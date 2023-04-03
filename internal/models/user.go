package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/SubochevaValeriya/face-recognition-app/internal/service"
)

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

func (u *User) SaveUser() (*User, error) {
	_, err := service.DB.NamedExec(`INSERT INTO user (username, password) VALUES (:username, :password)`, &u)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
