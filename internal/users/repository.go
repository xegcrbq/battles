package users

import (
	"battles/internal/answer"
	"battles/internal/users/user_models"
)

type UserRepo interface {
	CreateUserByUser(command user_models.CommandUserCreateByUser) *answer.Answer
	ReadUserByUserPublicAddress(query user_models.QueryUserReadByUserPublicAddress) *answer.Answer
}
