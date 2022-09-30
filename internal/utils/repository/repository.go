package repository

import (
	"battles/internal/answer"
	"battles/internal/db"
	"battles/internal/users"
	"battles/internal/users/user_models"
	"battles/internal/utils/errors_custom"
)

type Repository struct {
	ur users.UserRepo
}

func NewRepository() *Repository {
	return &Repository{ur: users.NewUserRepoSQL(db.Get())}
}

func (r *Repository) Exec(command interface{}) *answer.Answer {
	switch command.(type) {
	case user_models.CommandUserCreateByUser:
		return r.ur.CreateUserByUser(command.(user_models.CommandUserCreateByUser))
	case user_models.QueryUserReadByUserPublicAddress:
		return r.ur.ReadUserByUserPublicAddress(command.(user_models.QueryUserReadByUserPublicAddress))
	default:
		return &answer.Answer{
			Err: errors_custom.CommandNotFound,
		}
	}
}
