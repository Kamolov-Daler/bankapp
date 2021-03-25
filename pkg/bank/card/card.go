package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func IssueCard(currency types.Currency, color string, name string) (card types.Card) {
	card = types.Card{
		ID:         1000,
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    0,
		Currency:   currency,
		Color:      color,
		Name:       name,
		Active:     true,
		MinBalance: 0,
	}
	return
}

func Withdraw(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > card.Balance {
		return
	}
	if amount < 0 {
		return
	}
	if amount > 20_000_00 {
		return
	}
	card.Balance -= amount
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > 50_000_00 {
		return
	}
	if amount < 0 {
		return
	}
	if card.Balance < 0 {
		card.Balance += amount
		return
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}

	bonus := types.Money(((int(card.MinBalance) * percent) / 100) * daysInMonth / daysInYear)

	if bonus > 5_000_00 {
		return
	}

	card.Balance += bonus
}

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for i := 0; i < len(cards); i++ {
		if cards[i].Active {
			if cards[i].Balance > 0 {
				sum += cards[i].Balance
			}
		}
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSources {
	arrPayments := []types.PaymentSources{}

	for idx, card := range cards {
		if card.Balance > 0 {
			if card.Active {
				newIdx := "0" + fmt.Sprint(idx+1)
				arrPayments = append(arrPayments, types.PaymentSources{Type: "card", Number: newIdx, Balance: card.Balance})
			}
		}
	}
	return arrPayments
}
