package bankcore

import (
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	if account.Balance != 10 {
		t.Errorf("Balance should be 10 after deposit 10")
	}

	//if err := account.Deposit(-10); err != nil {
	//	t.Error("Could not deposit -10")
	//}
	account.Withdraw(10)
	if account.Balance != 0 {
		t.Errorf("Balance should be 0 after withdraw 10")
	}
	fmt.Println(account)
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}
