package handlers

import (
	"context"
	"encoding/json"
	"github.com/barisaydogdu/jwt-auth/service"
	"net/http"
)

type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	UserService service.UserService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func NewUserHandler(ctx context.Context, userService service.UserService) UserHandler {
	return &userHandler{UserService: userService}
}

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := u.UserService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Error with login", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(AuthResponse{Token: token})
}

func (u userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	token, err := u.UserService.Register(req.FirstName, req.LastName, req.Username, req.Email, req.Password)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(AuthResponse{Token: token})
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
