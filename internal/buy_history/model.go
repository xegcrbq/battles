package buy_history

type BuyHistory struct {
	BuyHistoryId int64 `db:"buy_history_id"`
	UserId       int64 `db:"userid"`
	CoinId       int64 `db:"coinid"`
	//Sum is price in USDT
	Sum int64 `db:"sum"`
}

type BuyHistorySimple struct {
	Ticker string `db:"ticker"`
	//Sum is price in USDT
	Sum int64 `db:"sum"`
}
