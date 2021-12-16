package api

import (
	"encoding/json"
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
	ErrorReadingRequest Error = Error{"There was a problem reading your request", http.StatusBadRequest}
)
