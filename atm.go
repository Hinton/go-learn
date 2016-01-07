package main

import (
	"bufio"
	"fmt"
	"github.com/hinton/go-learn/bank"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)

func main() {

	account := bank.CreateAccount(1)
	account.Deposit(decimal.NewFromFloat(1000.50))

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Go ATM service!\n")

	for {

		fmt.Println("1. View current balance.")
		fmt.Println("2. Withdraw.")
		fmt.Println("3. Exit.")

		actionStr, _ := reader.ReadString('\n')
		action, _ := strconv.ParseInt(strings.TrimSpace(actionStr), 10, 32)

		switch action {
		case 1:
			fmt.Println(account.GetBalance().String())
		case 2:
			fmt.Println("\nWithdraw amount:")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := decimal.NewFromString(strings.TrimSpace(amountStr))

			_, err := account.Withdraw(amount)

			if err != nil {
				fmt.Println(err)
			}

		case 3:
			os.Exit(0)
		}

	}

}
