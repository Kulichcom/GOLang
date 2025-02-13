package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name    string
	surname string
	age     int
}

func (person *Person) setName(personName string) {
	person.name = personName
}

func (person *Person) getName() string {
	return person.name
}

func (person *Person) setSurname(personSurname string) {
	person.surname = personSurname
}

func (person *Person) getSurname() string {
	return person.surname
}

func (person *Person) setAge(personAge int) {
	person.age = personAge
}

func (person *Person) getAge() int {
	return person.age
}

type Account struct {
	ownerInfo Person
	balance   int
	ID        int
}

var (
	accountIdCounter int
	mu               sync.Mutex
)

func generateId() int {
	mu.Lock()
	defer mu.Unlock()
	accountIdCounter++
	return accountIdCounter
}

func (account *Account) addMoney(money int) {
	account.balance += money
	fmt.Printf("New Balance: %d\n", account.balance)
}

func (account *Account) transferMoney(money int) {
	if money > account.balance {
		fmt.Println("Not enough money")
		return
	}
	account.balance -= money
	fmt.Printf("New Balance: %d\n", account.balance)
}

func (account *Account) getBalance() int {
	return account.balance
}

func (account *Account) setBalance(amount int) {
	account.balance = amount
}

func (account *Account) getOwnerInfo() *Person {
	return &account.ownerInfo
}

func (account *Account) getAccountID() int {
	return account.ID
}

func main() {
	var name, surname string
	var age, amountMoney int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your surname: ")
	fmt.Scanln(&surname)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("You cannot create an account because you are under 18.")
		return
	}

	person := Person{name: name, surname: surname, age: age}

	fmt.Print("How much money do you have? ")
	fmt.Scanln(&amountMoney)

	account := Account{ownerInfo: person, balance: amountMoney, ID: generateId()}

	run := true

	for run {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add money")
		fmt.Println("2. Transfer money")
		fmt.Println("3. Show account info")
		fmt.Println("4. Exit")

		var answer int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&answer)

		switch answer {
		case 1:
			fmt.Print("How much money do you want to add? ")
			var amount int
			fmt.Scanln(&amount)
			account.addMoney(amount)
		case 2:
			fmt.Print("How much money do you want to transfer? ")
			var amount int
			fmt.Scanln(&amount)
			account.transferMoney(amount)
		case 3:
			fmt.Printf("Owner Info: %+v\n", account.getOwnerInfo())
			fmt.Printf("Balance: %d\n", account.getBalance())
			fmt.Printf("Account ID: %d\n", account.getAccountID())
		case 4:
			run = false
		default:
			fmt.Println("Incorrect option, please try again.")
		}
	}
}
