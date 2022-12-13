package handlers

import (
	"Gateway/pkg/util"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type HttpErrorResponse struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
	Error     string    `json:"error"`
}

type ErrorResponse struct {
	Code         int    `json:"code"`
	Type         string `json:"type"`
	ErrorMessage string `json:"error"`
}

const httpLog string = "[HTTP]: "

// Writes an error response for http requests
func RenderErrorResponse(w http.ResponseWriter, message string, path string, err error) {
	log.Println(httpLog + err.Error())

	// Get the current timestamp
	now := time.Now()

	// Always set the status to Internal Server Error
	status := http.StatusInternalServerError

	var internalErr *util.Error
	if !errors.As(err, &internalErr) {
		// Create http error response
		response := &HttpErrorResponse{
			Message:   message,
			Status:    status,
			Timestamp: now,
			Path:      path,
			Error:     http.StatusText(status),
		}

		RenderResponse(w, status, response)
	} else {
		switch internalErr.ErrorCode() {
		case util.ErrorCodeNotFound:
			status = http.StatusNotFound
		case util.ErrorCodeInvalid:
			status = http.StatusBadRequest
		case util.ErrorCodeUnauthorized:
			status = http.StatusUnauthorized
		}

		// Create http error response
		response := &HttpErrorResponse{
			Message:   message,
			Status:    status,
			Timestamp: now,
			Path:      path,
			Error:     http.StatusText(status),
		}

		RenderResponse(w, status, response)
	}
}

// Writes a response for http requests with a payload
func RenderResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)

	w.Write(response)
}
