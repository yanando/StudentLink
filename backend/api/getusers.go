package api

import (
	"encoding/json"
	"net/http"
)

func (apiServer *APIServer) GetUsersHandler(rw http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("session")
	if cookie == "" {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	_, exists := apiServer.sessionManager.GetUserBySessionID(cookie)
	if !exists {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	users, err := apiServer.dataManager.GetUsers()
	if err != nil {
		rw.WriteHeader(ErrorFailedToGetUser.Code)
		rw.Write(ErrorFailedToGetUser.Unmarshal())
		return
	}

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	err = enc.Encode(users)
	if err != nil {
		rw.WriteHeader(ErrorWritingResponse.Code)
		rw.Write(ErrorWritingResponse.Unmarshal())
		return
	}
}
