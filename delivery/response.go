package delivery

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, status int) {

	payload, err := json.Marshal(Response{
		Code:    status,
		Message: data,
	})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
