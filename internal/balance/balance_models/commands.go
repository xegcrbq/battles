package balance_models

type CommandCreateBalanceByBalance struct {
	Balance *ClientBalance
}
type QueryReadBalanceByUserId struct {
	UserId string
}
