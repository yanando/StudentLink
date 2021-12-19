package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yanando/StudentLink/datamanager"
)

type AddMessageRequest struct {
	Content     string `json:"content"`
	RecipientID int    `json:"recipient_id"`
}

func (apiServer *APIServer) AddMessageHandler(rw http.ResponseWriter, r *http.Request) {
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

	var addMessageRequest AddMessageRequest
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&addMessageRequest)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}

	err = apiServer.dataManager.AddChatMessage(datamanager.Message{
		AuthorID:    user.ID,
		RecipientID: addMessageRequest.RecipientID,
		Content:     addMessageRequest.Content,
	})
	if err != nil {
		rw.WriteHeader(ErrorFailedToSendMessage.Code)
		rw.Write(ErrorFailedToSendMessage.Unmarshal())
		return
	}
	rw.WriteHeader(http.StatusOK)

	log.Printf("User %d sent a message", user.ID)
}
