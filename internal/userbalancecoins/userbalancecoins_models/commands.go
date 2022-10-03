package userbalancecoins_models

type QueryReadUserBalanceCoinsByUserPublicAddress struct {
	UserPublicAddress string
}
type QueryReadUserBalanceAllCoinsByUserPublicAddress struct {
	UserPublicAddress string
}

type CommandCreateBalanceByUserPublicAddressAndAmountAndTicker struct {
	UserPublicAddress string
	Amount            int64
	Ticker            string
}
type CommandUpdateOrCreateBalanceByUserIdAmountAndTicker struct {
	UserId int64
	Amount int64
	Ticker string
}
