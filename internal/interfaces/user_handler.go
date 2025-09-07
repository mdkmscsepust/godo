package user_handler

import (
	"fmt"
	"godo/internal/application/user"
	"net/http"
)

type UserHandler struct {
	usecase *user.UseCase
}

func NewUserHandler(uc *user.UseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (userHandler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	u, err := userHandler.usecase.RegisterUser(name, email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User registered: %s", u.Id)
}
