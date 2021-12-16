package api

import (
	"encoding/json"
	"net/http"
)

func (apiServer *APIServer) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	user, exists := apiServer.sessionManager.GetUserBySessionID(cookie.Value)
	if !exists {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	err = enc.Encode(user)
	if err != nil {
		rw.WriteHeader(ErrorWritingResponse.Code)
		rw.Write(ErrorWritingResponse.Unmarshal())
		return
	}
}
