package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

func (e Error) Unmarshal() []byte {
	bytez, _ := json.Marshal(e)
	return bytez
}

var (
	ErrorReadingRequest      Error = Error{"There was a problem reading your request", http.StatusBadRequest}
	ErrorWritingResponse     Error = Error{"There was a problem writing your response", http.StatusInternalServerError}
	ErrorUnauthorizedRequest Error = Error{"You are not authorized to use this", http.StatusUnauthorized}

	ErrorInvalidCredentials Error = Error{"Incorrect username or password", http.StatusUnauthorized}

	ErrorTooManyMessages func(amount int) Error = func(amount int) Error {
		return Error{fmt.Sprintf("You can only get %d messages at once", amount), http.StatusBadRequest}
	}
	ErrorFailedToGetMessages Error = Error{"There was a problem getting your messages", http.StatusInternalServerError}
	ErrorFailedToSendMessage Error = Error{"There was a problem sending your message", http.StatusInternalServerError}

	ErrorFailedToGetUser    Error = Error{"There was a problem getting your user", http.StatusInternalServerError}
	ErrorFailedToUpdateUser Error = Error{"There was a problem updating your user", http.StatusInternalServerError}
)
