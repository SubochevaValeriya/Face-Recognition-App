package service

import (
	"errors"
	"html"
	"strings"

	"github.com/SubochevaValeriya/face-recognition-app/internal/models"
	"github.com/SubochevaValeriya/face-recognition-app/internal/repository"
	"github.com/SubochevaValeriya/face-recognition-app/internal/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type UserApiService struct {
	repo repository.User
}

func newUserApiService(repo repository.User) *UserApiService {
	return &UserApiService{repo: repo}
}

func (s *UserApiService) GetUserByID(uid uint) (models.User, error) {
	u, err := s.repo.GetUser(uid)
	if err != nil {
		return *u, errors.New("User not found!")
	}

	s.PrepareGive(u)

	return *u, nil

}

func (s *UserApiService) PrepareGive(u *models.User) {
	u.Password = ""
}

func (s *UserApiService) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserApiService) LoginCheck(username string, password string) (string, error) {

	var err error

	u, err := s.repo.GetUserByName(username)

	if err != nil {
		return "", err
	}

	err = s.VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *UserApiService) SaveUser(u *models.User) (*models.User, error) {
	return s.repo.CreateUser(u)
}

func (s *UserApiService) BeforeSave(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
