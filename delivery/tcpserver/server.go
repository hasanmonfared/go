package main

import (
	"app/delivery/deliveryparam"
	"app/repository/memorystore"
	"app/service/task"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {

	const (
		network = "tcp"
		address = "127.0.1.1:8282"
	)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("can't listen on given address.", address, err)
	}
	fmt.Println("server listening on address:", listener.Addr())

	taskMemoryRepo := memorystore.NewTaskStore()
	taskService := task.NewService(taskMemoryRepo)

	for {
		connection, aErr := listener.Accept()
		if err != nil {
			log.Println("can't listen to new connection", aErr)
			continue
		}

		fmt.Println("connection address", connection.RemoteAddr(), connection.LocalAddr())
		var rawRequest = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(rawRequest)
		if rErr != nil {
			log.Println("can't read data form", rErr)
			continue
		}
		fmt.Println("number of read bytes", numberOfReadBytes)

		fmt.Printf("client address : %s,numberOfReadBytes: %d,data: %s\n",
			connection.RemoteAddr(), numberOfReadBytes, string(rawRequest))
		req := &deliveryparam.Request{}
		if uErr := json.Unmarshal(rawRequest[:numberOfReadBytes], req); uErr != nil {
			log.Println("bad request", uErr)
			continue
		}

		switch req.Command {
		case "create-task":
			response, cErr := taskService.Create(task.CreateRequest{
				Title:               req.CreateTaskRequest.Title,
				DueDate:             req.CreateTaskRequest.DueDate,
				CategoryID:          req.CreateTaskRequest.CategoryID,
				AuthenticatedUserID: 0,
			})

			if cErr != nil {
				_, wErr := connection.Write([]byte(cErr.Error()))
				if wErr != nil {
					log.Println("can't write data to connection", rErr)

					continue
				}
			}

			data, mErr := json.Marshal(&response)
			if mErr != nil {
				_, wErr := connection.Write([]byte(mErr.Error()))
				if wErr != nil {
					log.Println("can't marshal response", wErr)

					continue
				}

				continue
			}

			_, wErr := connection.Write(data)
			if wErr != nil {
				log.Println("can't write data to connection", rErr)

				continue
			}
		}

		connection.Close()
	}
}
