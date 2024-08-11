package controller

import (
	"net/http"
)

type userDomain interface {
	Create() error
}

type user struct {
	userDomain userDomain
}

func NewUserController(userDomain userDomain) *user {
	return &user{
		userDomain: userDomain,
	}
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	if err := u.userDomain.Create(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
