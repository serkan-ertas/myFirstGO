package main

import "fmt"

//lint:file-ignore U1000 dont wanna see unused val/func

func pointerFunc() {
	age := 32
	fmt.Println("Age: ", age)
	agePtr := &age
}
