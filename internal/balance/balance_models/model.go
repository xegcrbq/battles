package balance_models

type Balance struct {
	BalanceId int64 `db:"balanceid"`
	UserId    int64 `db:"userid"`
	//Amount хранит условные сатоши, умножаем на 10^-8, чтобы получить реальное число монет
	Amount int64 `db:"amount"`
	CoinId int32 `db:"coinid"`
	Spent  int64 `db:"spent"`
}

type Pair struct {
	First  string
	Second string
	Price  float64
}

func (p *Pair) String() string {
	return p.First + p.Second
}
