package datamanager

type userType string

type User struct {
	ID   int      `json:"id"`
	Type userType `json:"type"`

	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
