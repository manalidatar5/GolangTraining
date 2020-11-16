package main

import "fmt"

//fucntion

func sumOfnumbers(num1 int, num2 int) int {
	//function body
	var sum = num1 + num2
	return sum
}

//Recursion function

func facto(n int) int {
	if n == 0 {

		return 1
	}
	return n * facto(n-1)
}

func main() {

	//normal loop
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)

		} else {
			fmt.Println("Odd", i)
		}
		fmt.Println("i =", i)
	}
	//Infinite loop
	// for {
	// 	fmt.Printf("Hello")
	// }
	// simple range

	rvariable := []string{"ABC", "PQR", "GEF"}

	for i, j := range rvariable {
		fmt.Println(i, j)
	}
	//Only without index
	for _, j := range rvariable {
		fmt.Println(j)
	}
	var totalsum = sumOfnumbers(5, 7)
	fmt.Println("Sum is", totalsum)

	var fact = facto(7)
	fmt.Println("Factorial is", fact)

}
