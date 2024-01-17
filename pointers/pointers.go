package pointers

import "fmt"

func main() {

	age := 32

	agePointer := &age

	fmt.Print(calcIsAdult(*agePointer));

}

func calcIsAdult(age int) int {
	return age - 18
}