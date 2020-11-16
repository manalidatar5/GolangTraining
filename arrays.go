package main

import "fmt"

func main() {

	//array

	a := [3]int{3, 4, 5}
	fmt.Println(a)

	//slice
	a1 := [5]int{76, 77, 78, 79, 80}
	var b []int = a1[1:3]
	fmt.Println(b)

	//append
	fruits := []string{"Mango", "Banana", "Apple"}
	fmt.Println("Fruits old array:", fruits, len(fruits), cap(fruits))
	fruits = append(fruits, "Grapes")
	fmt.Println("Fruits new array:", fruits, len(fruits), cap(fruits))

	//map
	employeeID := make(map[int]string)
	employeeID[11] = "Manali"
	employeeID[12] = "Sayali"
	employeeID[13] = "Tejal"
	fmt.Println("employeeID map:", employeeID)

}
