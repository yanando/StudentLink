package datamanager

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type Datamanager interface {
	AddUser(user *User, password string) error
	GetUser(id int) (*User, error)
	GetUsers() ([]*User, error)
	UpdateUser(user *User) error
	VerifyAuth(username, password string) (int, error)

	AddChatMessage(msg Message) error
	GetChatMessages(authorID, recipientID, amount, offset int) ([]Message, error)
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
	user := &User{}

	err := s.db.QueryRow("SELECT id,type,username,firstname,lastname,email FROM users WHERE id=?", id).Scan(&user.ID, &user.Type, &user.Username, &user.Firstname, &user.Lastname, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *StudentLinkDatabase) GetUsers() ([]*User, error) {
	users := []*User{}

	rows, err := s.db.Query("SELECT id,type,username,firstname,lastname,email FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.ID, &user.Type, &user.Username, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *StudentLinkDatabase) UpdateUser(user *User) error {
	err := s.db.QueryRow("UPDATE users SET email = ?", user.Email).Err()

	if err != nil {
		return err
	}

	return nil
}

func (s *StudentLinkDatabase) VerifyAuth(username, password string) (int, error) {
	var dbpw string
	var id int

	err := s.db.QueryRow("SELECT id, pw_hash FROM users WHERE username = ?", username).Scan(&id, &dbpw)

	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbpw), []byte(password))

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *StudentLinkDatabase) AddChatMessage(msg Message) error {
	return s.db.QueryRow("INSERT INTO messages (author_id, recipient_id, content, created_timestamp) VALUES ($1, $2, $3, $4)",
		msg.AuthorID, msg.RecipientID, msg.Content, time.Now().UnixNano()).Err()
}

func (s *StudentLinkDatabase) GetChatMessages(authorID, recipientID, amount, offset int) ([]Message, error) {
	res, err := s.db.Query("SELECT * FROM messages WHERE (author_id = $1 AND recipient_id = $2) OR (author_id = $2 AND recipient_id = $1) ORDER BY created_timestamp ASC LIMIT $3 OFFSET $4",
		authorID, recipientID, amount, offset)

	if err != nil {
		return nil, err
	}

	// Empty message met var declaration zou nil slice value hebben met json => null. Zero-length wordt een empty array in json
	messages := []Message{}
	for res.Next() {
		var message Message
		var dateInt int64
		err := res.Scan(&message.ID, &message.AuthorID, &message.RecipientID, &message.Content, &dateInt)

		if err != nil {
			return nil, err
		}

		message.CreatedDate = time.Unix(0, dateInt)
		messages = append(messages, message)
	}

	return messages, nil
}

func hashPassword(plaintext string) string {
	pw, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(pw)
}
