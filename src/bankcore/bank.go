package bankcore

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func Hello() string {
	return "hey! im working"
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if amount > a.Balance {
		return errors.New("amount must be less than the balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// TODO 匿名字段好像可以继承方法，但不支持函数
// 匿名嵌入允许你访问嵌入类型的方法和字段，但并不会改变类型
// 和java中的继承相似，但不是继承
func (a *Account) Transfer(amount float64, toAccountPtr *Account) error {
	if amount < 0 {
		return errors.New("amount must be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("balance must be greater than zero")
	}
	err := a.Withdraw(amount)
	if err != nil {
		return err
	}
	err = toAccountPtr.Deposit(amount)
	if err != nil {
		return err
	}
	return nil
}

// Bank /**
type Bank interface {
	Statement() string
}

// Statement 接收接口作为函数/**
func Statement(b Bank) string {
	return b.Statement()
}


