package buy_history

import (
	"battles/internal/db"
	"battles/internal/utils/logger"
	"testing"
)

func TestReadBuyHistorySimpleByUserId(t *testing.T) {
	bhr := NewBuyHistoryRepoSQL(db.Get())
	answ := bhr.ReadBuyHistorySimpleByUserId(QueryReadBuyHistorySimpleByUserId{UserId: 1})
	logger.Get().Debugf("ReadBuyHistorySimpleByUserId map:%v\n err: %v", answ.BuyHistorySums, answ.Err)
}
