package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposite(amount float64) {
	b.Balance += amount
	fmt.Println("amount:", amount, "balance:", b.Balance)
}
func (b BankAccount) GetBalance() float64 {
	return b.Balance

}
func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance > amount {
		b.Balance -= amount
		fmt.Println("amount:", amount, "balance:", b.Balance)
	} else {
		fmt.Println("Not enough money")
		return
	}

}
func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposite(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}
func main() {

	account := BankAccount{
		AccountNumber: "123456789",
		HolderName:    "Aday",
		Balance:       5000.0,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the banking system!")
	fmt.Println("Your current balance:", account.GetBalance())

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter the amount to deposit: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Try again.")
				continue
			}
			account.Deposite(amount)

		case "2":
			fmt.Print("Enter the amount to withdraw: ")
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Try again.")
				continue
			}
			account.Withdraw(amount)

		case "3":
			fmt.Printf("Your current balance: %.2f\n", account.GetBalance())

		case "4":
			fmt.Println("Thank you for using the banking system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please select an option from the menu.")
		}
	}
}
