package coins

type Coin struct {
	CoinId int16  `db:"coinid"`
	Ticker string `db:"ticker"`
}
