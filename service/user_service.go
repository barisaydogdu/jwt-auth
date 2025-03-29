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
	ctx   context.Context
	repo  repository.UserRepository
	utils utils.Utils
	db    *pgx.Conn
}

func NewUserService(repo repository.UserRepository, utils2 utils.Utils, db *pgx.Conn) UserService {
	return &userService{repo: repo,
		utils: utils2,
		db:    db,
	}
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
	util := utils.NewUtils(u.ctx, u.db)

	token, err := util.CreateToken(email)
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

	util := utils.NewUtils(u.ctx, u.db)

	token, err := util.CreateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
