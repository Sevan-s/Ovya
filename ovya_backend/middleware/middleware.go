package middleware

import (
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
