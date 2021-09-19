package stats

import (
	"reflect"
	"testing"

	"github.com/ali-afk-code/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			Category: "shop",
			Amount:   3,
		},
		{
			Category: "shop",
			Amount:   2,
		},
		{
			Category: "stuff",
			Amount:   4,
		},
	}
	expected := map[types.Category]types.Money{
		"shop":  2,
		"stuff": 4,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, got: %v\n", expected, result)
	}
}

func TestPeriodsDynamic(t *testing.T) {
	payments1 := map[types.Category]types.Money{
		"shop":  3,
		"stuff": 4,
		"other": 10,
	}
	payments2 := map[types.Category]types.Money{
		"shop":  2,
		"stuff": 5,
	}

	result := PeriodsDynamic(payments1, payments2)
	expected := map[types.Category]types.Money{
		"shop":  -1,
		"stuff": 1,
		"other": -10,
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, got: %v\n", expected, result)
	}
}
