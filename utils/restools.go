package utils

import (
	"encoding/json"
	"net/http"
)

type standardAPIErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func outputJSON(respCode int, payload interface{}) (int, []byte) {
	output, err := json.Marshal(payload)
	if err != nil {
		internalServerError := &standardAPIErrorResponse{
			Message: err.Error(),
		}
		output, _ = json.Marshal(internalServerError)
		return http.StatusInternalServerError, output
	}
	return respCode, output
}

// OutputInternalServerError will return the standard error format with the internal server error
func OutputInternalServerError(w http.ResponseWriter, err error) {
	output := &standardAPIErrorResponse{
		Message: err.Error(),
	}

	code, payload := outputJSON(http.StatusInternalServerError, output)
	w.WriteHeader(code)
	w.Write(payload)
}

// OutputBadRequest ...
func OutputBadRequest(w http.ResponseWriter, err error) {
	output := &standardAPIErrorResponse{
		Message: err.Error(),
	}

	code, payload := outputJSON(http.StatusBadRequest, output)
	w.WriteHeader(code)
	w.Write(payload)
}

func OutputNotFound(w http.ResponseWriter, err error) {
	output := &standardAPIErrorResponse{
		Message: err.Error(),
	}

	code, payload := outputJSON(http.StatusNotFound, output)
	w.WriteHeader(code)
	w.Write(payload)
}

func OutputData(w http.ResponseWriter, code int, output interface{}) {
	code, payload := outputJSON(code, output)
	w.WriteHeader(code)
	w.Write(payload)
}
