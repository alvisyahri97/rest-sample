package user

import "github.com/alvisyahri97/rest-sample/repo/users"

type UserRepo interface {
	GetByID(id int) (users.User, error)
	InsertUser(user *users.User) error
	UpdateByID(id int, newUser users.User) error
	DeleteByID(id int) error
}
type Handler struct {
	userRepo UserRepo
}

func NewUserHandler(userRepo *users.UserRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
	}
}
