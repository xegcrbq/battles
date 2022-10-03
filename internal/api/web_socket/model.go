package web_socket

type WSReq struct {
	ReqType          string             `json:"req_type"`
	UserBalanceCoins WSUserBalanceCoins `json:"user_balance_coins"`
}
type WSUserBalanceCoins struct {
	//PublicAddress string `db:"public_address"`
	Ticker string `db:"ticker"`
	//Amount хранит float, нужно перевести в int умножив на 10^8, чтобы получить реальное число монет
	Amount string `db:"amount"`
}
