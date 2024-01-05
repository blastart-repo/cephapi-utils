package cautils

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0]
		// status -1 doesn't overwrite existing status code
		c.JSON(-1, ErrorResponse{Error: err.Error()})
	}
}
