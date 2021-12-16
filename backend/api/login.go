package api

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
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

	if request.Username == "" || request.Password == "" {
		rw.WriteHeader(ErrorInvalidCredentials.Code)
		rw.Write(ErrorInvalidCredentials.Unmarshal())
		return
	}

	authorized, err := apiServer.dataManager.VerifyAuth(request.Username, request.Password)
	if err != nil || authorized {
		rw.WriteHeader(ErrorInvalidCredentials.Code)
		rw.Write(ErrorInvalidCredentials.Unmarshal())
		return
	}

	apiServer.dataManager.GetUser()
}
