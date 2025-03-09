package main

import (
	"errors"
	"fmt"
	// "errors"
)

// создаем счет
type Account struct {
	balance float64 // переменная balance будет с плавающей точнкой
}	

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Пополнение счета на: %.2f. ", amount)
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("Недостаточно средств для снятия")
	}

	a.balance -= amount
	fmt.Printf("Снятие %.2f.", amount)
	return nil
}

func (a *Account) Balance() float64{
	return a.balance
}

func main() {
	account := Account{balance: 100.0}
	fmt.Printf("Счет создан. Текущий баланс: %.2f.\n", account.Balance())

	account.Deposit(50.0)
	fmt.Printf("Текущий баланс: %.2f.\n", account.Balance())

	err := account.Withdraw(20.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Текущий баланс: %.2f\n", account.Balance())
	}

	err = account.Withdraw(100.0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Текущий баланс: %.2f\n", account.Balance())
	}

}
