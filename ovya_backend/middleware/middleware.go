package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ExtractQueryId(req *http.Request) (int, error) {
	reqId := req.URL.Query().Get("id")
	if reqId == "" {
		return 0, fmt.Errorf("id is required")
	}
	id, err := strconv.Atoi(reqId)
	if err != nil {
		return 0, fmt.Errorf("id is invalid")
	}
	return id, nil
}

func JsonErrorResponse(message, details string) string {
	errorResponse := map[string]interface{}{
		"error":   message,
		"details": details,
	}
	response, _ := json.Marshal(errorResponse)
	return string(response)
}
