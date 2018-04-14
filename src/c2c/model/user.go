package model

import "fmt"

// User model
type User struct {
	ID       int
	Email    string
	Password string
}

func (user *User) String() string {
	return fmt.Sprintf("user(%d, %s, %s)", user.ID, user.Email, user.Password)
}
