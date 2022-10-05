package balance_models

type CommandCreateBalanceByBalance struct {
	Balance *Balance
}
type QueryReadBalancesByUserId struct {
	UserId int64
}
type QueryReadBalanceByUserIdAndCoinId struct {
	UserId int64
	CoinId int32
}
