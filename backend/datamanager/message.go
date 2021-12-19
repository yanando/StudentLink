package datamanager

import "time"

type Message struct {
	ID int `json:"id"`

	AuthorID    int `json:"author_id"`
	RecipientID int `json:"recipient_id"`

	Content     string    `json:"content"`
	CreatedDate time.Time `json:"created_date"`
}
