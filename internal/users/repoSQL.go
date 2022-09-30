package users

import (
	"battles/internal/answer"
	"battles/internal/users/user_models"
	"battles/internal/utils/logger"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepoSQL struct {
	db *sqlx.DB
	lg *logrus.Logger
}

func NewUserRepoSQL(db *sqlx.DB) UserRepo {
	return &UserRepoSQL{db: db, lg: logger.Get()}
}
func (r *UserRepoSQL) CreateUserByUser(command user_models.CommandUserCreateByUser) *answer.Answer {
	_, err := r.db.Exec("INSERT INTO users(public_address) VALUES ($1);", command.User.PublicAddress)
	return &answer.Answer{Err: err}
}
func (r *UserRepoSQL) ReadUserByUserPublicAddress(query user_models.QueryUserReadByUserPublicAddress) *answer.Answer {
	var user user_models.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE public_address = $1;", query.PublicAddress)
	if err != nil {
		r.lg.Debugf(`ReadUserByUserPublicAddress query:"%v" err:"%v"`, query, err)
		return &answer.Answer{Err: err}
	}
	return &answer.Answer{Err: err, User: &user}
}
