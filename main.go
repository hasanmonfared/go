package main

import (
	"interface_test/app"
	"interface_test/storage"
	"interface_test/user"
)

func main() {
	application := app.App{
		Name: "sample-app",
		//StorageFilePath: "./data.txt",
		UserStorage: &storage.Memory{},
	}
	u := user.User{
		ID:   1,
		Name: "Hassan",
	}
	application.CreateUser(u)
}
