package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello to TODO app")
	command := flag.String("command", "no command", "create a new task")
	flag.Parsed()
	scanner := bufio.NewScanner(os.Stdin)
	if *command == "create-task" {
		var name, duedate, category string
		fmt.Println("please enter the task title")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("please enter the task category")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("please enter the task du date")
		scanner.Scan()
		name = scanner.Text()
		fmt.Println("task:", name, category, duedate)
	}
}
