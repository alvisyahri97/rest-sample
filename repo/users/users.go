package users

import (
	"errors"
)

type UserRepo struct {
	lastID   int
	userData map[int]User
}

func NewUser() *UserRepo {
	return &UserRepo{
		lastID: 1,
		userData: map[int]User{
			1: {
				ID:      1,
				Name:    "farhan",
				Address: "jl. kampung dadap",
				Phone:   "0809899999",
			},
		},
	}
}

func (u *UserRepo) GetByID(id int) (User, error) {
	user, ok := u.userData[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (u *UserRepo) InsertUser(user *User) error {
	if user == nil {
		return errors.New("new user not valid")
	}

	u.lastID += 1

	user.ID = u.lastID
	u.userData[u.lastID] = *user
	return nil
}

func (u *UserRepo) UpdateByID(id int, newUser User) error {
	user, ok := u.userData[id]
	if !ok {
		return errors.New("user not found")
	}

	if user == newUser {
		return nil
	}

	newUser.ID = id
	u.userData[id] = newUser
	return nil
}

func (u *UserRepo) DeleteByID(id int) error {
	_, ok := u.userData[id]
	if !ok {
		return errors.New("user not found")
	}
	delete(u.userData, id)
	return nil
}
