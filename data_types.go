package main

import "fmt"

func main() {

	//Bolean
	a := true
	b := false
	fmt.Println(" Boolean a:", a, "b:", b)

	//complex
	a1 := complex(4, 10)
	a2 := 2 + 5i
	sum1 := a1 + a2
	fmt.Println(" complex a1", a1, "a2", a2)
	fmt.Println(" complex sum", sum1)

	//constant
	C := 5
	println("constant c", C)
	//constant 2
	const n = "manali"
	var name = n
	fmt.Printf("type %T value %v", name, name)
	fmt.Println()

	//type casting
	d := 44
	e := 14.6
	sum := d + int(e)
	fmt.Println(sum)

	//idomatic
	if num := 9; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}
