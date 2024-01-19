package cautils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UserIDRequest(req *http.Request, idUrl string) (string, error) {

	bearerToken, err := extractBearerToken(req)
	if err != nil {
		return "", err
	}

	newReq, err := http.NewRequest("GET", idUrl, nil)
	if err != nil {
		return "", fmt.Errorf("error creating UserIDRequest: %w", err)
	}

	newReq.Header.Set("Authorization", "Bearer "+bearerToken)

	client := http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		return "", fmt.Errorf("error making UserIDRequest: %w", err)
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
		return "", fmt.Errorf("error decoding response body: %w", err)
	}

	if userInfo.Sub != "" {
		return userInfo.Sub, nil
	}

	return "", errors.New("user id not found in the response body")
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
