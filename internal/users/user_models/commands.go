package user_models

type CommandUserCreateByUser struct {
	User *User
}
type QueryUserReadByUserPublicAddress struct {
	PublicAddress string
}
