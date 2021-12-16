package datamanager

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type Datamanager interface {
	AddUser(*User) error
	GetUser(id int) (*User, error)
	UpdateUser(*User) error
	VerifyAuth(username, password string) (int, error)

	AddChatMessage(*User, Message) error
	CreateChannel(*User) (int, error)
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

func (s *StudentLinkDatabase) AddUser(user *User, password string) error {
	hash := hashPassword(password)
	err := s.db.QueryRow("INSERT INTO users (email, username, firstname, lastname, type, pw_hash) Values ($1, $2, $3, $4, $5, $6)",
		user.Email, user.Username, user.Firstname, user.Lastname, user.Type, hash).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *StudentLinkDatabase) GetUser(id int) (*User, error) {
	var u *User = &User{}

	// todo: Populate scan
	err := s.db.QueryRow("SELECT * FROM users WHERE id=?", id).Scan()

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *StudentLinkDatabase) UpdateUser(user *User) error {
	err := s.db.QueryRow("UPDATE users SET email = ?", user.Email).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *StudentLinkDatabase) VerifyAuth(username, password string) (int, error) {
	if password == "hallo" {
		return 1, nil
	}

	return 0, errors.New("poop")
}

func (s *StudentLinkDatabase) AddChatMessage(user *User, msg Message) error {
	return nil
}

func (s *StudentLinkDatabase) CreateChannel(user *User) error {
	return nil
}

func hashPassword(plaintext string) string {
	pw, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(pw)
}
