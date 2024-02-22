package pointers

import "fmt"

func PointerFunc() {
	age := 32
	fmt.Println("Age: ", age)

	agePtr := &age
	fmt.Println(agePtr)
}
