package user

import "github.com/perlinleo/golang-todo/internal/model"

type Usecase interface {
	CreateUser(user *model.User) ([]model.User, error)
	Find(nickname string) (*model.User, error)
	Update(user *model.User) (*model.User, error, int)
	DuplicateUser(user *model.User) ([]model.User, error)
}