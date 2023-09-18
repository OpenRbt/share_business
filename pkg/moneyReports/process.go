package moneyreports

import (
	"washbonus/internal/entities"

	"github.com/shopspring/decimal"
)

func ProcessBonusesReward(report entities.UserMoneyReport, percent decimal.Decimal) decimal.Decimal {
	moneyTypes := []struct {
		name  string
		value int64
	}{
		{"Coins", int64(report.Coins)},
		{"Banknotes", int64(report.Banknotes)},
		{"Electronical", int64(report.Electronical)},
	}

	divider := decimal.NewFromInt(100)
	addAmount := decimal.Zero

	for _, moneyType := range moneyTypes {
		moneyValue := decimal.NewFromInt(moneyType.value).
			Div(divider).
			Mul(percent)
		addAmount = addAmount.Add(moneyValue)
	}

	return addAmount
}
