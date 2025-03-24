package service

import (
	"github.com/barisaydogdu/jwt-auth/repository"
)

type UserService interface {
	Login(email, password string) (string, error)
	Register(firstname, lastname, username, email, password string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

//func NewUserService(repo repository.UserRepository) UserService {
//	return &userService{repo: repo}
//}

//func (u *userService) Login(email, password string) (string, error) {
//	user, err := u.repo.FindByEmail(email)
//}
//
//func (u *userService) Register(firstname, lastname, username, email, password string) (string, error) {
//	existingUser, _ := u.repo.FindByEmail(email)
//	if existingUser != nil {
//		return "", errors.New("User already exists")
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return "", errors.New("Error generating password")
//	}
//
//	user := &domain.User{FirstName: firstname, LastName: lastname, Username: username, Email: email, Password: string(hashedPassword)}\
//	err = u.repo.Create(user)
//	if err != nil {
//		return "", err
//	}
//
//}
