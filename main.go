package main

import (
	"fmt"
	"interface_test/log"
	"interface_test/richerror"
	"os"
	"strconv"
)

type User struct {
	ID   uint
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id:%d, name: %s}", u.ID, u.Name)
}
func main() {
	//u := User{
	//	ID:   123,
	//	Name: "Hassan",
	//}
	//fmt.Println(u)
	logger := log.Log{}

	f, oErr := os.OpenFile("storage/data.txt", os.O_RDWR, 0777)
	if oErr != nil {
		logger.Append(oErr)

	}
	user, gErr := getUserByID(0)
	if gErr != nil {
		logger.Append(gErr)

	}
	fmt.Println("user", user)
}

func getUserByID(id int) (User, error) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByID",
		}
	}
	return User{}, nil
}
