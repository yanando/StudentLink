package datamanager

import "time"

type Message struct {
	ID int

	AuthorID  int
	ChannelID int

	Content     string
	CreatedDate time.Time
}
