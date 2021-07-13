package main

import (
	"fmt"

	"github.com/lihakim/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("bilbo")
	account.Deposit(10)
	fmt.Println(account)
}
