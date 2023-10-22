package app

import (
	"interface_test/user"
)

type userStore interface {
	CreateUser(u user.User)
	ListUsers() []user.User
	GetUserByID(id uint) user.User
}

type App struct {
	Name string
	//StorageFilePath string
	UserStorage userStore
}

func (a App) CreateUser(u user.User) {
	//if u.Name == "" {
	//	fmt.Println("name can't be empty")
	//
	//	return
	//}
	//var fileHandler *os.File
	//if f, err := os.OpenFile(a.StorageFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777); err != nil {
	//	fmt.Println("can't open file", err)
	//	return
	//} else {
	//	fileHandler = f
	//}
	//defer fileHandler.Close()
	//data, mErr := json.Marshal(u)
	//if mErr != nil {
	//	fmt.Println("can't marshal json", mErr)
	//
	//	return
	//}
	//
	//if _, wErr := fileHandler.Write(data); wErr != nil {
	//	fmt.Println("can't writ to the file.", mErr)
	//
	//	return
	//}
	a.UserStorage.CreateUser(u)
}
