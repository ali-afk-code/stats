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
