package main

import (
	"fmt"
)

var userIDCounter int
var taskIDCounter int

func genUserID() int {
	userIDCounter++
	return userIDCounter
}

func genTaskID() int {
	taskIDCounter++
	return taskIDCounter
}

type Task struct {
	taskName   string
	taskDate   string
	taskID     int
	taskStatus bool
}

func (task *Task) getTaskName() string {
	return task.taskName
}

func (task *Task) getTaskDate() string {
	return task.taskDate
}

func (task *Task) getTaskID() int {
	return task.taskID
}

func (task *Task) getTaskStatus() bool {
	return task.taskStatus
}

func (task *Task) setTaskName(taskName string) {
	task.taskName = taskName
}

func (task *Task) setTaskDate(taskDate string) {
	task.taskDate = taskDate
}

func (task *Task) setTaskStatus(taskStatus bool) {
	task.taskStatus = taskStatus
}

type TaskList struct {
	listCategory string
	listName     string
	list         []Task
}

func (list *TaskList) getListName() string {
	return list.listName
}

func (list *TaskList) getListCategory() string {
	return list.listCategory
}

func (list *TaskList) showTaskList() {
	if len(list.list) == 0 {
		fmt.Printf("List: %s is empty\n", list.listName)
	}

	for _, task := range list.list {
		fmt.Printf("ID: %d | TASK NAME: %s | TASK DATE: %s | TASK STATUS: %t\n", task.getTaskID(), task.getTaskName(), task.getTaskDate(), task.getTaskStatus())
	}
}

type User struct {
	userName    string
	userSurname string
	password    string
	userID      int
	userLists   []TaskList
}

func (user *User) getUserName() string {
	return user.userName
}

func (user *User) getUserSurname() string {
	return user.userSurname
}

func (user *User) getUserID() int {
	return user.userID
}

func (user *User) setPassword(pass string) {
	user.password = pass
}

func (user *User) createNewTaskList() {
	fmt.Print("Enter new list name: ")
	var listName string
	fmt.Scanln(&listName)

	for _, i := range user.userLists {
		if i.listName == listName {
			fmt.Print("This list is already exists\n")
			return
		}
	}

	fmt.Print("Enter list category: [WORK, PERSONAL, STUDY]: ")
	var listCategory string
	fmt.Scanln(&listCategory)

	newList := TaskList{listName: listName, listCategory: listCategory, list: []Task{}}
	user.userLists = append(user.userLists, newList)
	fmt.Printf("List: %s, was created\n", listName)
}

func (user *User) showAllLists() {
	if len(user.userLists) == 0 {
		fmt.Print("You don't have any task lists\n")
		return
	}

	for _, list := range user.userLists {
		list.showTaskList()
	}
}

func (user *User) addNewTask() {
	if len(user.userLists) == 0 {
		fmt.Print("You don't have any lists\n")
		return
	}

	fmt.Print("Enter list name: ")
	var listName string
	fmt.Scanln(&listName)

	var found = -1
	for i, list := range user.userLists {
		if list.listName == listName {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Printf("List: %s not found", listName)
		return
	}

	fmt.Print("Enter new task name: ")
	var newTaskName string
	fmt.Scanln(&newTaskName)

	fmt.Print("Enter new task date: ")
	var newTaskDate string
	fmt.Scanln(&newTaskDate)

	newTask := Task{taskName: newTaskName, taskDate: newTaskDate, taskID: genTaskID()}
	user.userLists[found].list = append(user.userLists[found].list, newTask)
	fmt.Printf("Task: %s was added to list: %s\n", newTaskName, listName)
}

func (user *User) setAsComplete() {
	if len(user.userLists) == 0 {
		fmt.Print("You don't have any lists\n")
		return
	}

	fmt.Print("Enter task name: ")
	var name string
	fmt.Scanln(&name)

	var listIndex = -1
	var taskIndex = -1
	for i := 0; i < len(user.userLists); i++ {
		for j := 0; j < len(user.userLists[i].list); j++ {
			if user.userLists[i].list[j].taskName == name {
				listIndex = i
				taskIndex = j
				break
			}
		}
	}

	if listIndex == -1 {
		fmt.Print("No list found\n")
		return
	}
	if taskIndex == -1 {
		fmt.Print("No task found\n")
		return
	}

	user.userLists[listIndex].list[taskIndex].setTaskStatus(true)
	fmt.Printf("Task: %s statuc set to TRUE\n", name)
}

func main() {
	fmt.Print("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter your surname: ")
	var surname string
	fmt.Scanln(&surname)

	fmt.Print("Set your password: ")
	var password string
	fmt.Scanln(&password)

	user := User{userName: name, userSurname: surname, password: password, userID: genUserID()}

	run := true

	for run {
		fmt.Print("MENUE\n")
		fmt.Print("1. Create new task list\n")
		fmt.Print("2. Add new task to task list\n")
		fmt.Print("3. Show all task lists\n")
		fmt.Print("4. Mark task as complete\n")
		fmt.Print("5. Exit\n")

		fmt.Print("Enter your choose: ")
		var answer int
		fmt.Scanln(&answer)

		switch answer {
		case 1:
			user.createNewTaskList()
		case 2:
			user.addNewTask()
		case 3:
			user.showAllLists()
		case 4:
			user.setAsComplete()
		case 5:
			run = false
		default:
			fmt.Print("Error\n")
		}
	}
	fmt.Print("Exit")
}
