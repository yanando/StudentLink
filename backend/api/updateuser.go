package api

import (
	"encoding/json"
	"net/http"
)

type UpdateUserRequest struct {
	Email string `json:"email"`
}

func (apiServer *APIServer) UpdateUserHandler(rw http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("session")
	if cookie == "" {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	var request UpdateUserRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&request)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}

	user, exists := apiServer.sessionManager.GetUserBySessionID(cookie)
	if !exists {
		rw.WriteHeader(ErrorUnauthorizedRequest.Code)
		rw.Write(ErrorUnauthorizedRequest.Unmarshal())
		return
	}

	user.Email = request.Email

	err = apiServer.dataManager.UpdateUser(user)
	if err != nil {
		rw.WriteHeader(ErrorFailedToUpdateUser.Code)
		rw.Write(ErrorFailedToUpdateUser.Unmarshal())
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
