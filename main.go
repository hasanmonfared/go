package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}
type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var (
	userStorage       []User
	AuthenticateUser  *User
	taskStorage       []Task
	CategoryStorage   []Category
	serializationMode string
)

const (
	userStoragePath               = "user.txt"
	ManDarAvardiSerializationMode = "mandaravardi"
	JsonSerializationMode         = "json"
)

func main() {

	serializeMode := flag.String("serialize-mode", ManDarAvardiSerializationMode, "serialization code for store file")
	command := flag.String("command", "no command", "create a new task")
	flag.Parse()
	loadUserStorageFromFile(*serializeMode)
	fmt.Println("Hello to TODO app")

	switch *serializeMode {
	case ManDarAvardiSerializationMode:
		serializationMode = ManDarAvardiSerializationMode
	default:
		serializationMode = JsonSerializationMode
	}

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

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()
	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer %v\n\n", err)
		return
	}
	isFound := false
	for _, c := range CategoryStorage {
		if c.ID == categoryID && c.UserID == AuthenticateUser.ID {
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Println("category id is not found.\n")
		return
	}
	fmt.Println("please enter the task du date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     AuthenticateUser.ID,
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
	category := Category{
		ID:     len(CategoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: AuthenticateUser.ID,
	}
	CategoryStorage = append(CategoryStorage, category)
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
		Password: hashThePassword(password),
	}
	userStorage = append(userStorage, user)
	writeUserToFile(user)
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
func listTask() {
	for _, task := range taskStorage {
		if task.UserID == AuthenticateUser.ID {
			fmt.Println(task)
		}
	}
}
func loadUserStorageFromFile(serializationMode string) {
	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Println("can't open the file.")
	}
	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)
		return
	}
	var dataStr = string(data)
	userSlice := strings.Split(dataStr, "\n")
	for _, u := range userSlice {
		var userStruct = User{}

		switch serializationMode {
		case ManDarAvardiSerializationMode:
			var dErr error
			userStruct, dErr = deSerializeFromManDarAvardi(u)
			if dErr != nil {
				fmt.Println("can't deserialize user record to user struct", dErr)
				return
			}
		case JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("can't deserialize user record to user struct from json mode", uErr)
				return
			}
		default:
			fmt.Println("invalid serialization mode")
		}

		userStorage = append(userStorage, userStruct)
	}
}
func writeUserToFile(user User) {
	var file *os.File
	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("can't create or open file.", err)
		return
	}
	defer file.Close()

	var data []byte

	if serializationMode == ManDarAvardiSerializationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, password: %s\n", user.ID, user.Name, user.Email, user.Password))
	} else if serializationMode == JsonSerializationMode {
		var jErr error
		data, jErr = json.Marshal(user)
		if jErr != nil {
			fmt.Println("can't marshal user to json", jErr)
			return
		}
		data = append(data, []byte("\n")...)
	} else {
		fmt.Println("invalid serialization mode")
		return
	}

	file.Write(data)
}
func deSerializeFromManDarAvardi(userStr string) (User, error) {
	if userStr == "" {
		return User{}, errors.New("user string is empty")
	}
	var user = User{}
	userFields := strings.Split(userStr, ",")
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Println("field is not valid, skipping...", len(values))
			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]
		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("strconv error", err)
				return User{}, errors.New("strconv error")
			}
			user.ID = id
		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}
	}
	fmt.Printf("user: %+v\n", user)
	return user, nil
}
func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
