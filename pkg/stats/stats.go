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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	res := map[types.Category]types.Money{}
	counter := map[types.Category]int{}
	for _, payment := range payments {
		if _, ok := res[payment.Category]; ok {
			res[payment.Category] += payment.Amount
			counter[payment.Category] += 1
		} else {
			res[payment.Category] = payment.Amount
			counter[payment.Category] = 1
		}
	}
	for category, money := range res {
		res[category] = money / types.Money(counter[category])
	}
	return res
}
