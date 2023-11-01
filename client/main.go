package main

import (
	"app/delivery/deliveryparam"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	message := "default message"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	connection, err := net.Dial("tcp", "127.0.1.1:8282")
	if err != nil {
		log.Fatal("can't dial given address.", err)
	}
	defer connection.Close()

	fmt.Println(connection.LocalAddr())
	req := deliveryparam.Request{Command: message}
	if req.Command == "create-task" {
		req.CreateTaskRequest = deliveryparam.CreateTaskRequest{
			Title:      "test",
			DueDate:    "test",
			CategoryID: 1,
		}
	}
	serializeData, mErr := json.Marshal(&req)
	if mErr != nil {
		log.Fatalln("can't write data to connection", mErr)
	}

	numberOfWriteBytes, rErr := connection.Write(serializeData)
	if rErr != nil {
		log.Fatalln("can't write data in connection")
	}
	fmt.Println("number of data", numberOfWriteBytes)

	var data = make([]byte, 1024)
	_, reErr := connection.Read(data)
	if reErr != nil {
		log.Fatalln("can't read data form", reErr)
	}
	fmt.Println("server response", string(data))
}
