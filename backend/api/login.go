package api

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (apiServer APIServer) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&request)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}
}
