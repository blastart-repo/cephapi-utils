package cautils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blastart-repo/cephapi-utils/config"
	"io"
	"net/http"
	"strings"
)

func init() {
	var err error
	config.Data, err = config.LoadConfig()
	if err != nil {
		fmt.Println(fmt.Sprintf("Reading UserIdRequest config failed: %s", err.Error()))
	}
}

func UserIDRequest(req *http.Request) (string, error) {

	bearerToken, err := extractBearerToken(req)
	if err != nil {
		return "", err
	}

	url := config.Data.UidUrl

	newReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating UserIDRequest:", err)
		return "", err
	}

	newReq.Header.Set("Authorization", "Bearer "+bearerToken)

	client := http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		fmt.Println("Error making UserIDRequest:", err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing UserIDRequest response body: %v\n", err)
		}
	}(resp.Body)
	var userInfo struct {
		Sub string `json:"sub"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		fmt.Printf("Error decoding response body: %v\n", err)
		return "", err
	}

	if userInfo.Sub != "" {
		return userInfo.Sub, nil
	}

	return "", errors.New("user id not found in response")
}

func extractBearerToken(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader != "" {
		authParts := strings.Split(authHeader, " ")
		if len(authParts) == 2 && strings.ToLower(authParts[0]) == "bearer" {
			return authParts[1], nil
		}
	}

	return "", errors.New("bearer token not found in Authorization header")
}
