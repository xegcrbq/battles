package userbalancecoins

import (
	"battles/internal/answer"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"github.com/jmoiron/sqlx"
)

type UserBalanceCoinsRepoSQL struct {
	db *sqlx.DB
}

func NewUserBalanceCoinsRepoSQL(db *sqlx.DB) UserBalanceCoinsRepo {
	return &UserBalanceCoinsRepoSQL{db: db}
}

func (r *UserBalanceCoinsRepoSQL) ReadUserBalanceCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress) *answer.Answer {
	var userBalanceCoins []userbalancecoins_models.UserBalanceCoins
	err := r.db.Select(&userBalanceCoins, `
SELECT balances.amount, c.ticker
FROM
    users
    INNER JOIN balances
        ON users.userid = balances.userid
    INNER JOIN coins c on c.coinid = balances.coinid
WHERE  users.public_address = $1;
`, query.UserPublicAddress)
	return &answer.Answer{Err: err, UserBalanceCoins: &userBalanceCoins}
}
func (r *UserBalanceCoinsRepoSQL) ReadUserBalanceAllCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress) *answer.Answer {
	var userBalanceCoins []userbalancecoins_models.UserBalanceCoins
	err := r.db.Select(&userBalanceCoins, `
SELECT COALESCE(balances.amount, 0) as amount, c.ticker
FROM
    (SELECT public_address, userid from users where public_address = $1) as users
        INNER JOIN balances
                   ON users.userid = balances.userid
        Right JOIN coins c on c.coinid = balances.coinid;
`, query.UserPublicAddress)
	return &answer.Answer{Err: err, UserBalanceCoins: &userBalanceCoins}
}

func (r *UserBalanceCoinsRepoSQL) CreateBalanceByUserPublicAddressAndAmountAndTicker(command userbalancecoins_models.CommandCreateBalanceByUserPublicAddressAndAmountAndTicker) *answer.Answer {
	_, err := r.db.Exec(`
INSERT INTO balances(userid, amount, coinid)
VALUES
    (
     (Select userid from users where public_address = $1),
     $2,
     (Select coinid FROM coins where ticker=$3)
     );
`, command.UserPublicAddress, command.Amount, command.Ticker)
	return &answer.Answer{Err: err}
}

// UpdateOrCreateBalanceByUserIdAmountAndTicker нужно деласть через сервис с выховом из репозитория чтения, затем селекта
func (r *UserBalanceCoinsRepoSQL) UpdateOrCreateBalanceByUserIdAmountAndTicker(command userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountAndTicker) *answer.Answer {
	count := 0
	err := r.db.Get(&count, `
	select count(*) 
	from 
	    balances
		INNER JOIN coins c on c.coinid = balances.coinid
	where userid = $1 and ticker = $2 `,
		command.UserId, command.Ticker)
	if err != nil {
		return &answer.Answer{Err: err}
	}
	if count == 0 {
		_, err = r.db.Exec(`insert into balances(userid, amount, coinid) VALUES ($1, $2, (select coinid from coins where ticker = $3));`,
			command.UserId, command.Amount, command.Ticker)
	} else {
		_, err = r.db.Exec(`UPDATE balances set amount = $1 where coinid = (select coinid from coins where ticker = $2) and userid = $3;`,
			command.Amount, command.Ticker, command.UserId)
	}
	return &answer.Answer{Err: nil}
}
