package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (apiServer *APIServer) LoginHandler(rw http.ResponseWriter, r *http.Request) {
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

	userId, err := apiServer.dataManager.VerifyAuth(request.Username, request.Password)
	if err != nil {
		rw.WriteHeader(ErrorInvalidCredentials.Code)
		rw.Write(ErrorInvalidCredentials.Unmarshal())
		return
	}

	user, err := apiServer.dataManager.GetUser(userId)
	if err != nil {
		rw.WriteHeader(ErrorFailedToGetUser.Code)
		rw.Write(ErrorFailedToGetUser.Unmarshal())
		return
	}

	sessionId := apiServer.sessionManager.CreateSession(user)
	rw.Header().Set("auth", sessionId)

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	err = enc.Encode(user)
	if err != nil {
		rw.WriteHeader(ErrorWritingResponse.Code)
		rw.Write(ErrorWritingResponse.Unmarshal())
		return
	}

	log.Printf("User %s logged in", user.Username)
}
