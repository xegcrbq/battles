package userbalancecoins_models

type CommandCreateBalanceByUserPublicAddressAndAmountAndTicker struct {
	UserPublicAddress string
	Amount            int64
	Ticker            string
}
type CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker struct {
	UserId int64
	Amount int64
	Spent  int64
	Ticker string
}
type CommandCreateBalanceByUserIdAmountSpentAndTicker struct {
	UserId int64
	Amount int64
	Spent  int64
	Ticker string
}
type CommandUpdateBalanceByUserIdAmountSpentAndTicker struct {
	UserId int64
	Amount int64
	Spent  int64
	Ticker string
}
type QueryReadCountByUserIdAndTicker struct {
	UserId int64
	Ticker string
}
type QueryReadUserBalanceCoinsByUserPublicAddress struct {
	UserPublicAddress string
}
type QueryReadUserBalanceAllCoinsByUserPublicAddress struct {
	UserPublicAddress string
}
