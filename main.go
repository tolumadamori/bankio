package main

import (
	"fmt"
	"os"
)

type user struct {
	username string
	password int
	balance  float32
}

var defaultUser = user{
	"tolu",
	1234,
	1000.00,
}

func menu() {
	fmt.Printf("\n\n")
	fmt.Println("Please select an option:........\n #1 -Change Pin \t \t #2 - Show Account Balance \n #3 - Withdraw \t \t #4 - Deposit \n \t #5 - Exit program")
	fmt.Println("\n enter the option number and press enter:")
	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		changePin()
	case 2:
		accountBalance()
	case 3:
		withdrawFunds()

	case 4:
		depositFunds()

	case 5:
		exit()

	default:
		fmt.Printf("\n\n")
		fmt.Println("The option you selected is invalid. Please try again!")
		menu()

	}

}

// Collect user input and store as username and password.
func credentials() (username string, password int) {
	fmt.Println("Type in the username and press enter") //collect username from user input
	fmt.Scanln(&username)
	fmt.Println("Type in your current password and press enter") //collect password from user input.
	fmt.Scanln(&password)

	return username, password
}

// funtion to change pin. Takes the first password and the
func changePin() {
	username, password := credentials()
	if username == defaultUser.username {
		if password == defaultUser.password {
			fmt.Println("Input the Newpassword and press enter. Please note, Digits only")
			var newPassword int
			if _, err := fmt.Scanln(&newPassword); err != nil {
				fmt.Printf("There was an error: %v", err)
			} else {
				defaultUser.password = newPassword
				fmt.Printf("\n\n")
				fmt.Println("Password Changed! Going back to main menu...")
				menu()
			}

		} else {
			fmt.Println("passwords do not match. Returning to main menu")
			menu()
		}
	} else {
		fmt.Println("Username does not match. Returning to main menu")
		menu()
	}

}

// takes the user name and password and returns the balance if credentials match
func accountBalance() {
	username, password := credentials()
	if username == defaultUser.username {
		if password == defaultUser.password {
			fmt.Printf("\n\n")
			fmt.Printf("\n %v your balance is: %v\n", defaultUser.username, defaultUser.balance)
			fmt.Printf("\n\n")
			fmt.Println("press enter to return to main menu")
			fmt.Scanln()
			menu()
		} else {
			fmt.Printf("\n\n")
			fmt.Println("passwords do not match. Returning to main menu")
			menu()
		}
	} else {
		fmt.Printf("\n\n")
		fmt.Println("Username does not match. Returning to main menu")
		menu()
	}

}

// Removes the amount entered from the user's balance
func withdrawFunds() {
	username, password := credentials()
	if username == defaultUser.username {
		if password == defaultUser.password {
			var amount float32
			fmt.Printf("Enter Amount to Withdraw and press enter")
			fmt.Printf("\n")
			if _, err := fmt.Scanln(&amount); err != nil {
				fmt.Printf("There was an error: %v. RETURNING TO MAIN MENU!", err)
				menu()
			} else if amount >= defaultUser.balance {
				fmt.Println("INSUFFICIENT BALANCE. RETURNING TO MAIN MENU!")
				menu()
			} else {
				defaultUser.balance = defaultUser.balance - amount
				fmt.Printf("%v has been withdrawn from your account. Current Balance is %v", amount, defaultUser.balance)
				fmt.Printf("\n\n")
				fmt.Println("press enter to return to main menu")
				fmt.Scanln()
				menu()
			}

		} else {
			fmt.Println("passwords do not match. Returning to main menu")
			menu()
		}
	} else {
		fmt.Println("Username does not match. Returning to main menu")
		menu()
	}

}

// Adds the amount entered to the users balance.
func depositFunds() {
	username, password := credentials()
	if username == defaultUser.username {
		if password == defaultUser.password {
			var amount float32
			fmt.Printf("Enter Amount to Deposit and press enter")
			fmt.Printf("\n")
			if _, err := fmt.Scanln(&amount); err != nil {
				fmt.Printf("There was an error: %v. Returning to Main Menu!", err)
				menu()
			} else {
				defaultUser.balance = defaultUser.balance + amount
				fmt.Printf("%v has been deposited to your account. Current Balance is %v", amount, defaultUser.balance)
				fmt.Printf("\n\n")
				fmt.Println("press enter to return to main menu")
				fmt.Scanln()
				menu()
			}

		} else {
			fmt.Println("passwords do not match. Returning to main menu")
			fmt.Printf("\n\n")
			menu()
		}
	} else {
		fmt.Println("Username does not match. Returning to main menu")
		fmt.Printf("\n\n")
		menu()
	}

}

// exits the CLI
func exit() {
	fmt.Printf("\n\n")
	fmt.Println("Exiting...")
	os.Exit(1)
}

// func welcome:
func welcome() {

	fmt.Printf("\n\n")
	fmt.Println("The default name and password are 'tolu' and '1234'. To validate you are a real user, type out the default password and press enter.")
	fmt.Printf("\n")
	fmt.Println("Note: If you do not want to go further, type in 99 and press enter!")
	var password int
	fmt.Scanln(&password)
	switch password {
	case 1234:
		menu()
	case 99:
		exit()
	default:
		welcome()
	}
}

func main() {
	fmt.Println("Starting...")
	fmt.Printf("\n")
	fmt.Println("Welcome to bankio...")
	fmt.Printf("\n")
	welcome()
	menu()

}
