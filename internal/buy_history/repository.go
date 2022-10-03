package buy_history

import "battles/internal/answer"

type BuyHistoryRepo interface {
	CreateBuyHistoryByBuyHistory(command CommandCreateBuyHistoryByBuyHistory) *answer.Answer
	ReadBuyHistorySimpleByUserId(query QueryReadBuyHistorySimpleByUserId) *answer.Answer
}