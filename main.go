package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func main() {
	name := "Hassn"
	stringRider := strings.NewReader(name)
	scanner := bufio.NewScanner(stringRider)
	scanner.Scan()
	fmt.Println("output")
	fmt.Println(scanner.Text())
	var scores = Int{6, 78, 25, 74, 36, 4, 1, 5, 7}

	fmt.Println("before", scores)
	sort.Sort(scores)
	fmt.Println("after", scores)

}

type Int []int

func (in Int) Len() int {
	return len(in)
}
func (in Int) Less(i, j int) bool {
	return in[i] < in[j]
}
func (in Int) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

type User struct {
	ID   uint
	Name string
}
type userStore map[uint]User
