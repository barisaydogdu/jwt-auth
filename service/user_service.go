package service

import (
	"context"
	"errors"
	"github.com/barisaydogdu/jwt-auth/domain"
	"github.com/barisaydogdu/jwt-auth/repository"
	"github.com/barisaydogdu/jwt-auth/utils"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(email, password string) (string, error)
	Register(firstname, lastname, username, email, password string) (string, error)
}

type userService struct {
	ctx  context.Context
	repo repository.UserRepository
	db   *pgx.Conn
}

func NewUserService(ctx context.Context, repo repository.UserRepository) UserService {
	return &userService{repo: repo,
		ctx: ctx}
}

func (u *userService) Login(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.CreateToken(email)
	if err != nil {
		return "", errors.New("error creating token")
	}

	return token, nil

}

func (u *userService) Register(firstname, lastname, username, email, password string) (string, error) {
	existingUser, _ := u.repo.FindByEmail(email)
	if existingUser != nil {
		return "", errors.New("User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Error generating password")
	}

	user := &domain.User{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword)}

	err = u.repo.Create(user)
	if err != nil {
		return "", err
	}

	token, err := utils.CreateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
