package main

import (
	"bank-system/banksystem"
	"bank-system/fixeddeposit"
	"bank-system/useraccount"
	"fmt"
)

func main() {
	bankSystem := banksystem.NewBankSystem()

	for {
		fmt.Println("Menu Options:")
		fmt.Println("1. Add User Account")
		fmt.Println("2. Open Fixed Deposit")
		fmt.Println("3. Display User Accounts")
		fmt.Println("4. Display Fixed Deposits")
		fmt.Println("5. Exit")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			var mobileNumber, name string
			var age int
			fmt.Print("Enter Mobile Number: ")
			fmt.Scan(&mobileNumber)
			fmt.Print("Enter Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Age: ")
			fmt.Scan(&age)

			if age <= 1 {
				fmt.Println("Person opening account must be older than 1 year.")
				break
			}
			user := useraccount.NewUserAccount(mobileNumber, name, age)
			bankSystem.AddUser(user)

		case 2:
			var mobileNumber string
			fmt.Print("Enter Mobile Number: ")
			fmt.Scan(&mobileNumber)

			accNum := useraccount.GenerateAccountNumber(mobileNumber)
			if user, exists := bankSystem.Users[accNum]; exists {
				var amount int64
				var duration int16
				fmt.Print("Enter FD Amount: ")
				fmt.Scan(&amount)
				fmt.Print("Enter Duration (years): ")
				fmt.Scan(&duration)

				if amount <= 0 || duration <= 0 {
					fmt.Println("Invalid amount or duration. Please enter positive values.")
					break
				}

				fd := fixeddeposit.NewFixedDeposit(user, amount, duration)
				bankSystem.AddFixedDeposit(fd)
				fmt.Println("Fixed Deposit created successfully.")
			} else {
				fmt.Println("You don't have an existing account! Please provide details to open FD.")
				var name string
				var age int
				fmt.Print("Enter Name: ")
				fmt.Scan(&name)
				fmt.Print("Enter Age: ")
				fmt.Scan(&age)

				if age <= 1 {
					fmt.Println("Person opening account must be older than 1 year.")
					break
				}
				user := useraccount.NewUserAccount(mobileNumber, name, age)
				bankSystem.AddUser(user)

				var amount int64
				var duration int16
				fmt.Print("Enter FD Amount: ")
				fmt.Scan(&amount)
				fmt.Print("Enter Duration (years): ")
				fmt.Scan(&duration)

				if amount <= 0 || duration <= 0 {
					fmt.Println("Invalid amount or duration. Please enter positive values.")
					break
				}

				fd := fixeddeposit.NewFixedDeposit(user, amount, duration)
				bankSystem.AddFixedDeposit(fd)
				fmt.Println("Fixed Deposit created successfully.")
			}

		case 3:
			bankSystem.DisplayUsers()

		case 4:
			bankSystem.DisplayFixedDeposits()

		case 5:
			fmt.Println("Exiting...")
			return
		}
		fmt.Println()
	}
}
