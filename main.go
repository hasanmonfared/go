package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category string
	IsDone   bool
	UserID   int
}

var userStorage []User
var AuthenticateUser *User
var taskStorage []Task

func main() {
	fmt.Println("Hello to TODO app")
	command := flag.String("command", "no command", "create a new task")
	flag.Parse()

	for {
		runCommand(*command)
		fmt.Println("please enter another command")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}
func runCommand(command string) {

	if command != "register-user" && command != "exit" && AuthenticateUser == nil {
		login()
		if AuthenticateUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "list-task":
		listTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "exist":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}
func createTask() {
	var title, duedate, category string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task du date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		DueDate:  duedate,
		Category: category,
		IsDone:   false,
		UserID:   AuthenticateUser.ID,
	}
	taskStorage = append(taskStorage, task)
	fmt.Println("task:", title, category, duedate)
}
func createCategory() {
	var title, color string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()
	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category", title, color)
}
func registerUser() {
	var id, email, name, password string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()
	id = email

	fmt.Println("user", id, email, password)
	user := User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)
}
func login() {
	fmt.Println("login process")
	var email, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			AuthenticateUser = &user

			break
		} else {
			fmt.Println("The email and password not correct.")
		}
	}
	if AuthenticateUser == nil {
		fmt.Println("the email or password is not correct.")
		return
	}

	fmt.Println("user", email, password)
}
func listTask() {
	for _, task := range taskStorage {
		if task.UserID == AuthenticateUser.ID {
			fmt.Println(task)
		}
	}
}
