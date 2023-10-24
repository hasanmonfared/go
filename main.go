package main

import (
	"fmt"
	"interface_test/log"
	"interface_test/richerror"
	"interface_test/simpleerror"
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
	logger := log.Log{}

	_, oErr := os.OpenFile("storage/data.txt", os.O_RDWR, 0777)
	if oErr != nil {

		logger.Append(oErr)
		fmt.Println(oErr.Error())
	}
	_, gErr := getUserByID(0)
	if gErr != nil {
		logger.Append(gErr)
	}
	_, g2Err := getUserByIDTwo(0)
	if g2Err != nil {
		logger.Append(g2Err)
	}
	_, g3Err := getUserByIDThree(0)
	if g3Err != nil {
		logger.Append(g3Err)
	}
	logger.Save()
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
func getUserByIDTwo(id int) (User, error) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByIDTwo",
		}
	}
	return User{}, nil
}

func getUserByIDThree(id int) (User, error) {
	if id == 0 {
		return User{}, &simpleerror.SimpleError{
			Output:    "id is 0",
			Operation: "getUserByIDThree",
		}
	}
	return User{}, nil
}
