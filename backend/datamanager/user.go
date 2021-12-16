package datamanager

type userType string

type User struct {
	ID   int
	Type userType

	Username  string
	Firstname string
	Lastname  string
	Email     string
}
