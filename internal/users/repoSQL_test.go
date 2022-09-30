package users

import (
	"battles/internal/db"
	"battles/internal/users/user_models"
	"battles/internal/utils/logger"
	"testing"
)

var lg = logger.Get()

func TestCreate(t *testing.T) {
	repo := NewUserRepoSQL(db.Get())
	answ := repo.CreateUserByUser(user_models.CommandUserCreateByUser{User: &user_models.User{PublicAddress: `0x8a8cB39FBE932c2fBED13B4982e4fE1BE364d58C`}})
	lg.Debug(answ)
}
func TestRead(t *testing.T) {
	repo := NewUserRepoSQL(db.Get())
	answ := repo.ReadUserByUserPublicAddress(user_models.QueryUserReadByUserPublicAddress{PublicAddress: `0x8a8cB99FBE932c2fBED13B4982e4fE1BE364d58C`})
	lg.Debug(*answ.User)
}
