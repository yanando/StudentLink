package datamanager

import (
	"zombiezen.com/go/sqlite"
)

type Datamanager interface {
	GetUser() (User, error)
	UpdateUser(User) error
	VerifyAuth(email, password string) (bool, error)

	AddChatMessage(User, Message) error
	CreateChannel(User) (int, error)
}

type StudentLinkDatabase struct {
	conn *sqlite.Conn
}

func (s *StudentLinkDatabase) Start() error {
	conn, err := sqlite.OpenConn("database.db")

	if err != nil {
		return err
	}

	s.conn = conn
	return nil
}

func (s *StudentLinkDatabase) Close() error {
	return s.conn.Close()
}
