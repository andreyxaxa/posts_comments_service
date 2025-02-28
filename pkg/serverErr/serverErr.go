package servererr

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ServerErr struct {
	StatusCode int
	Message    string
}

func (e ServerErr) Err() string {
	return e.Message
}

func ErrorResponse(w http.ResponseWriter, err error) {
	var serverErr ServerErr
	var statusCode int
	var errorMessage string

	ok := errors.As(err, &serverErr)
	if !ok {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = serverErr.StatusCode
	}

	errorMessage = err.Error()
	w.WriteHeader(statusCode)

	if errorMessage != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response := make(map[string]string)
		response["error"] = errorMessage
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
