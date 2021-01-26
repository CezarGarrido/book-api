package delivery

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON : Retorna um json
func RespondWithJSON(w http.ResponseWriter, data interface{}, status int) {

	payload, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
