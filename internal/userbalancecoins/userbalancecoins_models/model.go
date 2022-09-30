package userbalancecoins_models

type UserBalanceCoins struct {
	PublicAddress string `db:"public_address"`
	Ticker        string `db:"ticker"`
	//Amount хранит условные сатоши, умножаем на 10^-8, чтобы получить реальное число монет
	Amount int64 `db:"amount"`
}
