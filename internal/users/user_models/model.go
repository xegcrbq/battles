package user_models

type User struct {
	UserId        int64  `db:"userid"`
	PublicAddress string `db:"public_address"`
}
