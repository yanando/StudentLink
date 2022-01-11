package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	MaxMessageAmount = 100
)

type GetMessagesRequest struct {
	RecipientID int `json:"recipient_id"`
	Amount      int `json:"amount"`
	Offset      int `json:"offset"`
}

func (apiServer *APIServer) GetMessagesHandler(rw http.ResponseWriter, r *http.Request) {
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

	values := r.URL.Query()
	recipientIdStr := values.Get("recipient_id")
	amountStr := values.Get("amount")
	offsetStr := values.Get("offset")

	recipientId, err := strconv.Atoi(recipientIdStr)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		rw.WriteHeader(ErrorReadingRequest.Code)
		rw.Write(ErrorReadingRequest.Unmarshal())
		return
	}

	if amount > MaxMessageAmount {
		rw.WriteHeader(ErrorTooManyMessages(MaxMessageAmount).Code)
		rw.Write(ErrorTooManyMessages(MaxMessageAmount).Unmarshal())
		return
	}

	messages, err := apiServer.dataManager.GetChatMessages(user.ID, recipientId, amount, offset)
	if err != nil {
		rw.WriteHeader(ErrorFailedToGetMessages.Code)
		rw.Write(ErrorFailedToGetMessages.Unmarshal())
		return
	}

	rw.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(rw)
	err = enc.Encode(messages)
	if err != nil {
		rw.WriteHeader(ErrorWritingResponse.Code)
		rw.Write(ErrorWritingResponse.Unmarshal())
		return
	}
}
