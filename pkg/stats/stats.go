package stats

import "github.com/ali-afk-code/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, payment := range payments {
		if payment.Status != types.Status("FAIL") {
			sum += int(payment.Amount)
		}
	}
	return types.Money(sum / len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}
	return sum
}
