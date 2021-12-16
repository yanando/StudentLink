package datamanager

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

type Datamanager interface {
	GetUser(id int) (User, error)
	UpdateUser(User) error
	VerifyAuth(username, password string) (int, error)

	AddChatMessage(User, Message) error
	CreateChannel(User) (int, error)
}

type StudentLinkDatabase struct {
	db *sql.DB
}

func (s *StudentLinkDatabase) Start() error {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *StudentLinkDatabase) Close() error {
	return s.db.Close()
}

// func (s *StudentLinkDatabase) GetUser(id int) (User, error) {
// 	s.db.QueryRow("SELECT * FROM Users WHERE ID=?", id).Scan()
// }

func (s *StudentLinkDatabase) VerifyAuth(username, password string) (int, error) {
	if password == "hallo" {
		return 1, nil
	}

	return 0, errors.New("poop")
}
