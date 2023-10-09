package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	id       int
	email    string
	password string
}

var userStorage []User

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
	fmt.Printf("userStorage %v\n", userStorage)
}
func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exist":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}
func createTask() {
	var name, duedate, category string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("please enter the task du date")
	scanner.Scan()
	duedate = scanner.Text()
	fmt.Println("task:", name, category, duedate)
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
	var id, email, password string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()
	id = email

	fmt.Println("user", id, email, password)
	user := User{
		id:       len(userStorage) + 1,
		email:    email,
		password: password,
	}
	userStorage = append(userStorage, user)
}
func login() {
	var email, password string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()
	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()
	fmt.Println("user", email, password)
}
