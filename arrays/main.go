package main

import "fmt"

type products struct {
	title string
	price float64
	description string
}

func main() {
	// var testingArr = [][4]string{{"Seif", "Hesham", "Salem", "Sayed"}, {"Mohamed", "Ebrahim", "Hussein", "Salama"}}

	// for i := 0; i < len(testingArr[0]); i += 1 {
	// 	name := testingArr[0][i]
	// 	fmt.Println(name)
	// }

	// fmt.Print(testingArr[0][0], ", ",testingArr[0][3])
	// fmt.Println()
	// fmt.Print(testingArr[1][0], ", ",testingArr[1][3])

	var hobbies = [3]string{"Programming", "Gaming", "Reading"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1],hobbies[2])
	firstFirstSlice := hobbies[0:2]
	firstSecondSlice := hobbies[0:1]
	secondSecondSlice := hobbies[1:2]

	fmt.Println(firstFirstSlice)
	fmt.Println(firstSecondSlice, secondSecondSlice)

	firstFirstSlice = firstFirstSlice[1:3]
	fmt.Println(firstFirstSlice)
	

	myProducts := []products{{title: "Cakes", price: 22.22, description: "This is a great cake"},
	{title: "Grokking Algorithms", price: 21.22, description: "This is a great book"}}
	
	fmt.Print(myProducts)
}