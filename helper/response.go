package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Angle   int    `json:"angle"`
}

func BuildResponseBody(message string, success bool, angle int) ResponseBody {
	response := ResponseBody{
		Success: success,
		Message: message,
		Angle:   angle,
	}
	return response
}

func WriteResponse(w http.ResponseWriter, status int, response ResponseBody) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
