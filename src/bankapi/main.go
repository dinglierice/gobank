package main

import (
	"bankcore"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bankcore.Account{}

func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqstr := req.URL.Query().Get("number")
	amountstr := req.URL.Query().Get("amount")

	if numberqstr == "" {
		fmt.Fprintf(w, "Account number and amount is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqstr, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountstr, 64); err != nil {
		fmt.Fprintf(w, "Invalid account amount!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with amount %v can't be found!", amount)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "Account with amount %v deposit failed!", amount)
			} else {
				fmt.Fprintf(w, "Account with amount %v successfully deposited!", amount)
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqstr := req.URL.Query().Get("number")
	amountstr := req.URL.Query().Get("amount")

	if numberqstr == "" {
		fmt.Fprintf(w, "Account number and amount is missing!")
	}

	if number, err := strconv.ParseFloat(numberqstr, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountstr, 64); err != nil {
		fmt.Fprintf(w, "Invalid account amount!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		}
		err := account.Withdraw(amount)
		if err != nil {
			fmt.Fprintf(w, "Account with amount %v deposit failed!", amount)
		} else {
			fmt.Fprintf(w, "Account with amount %v successfully withdrawn!", amount)
		}
	}
}
