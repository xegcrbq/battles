package coins

import (
	"battles/internal/answer"
	"battles/internal/coins/coins_model"
	"github.com/jmoiron/sqlx"
)

type CoinsRepoSQL struct {
	db *sqlx.DB
}

func NewCoinsRepoSQL(db *sqlx.DB) CoinsRepo {
	return &CoinsRepoSQL{db: db}
}

func (r *CoinsRepoSQL) CreateCoin(command coins_model.CommandCreateCoinByTicker) *answer.Answer {
	_, err := r.db.Exec(`INSERT INTO coins(ticker) values ($1);`, command.Ticker)
	return &answer.Answer{Err: err}
}

func (r *CoinsRepoSQL) GetCoinsCount(query coins_model.QueryReadCoinsCount) *answer.Answer {
	count := 0
	err := r.db.Get(&count, `SELECT count(*) FROM coins`)
	return &answer.Answer{Err: err, Count: count}
}
