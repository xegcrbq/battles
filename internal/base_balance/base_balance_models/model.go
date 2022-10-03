package base_balance_models

type BaseBalance struct {
	BaseBalanceId int64 `db:"base_balance_id"`
	UserId        int64 `db:"userid"`
	Amount        int64 `db:"amount"`
}
