package datamanager

import "time"

type Message struct {
	ID int

	AuthorID    int
	RecipientID int

	Content     string
	CreatedDate time.Time
}
