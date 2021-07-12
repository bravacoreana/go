package main

import (
	"fmt"

	"github.com/lihakim/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bilbo")
	fmt.Println("account", account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}
