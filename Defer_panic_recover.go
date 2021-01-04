package main

import (
	"fmt"
	"os"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func sub(n1 int, n2 int) int {

	add := n1 + n2
	fmt.Println("Add", add)
	return add
}

func recoverfrompanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic ", r)
	}

}

func main() {

	//Error handling
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")

	//defer
	defer sub(1, 2)
	fmt.Println("defer")

	//panic & recover
	defer recoverfrompanic()
	firstName := "manali"
	//lastname := "datar"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	//channel
	var a chan int
	a = make(chan int)
	fmt.Printf("Channle is %T", a)

}
