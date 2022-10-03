package base_balance_models

type CommandCreateBaseBalanceByBaseBalance struct {
	BaseBalance *BaseBalance
}
type QueryReadBaseBalanceByUserId struct {
	UserId int64
}
type CommandUpdateBaseBalanceByBaseBalance struct {
	BaseBalance *BaseBalance
}
