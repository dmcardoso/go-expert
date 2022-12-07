package database

import "github.com/dmcardoso/go-expert/7-apis/5-user-db/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
