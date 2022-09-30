package balance_models

type ClientBalance struct {
	BalanceId int64
	UserId    int64 `db:"userid"`
	//Amount хранит условные сатоши, делим на 10^-8, чтобы получить реальное число монет
	Amount int64
	CoinId int64
}
type Pair struct {
	First  string
	Second string
	Price  float64
}

func (p *Pair) String() string {
	return p.First + p.Second
}
