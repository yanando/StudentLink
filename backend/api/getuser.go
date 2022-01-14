package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (apiServer *APIServer) GetUserHandler(rw http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("session")
	if cookie == "" {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	user, exists := apiServer.sessionManager.GetUserBySessionID(cookie)
	if !exists {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	// TODO: Temp code
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr != "" {
		var err error
		id, _ := strconv.Atoi(idStr)
		user, err = apiServer.dataManager.GetUser(id)
		if err != nil {
			rw.WriteHeader(ErrorFailedToGetUser.Code)
			rw.Write(ErrorFailedToGetUser.Unmarshal())
			return
		}
	}
	// ==============

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	err := enc.Encode(user)
	if err != nil {
		rw.WriteHeader(ErrorWritingResponse.Code)
		rw.Write(ErrorWritingResponse.Unmarshal())
		return
	}
}
