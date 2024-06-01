package main

import (
	"bankcore"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//map[float64]*CustomAccount{} 创建了一个空的 accounts 映射，其中键的类型为 float64，值的类型为 *CustomAccount
var accounts = map[float64]*CustomAccount{}

func main() {
	accounts[1001] = &CustomAccount{
		Account: &bankcore.Account{
			Customer: bankcore.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number: 1001,
		},
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func transfer(writer http.ResponseWriter, request *http.Request) {
	numberqs := request.URL.Query().Get("number")
	amountqs := request.URL.Query().Get("amount")
	toAccountqs := request.URL.Query().Get("dest")

	if numberqs == "" {
		fmt.Fprintf(writer, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(writer, "Account number is invalid!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(writer, "Account amount is invalid!")
	} else if toAccount, err := strconv.ParseFloat(toAccountqs, 64); err != nil {
		fmt.Fprintf(writer, "Account number is invalid!")
	} else {
		if accountFrom, ok := accounts[number]; !ok {
			fmt.Fprintf(writer, "Account number is invalid!")
		} else if accountTo, ok := accounts[toAccount]; !ok {
			fmt.Fprintf(writer, "Account number is invalid!")
		} else {
			err := accountFrom.Transfer(amount, accountTo.Account)
			if err != nil {
				fmt.Fprintf(writer, "transfer is invalid!")
			} else {
				fmt.Fprintf(writer, accountFrom.Statement())
			}
		}
	}
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
			json.NewEncoder(w).Encode(bankcore.Statement(account))
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


type CustomAccount struct {
	*bankcore.Account
}

func (c *CustomAccount) Statement() string {
	jsonObj, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}

	return string(jsonObj)
}