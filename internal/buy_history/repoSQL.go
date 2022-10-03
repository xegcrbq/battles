package buy_history

import (
	"battles/internal/answer"
	"battles/internal/utils/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math"
)

type BuyHistoryRepoSQL struct {
	db *sqlx.DB
}

func NewBuyHistoryRepoSQL(db *sqlx.DB) BuyHistoryRepo {
	return &BuyHistoryRepoSQL{db: db}
}

func (r *BuyHistoryRepoSQL) CreateBuyHistoryByBuyHistory(command CommandCreateBuyHistoryByBuyHistory) *answer.Answer {
	_, err := r.db.Exec(`insert into buy_history(userid, coinid, sum) 
		VALUES ($1, $2, $3)`, command.BuyHistory.UserId, command.BuyHistory.CoinId, command.BuyHistory.Sum)
	logger.Get().Debugf("CreateBuyHistoryByBuyHistory executed with command %v\n err %v", command.BuyHistory, err)
	return &answer.Answer{Err: err}
}

func (r *BuyHistoryRepoSQL) ReadBuyHistorySimpleByUserId(query QueryReadBuyHistorySimpleByUserId) *answer.Answer {
	var buyHistorySimple []BuyHistorySimple
	err := r.db.Select(&buyHistorySimple, `
		SELECT ticker, sum(sum) as sum FROM
			buy_history
			inner join coins c on c.coinid = buy_history.coinid
		where userid = $1 group by ticker;`, query.UserId)
	if err != nil {
		return &answer.Answer{Err: err}
	}
	answMap := make(map[string]string)
	for i := range buyHistorySimple {
		answMap[buyHistorySimple[i].Ticker] = fmt.Sprintf("%.8f", float64(buyHistorySimple[i].Sum%int64(math.Pow10(8)))*math.Pow10(-8)+float64(buyHistorySimple[i].Sum/int64(int(math.Pow10(8)))))
	}
	return &answer.Answer{
		Err:            err,
		BuyHistorySums: &answMap,
	}
}
