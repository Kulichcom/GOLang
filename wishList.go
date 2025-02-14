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

func (thing *Thing) setThingName(thingName string) {
	thing.thingName = thingName
}

func (thing *Thing) setThingLink(thingLink string) {
	thing.thingLink = thingLink
}

func (thing *Thing) setThingCost(thingCost int) {
	thing.thingCost = thingCost
}

// function for create new thing
func (thing *Thing) createThing(name, link string, cost int) {
	thing.thingName = name
	thing.thingLink = link
	thing.thingCost = cost
}

type WishList struct {
	listName  string
	comment   string
	listItems []Thing
}

func (wishList *WishList) getWishtListName() string {
	return wishList.listName
}

func (wishList *WishList) getWishListComment() string {
	return wishList.comment
}

// function for list items in wish list
func (wishList *WishList) showListItems() {
	// if list is empty, print a message that list is empty
	if len(wishList.listItems) == 0 {
		fmt.Printf("List: %s is empty\n", wishList.listName)
		return
	}

	for _, thing := range wishList.listItems {
		fmt.Printf("LIST NAME =  %s [ NAME: %s | LINK: %s | COST: %d ] COMMENT: %s\n", wishList.listName, thing.thingName, thing.thingLink, thing.thingCost, wishList.comment)
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

	for _, list := range account.Lists {
		if list.listName == listName {
			fmt.Print("This list already exists\n")
			return
		}
	}

	fmt.Print("Enter wish list comment: ")
	var wishListComment string
	fmt.Scanln(&wishListComment)

	newList := WishList{listName: listName, comment: wishListComment, listItems: []Thing{}}
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

	var listIndex = -1
	for i, list := range account.Lists {
		if list.listName == listName {
			listIndex = i
			break
		}
	}
	if listIndex == -1 {
		fmt.Print("No such list found")
		return
	}
	fmt.Print("Enter new thing name: ")
	fmt.Scanln(&newName)
	fmt.Print("Enter new thing link: ")
	fmt.Scanln(&newLink)
	fmt.Print("Enter new thing cost: ")
	fmt.Scanln(&newCost)

	for _, thing := range account.Lists[listIndex].listItems {
		if thing.thingName == newName {
			fmt.Print("This thing is already in the wish list\n")
			return
		}
	}

	newThing := Thing{thingName: newName, thingLink: newLink, thingCost: newCost}

	account.Lists[listIndex].listItems = append(account.Lists[listIndex].listItems, newThing)
}

// function to delete thing from wish list
func (account *Account) deleteThingFromWishList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists")
		return
	}

	fmt.Print("Enter wish list name: ")
	var listname string
	fmt.Scanln(&listname)

	var listIndex = -1
	for i, list := range account.Lists {
		if list.listName == listname {
			listIndex = i
			break
		}
	}

	if listIndex == -1 {
		fmt.Print("No such wish list found")
		return
	}

	if len(account.Lists[listIndex].listItems) == 0 {
		fmt.Print("Thing list is empty")
		return
	}

	fmt.Print("Enter name of the thing to delte: ")
	var itemName string
	fmt.Scanln(&itemName)

	var itemIndex = -1
	for i, thing := range account.Lists[listIndex].listItems {
		if thing.thingName == itemName {
			itemIndex = i
		}
	}

	if itemIndex == -1 {
		fmt.Print("No such thing in the list")
		return
	}

	account.Lists[listIndex].listItems = append(account.Lists[listIndex].listItems[:itemIndex], account.Lists[listIndex].listItems[itemIndex+1:]...)
	fmt.Printf("ITEM: %s", itemName, " was deleted from the list: %s\n", listname)
}

// function for delte wish list
func (account *Account) deleteWishList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists\n")
		return
	}

	fmt.Print("Enter wish list name: ")
	var listname string
	fmt.Scanln(&listname)

	var listIndex = -1
	for i, list := range account.Lists {
		if list.listName == listname {
			listIndex = i
			break
		}
	}

	if listIndex == -1 {
		fmt.Print("No such list found\n")
		return
	}

	account.Lists = append(account.Lists[:listIndex], account.Lists[listIndex+1:]...)
	fmt.Printf("WISH LIST: %s", listname, " was deleted\n")
}

func (account *Account) changeThingInWishList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists\n")
		return
	}

	fmt.Print("Enter wish list name: ")
	var listName string
	fmt.Scanln(&listName)

	var listFound = -1
	for i, list := range account.Lists {
		if list.listName == listName {
			listFound = i
			break
		}
	}

	if listFound == -1 {
		fmt.Print("No such list found\n")
		return
	}

	fmt.Print("Enter thing name wish you want to change: ")
	var thingName string
	fmt.Scanln(&thingName)

	var thingFound = -1
	for i, thing := range account.Lists[listFound].listItems {
		if thing.thingName == thingName {
			thingFound = i
			break
		}
	}

	if thingFound == -1 {
		fmt.Print("No such thing found\n")
		return
	}

	fmt.Print("Which info you want to cahgne?\n")
	fmt.Print("1. Thing name\n")
	fmt.Print("2. Thing link\n")
	fmt.Print("3. Thing cost\n")

	var answer int
	fmt.Scanln(&answer)

	switch answer {
	case 1:
		fmt.Print("Enter new thing name: ")
		var newThingName string
		fmt.Scanln(&newThingName)

		account.Lists[listFound].listItems[thingFound].setThingName(newThingName)
		fmt.Printf("Name of: %s was changed to: %s\n", account.Lists[listFound].listItems[thingFound].thingName, newThingName)

	case 2:
		fmt.Print("Enter new thing link: ")
		var newThingLink string
		fmt.Scanln(&newThingLink)

		account.Lists[listFound].listItems[thingFound].setThingLink(newThingLink)
		fmt.Printf("Link of: %s was changed to: %s\n", account.Lists[listFound].listItems[thingFound].thingName, newThingLink)

	case 3:
		fmt.Print("Enter new thing cost: ")
		var newThingCost int
		fmt.Scanln(&newThingCost)

		account.Lists[listFound].listItems[thingFound].setThingCost(newThingCost)
		fmt.Printf("Cost of: %s was changed to: %d\n", account.Lists[listFound].listItems[thingFound].thingName, newThingCost)

	default:
		fmt.Print("Error")
	}
}

// function to show all wish lists
func (account *Account) showAllWishList() {
	if len(account.Lists) == 0 {
		fmt.Print("You don't have any wish lists\n")
		return
	}

	for _, list := range account.Lists {
		list.showListItems()
	}
}

// functiion for show User info
func (account *Account) showAccountInfo() {
	fmt.Printf("User name: %s\n", account.userName)
	fmt.Printf("Age: %d\n", account.userAge)
	fmt.Printf("Birthday: %s\n", account.Bdate)
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
		fmt.Print("4. Delete thing from wish list\n")
		fmt.Print("5. Delete wish list\n")
		fmt.Print("6. Show account innfo\n")
		fmt.Print("7. Change thing in wish list\n")
		fmt.Print("8. Exit\n")

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
			account.deleteThingFromWishList()
		case 5:
			account.deleteWishList()
		case 6:
			account.showAccountInfo()
		case 7:
			account.changeThingInWishList()
		case 8:
			run = false
		default:
			fmt.Print("Error")
		}
	}

	fmt.Print("Exit")
}
