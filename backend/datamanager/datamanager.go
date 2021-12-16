package datamanager

import (
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type Datamanager interface {
	GetUser(id int) (User, error)
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

func (s *StudentLinkDatabase) GetUser(id int) (User, error) {
	sqlitex.Exec(s.conn, "SELECT * FROM Users WHERE ID = ")
}
