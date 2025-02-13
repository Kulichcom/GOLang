package main

import (
	"fmt"
	"sync"
)

type Thing struct {
	thingName string
	thingLink string
	thingCost int
}

func (thing *Thing) getThingName() string {
	return thing.thingName
}

func (thing *Thing) getThingLink() string {
	return thing.thingLink
}

func (thing *Thing) getThingCost() int {
	return thing.thingCost
}

// function for create new thing
func (thing *Thing) createThing(name, link string, cost int) {
	thing.thingName = name
	thing.thingLink = link
	thing.thingCost = cost
}

type WishList struct {
	listName  string
	listItems []Thing
}

func (wishList *WishList) getWishtListName() string {
	return wishList.listName
}

// function for list items in wish list
func (wishList *WishList) showListItems() {
	// if list is empty, print a message that list is empty
	if len(wishList.listItems) == 0 {
		fmt.Print("List is empty")
		return
	}

	for _, thing := range wishList.listItems {
		fmt.Printf("NAME: %s", thing.thingName)
		fmt.Printf("LINK: %s", thing.thingLink)
		fmt.Printf("COSR: %d", thing.thingCost)
	}
}

type Account struct {
	// User's info
	userName string
	userAge  int
	Bdate    string
	ID       int

	// All wishlists
	Lists []WishList
}

// START: Block for creating account ID
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

//END

// function for create new wish list
func (account *Account) createWishList() {
	fmt.Print("Enter wish list name: ")
	var listName string
	fmt.Scanln(&listName)

	newList := WishList{listName: listName, listItems: []Thing{}}
	account.Lists = append(account.Lists, newList)
}

// function for add new thing to wish list
func (account *Account) addThingToWhisList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists")
		return
	}

	var newName string
	var newLink string
	var newCost int

	fmt.Print("Enter name of the wish list: ")
	var listName string
	fmt.Scanln(&listName)

	for i := 0; i < len(account.Lists); i++ {
		if account.Lists[i].listName == listName {
			fmt.Print("Enter thing name: ")
			fmt.Scanln(&newName)

			fmt.Print("Enter thing link: ")
			fmt.Scanln(&newLink)

			fmt.Print("Enter thing cost: ")
			fmt.Scanln(&newCost)
		} else {
			fmt.Print("No such list")
			return
		}
	}

	newThing := Thing{thingName: newName, thingLink: newLink, thingCost: newCost}

	for i := 0; i < len(account.Lists); i++ {
		if account.Lists[i].listName == listName {
			account.Lists[i].listItems = append(account.Lists[i].listItems, newThing)
		}
	}
}

func (account *Account) showAllWishList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists")
		return
	}

	for _, list := range account.Lists {
		list.showListItems()
	}
}

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)

	fmt.Print("Enter your birthday: ")
	var bdate string
	fmt.Scanln(&bdate)

	account := Account{userName: name, userAge: age, Bdate: bdate, ID: generateId()}

	run := true

	for run {
		fmt.Print("Menu\n")
		fmt.Print("1. Create wish list\n")
		fmt.Print("2. Add thing to wish list\n")
		fmt.Print("3. Show all wish lists\n")
		fmt.Print("4. Exit\n")

		var answer int
		fmt.Print("Enter your answer: ")
		fmt.Scanln(&answer)

		switch answer {
		case 1:
			account.createWishList()
		case 2:
			account.addThingToWhisList()
		case 3:
			account.showAllWishList()
		case 4:
			run = false
		default:
			fmt.Print("Error")
		}
	}

	fmt.Print("Exit")
}
