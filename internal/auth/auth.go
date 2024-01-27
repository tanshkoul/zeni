package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentication information")
	}
	if vals[0] != "APIKey" {
		return "", errors.New("invalid authentication information")
	}
	return vals[1], nil

}
