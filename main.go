package main

import (
	"app/constant"
	"app/contract"
	"app/entity"
	"app/repository/filestore"
	"app/repository/memorystore"
	"app/service/task"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	userStoragePath = "user.txt"
)

var (
	userStorage       []entity.User
	AuthenticateUser  *entity.User
	CategoryStorage   []entity.Category
	serializationMode string
)

func main() {
	taskMemoryRepo := memorystore.NewTaskStore()
	taskService := task.NewService(taskMemoryRepo)

	serializeMode := flag.String("serialize-mode", constant.ManDarAvardiSerializationMode, "serialization code for store file")
	command := flag.String("command", "no command", "create a new task")
	flag.Parse()

	fmt.Println("Hello to TODO app")

	switch *serializeMode {
	case constant.ManDarAvardiSerializationMode:
		serializationMode = constant.ManDarAvardiSerializationMode
	default:
		serializationMode = constant.JsonSerializationMode
	}
	var userFileStore = filestore.New(userStoragePath, serializationMode)
	users := userFileStore.Load()
	userStorage = append(userStorage, users...)
	for {
		runCommand(userFileStore, *command, &taskService)
		fmt.Println("please enter another command")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}

}
func runCommand(userFileStore contract.UserWriteStore, command string, taskService *task.Service) {

	if command != "register-user" && command != "exit" && AuthenticateUser == nil {
		login()
		if AuthenticateUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask(taskService)
	case "list-task":
		listTask(taskService)
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(userFileStore)
	case "exist":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}
func createTask(taskService *task.Service) {
	var title, duedate, category string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()
	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer %v\n\n", err)
		return
	}

	fmt.Println("please enter the task du date")
	scanner.Scan()
	duedate = scanner.Text()
	response, err := taskService.Create(task.CreateRequest{
		Title:               title,
		DueDate:             duedate,
		CategoryID:          categoryID,
		AuthenticatedUserID: AuthenticateUser.ID,
	})
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("create task :", response.Task)
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
	category := entity.Category{
		ID:     len(CategoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: AuthenticateUser.ID,
	}
	CategoryStorage = append(CategoryStorage, category)
}

func registerUser(store contract.UserWriteStore) {
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
	user := entity.User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}
	userStorage = append(userStorage, user)
	store.Save(user)
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
		if user.Email == email && user.Password == hashThePassword(password) {
			AuthenticateUser = &user

			break
		} else {
			fmt.Println("The email and password not correct.")
		}
	}

	fmt.Println("user", email, password)
}
func listTask(taskService *task.Service) {
	userTask, err := taskService.List(task.ListRequest{UserID: AuthenticateUser.ID})
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("user tasks", userTask.Tasks)
}

func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
