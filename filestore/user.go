package filestore

import (
	"app/constant"
	"app/entity"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileStore struct {
	filePath          string
	serializationMode string
}

func New(path string, serializationMode string) FileStore {
	return FileStore{
		filePath:          path,
		serializationMode: serializationMode,
	}
}
func (f FileStore) Save(u entity.User) {
	f.writeUserToFile(u)
}
func (f FileStore) Load() []entity.User {
	var uStore []entity.User

	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("can't open the file.")
	}
	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)
		return nil
	}
	var dataStr = string(data)
	userSlice := strings.Split(dataStr, "\n")
	for _, u := range userSlice {
		var userStruct = entity.User{}

		switch f.serializationMode {
		case constant.ManDarAvardiSerializationMode:
			var dErr error
			userStruct, dErr = deSerializeFromManDarAvardi(u)
			if dErr != nil {
				fmt.Println("can't deserialize user record to user struct", dErr)
				return nil
			}
		case constant.JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("can't deserialize user record to user struct from json mode", uErr)
				return nil
			}
		default:
			fmt.Println("invalid serialization mode")
		}

		uStore = append(uStore, userStruct)
	}
	return uStore
}
func (f FileStore) writeUserToFile(user entity.User) {
	var file *os.File
	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("can't create or open file.", err)
		return
	}
	defer file.Close()

	var data []byte

	if f.serializationMode == constant.ManDarAvardiSerializationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, password: %s\n", user.ID, user.Name, user.Email, user.Password))
	} else if f.serializationMode == constant.JsonSerializationMode {
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
func deSerializeFromManDarAvardi(userStr string) (entity.User, error) {
	if userStr == "" {
		return entity.User{}, errors.New("user string is empty")
	}
	var user = entity.User{}
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
				return entity.User{}, errors.New("strconv error")
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
