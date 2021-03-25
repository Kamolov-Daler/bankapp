package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 1000000

}
func ExampleWithdraw() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)

	fmt.Println(card.Balance)
	// Output:
	// 1000000

}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 0)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 21_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleCard() {
	card := types.Card{Balance: 20_000_00, Active: true}
	clone := card
	card.Balance -= 10_000_00
	fmt.Println(clone.Balance)
	// Output: 2000000
}

func ExampleDeposit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleDeposit_negativeBalance() {
	card := types.Card{Balance: -1_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 900000
}
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleDeposit_linit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 51_000_00)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 1, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 2466
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -3, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: -3
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 1, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1
}

func ExampleAddBonus_limitEqual() {
	card := types.Card{Balance: 1, Active: false, MinBalance: 5_000_000_00}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)
	// Output: 1
}

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))
	// Output:
	// 100000
	// 300000
	// 0
	// 0
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 1_000_00,
			Active:  false,
		},
		{
			Balance: 3_000_00,
			Active:  true,
		},
	}

	payments := PaymentSources(cards)

	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

	// Output:
	// 5058 xxxx xxxx 8888
	// 5058 xxxx xxxx 8888
}
