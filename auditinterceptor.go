package cautils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blastart-repo/cephapi-utils/audit"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateLog(c audit.AuditServiceClient) gin.HandlerFunc {
	return func(g *gin.Context) {
		req := g.Request
		pathvars := strings.Split(req.URL.Path, "/")[2:]
		var err error
		var sb strings.Builder

		if req.Method == http.MethodGet || req.Method == http.MethodDelete {
			sb.WriteString(pathvars[len(pathvars)-1])
		} else {
			body := struct {
				Name   string `json:"cluster_name,omitempty"`
				Bucket string `json:"bucket,omitempty"`
				UserID string `json:"user_id,omitempty"`
			}{}
			err = json.NewDecoder(req.Body).Decode(&body)
			if err != nil {
				g.Error(err)
				sb.WriteString("")
			} else {
				sb.WriteString(body.Name)
				sb.WriteString(body.Bucket)
				if body.Bucket == "" {
					sb.WriteString(body.UserID)
				}
			}
		}

		userID := ""
		sub, err := UserIDRequest(req)
		if err != nil {
			fmt.Printf("error geting the userID: %v\n", err)
		} else {
			userID = sub
		}

		logIn := audit.LogMessage{
			Status:      "",
			Endpoint:    req.URL.Path,
			Method:      req.Method,
			Description: sb.String(),
			User:        userID,
			IP:          g.ClientIP(),
			Timestamp:   time.Now().String(),
		}

		_, err = c.SendLog(context.Background(), &logIn)
		if err != nil {
			g.AbortWithError(http.StatusInternalServerError, errors.New(fmt.Sprintf("error writing audit log: %s", err.Error())))
			return
		}

		g.Next()

		logIn.Status = strconv.Itoa(g.Writer.Status())
		_, _ = c.SendLog(context.Background(), &logIn)
	}
}
