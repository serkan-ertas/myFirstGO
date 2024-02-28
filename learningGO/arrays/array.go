package arrays

import "fmt"

func ArrayFunc() {
	// create an array
	var productNames [4]string
	productNames = [4]string{"a", "b", "c", "d"}
	fmt.Println(productNames)

	// specific index
	fmt.Println(productNames[0])

	productNames[0] = "changed"
	fmt.Println(productNames[0])

	// another array
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	// slice	//first include //last exclude // slices are references
	featuredPrices := prices[1:]
	fmt.Println(featuredPrices)

	// slice features
	// for capacity, CANT count left, CAN count right
	fmt.Println("length of featuredPrices is:", len(featuredPrices))
	fmt.Println("capacity of featuredPrices is:", cap(featuredPrices))

	// can turn back to original arr  ***cant turn back to left!!***
	highlightedPrices := featuredPrices[:1]
	fmt.Println("highlightedPrices array is:", highlightedPrices)
	fmt.Println("length is:", len(highlightedPrices), "\ncapacity is:", cap(highlightedPrices))
	// turning back to original
	highlightedPrices = highlightedPrices[:cap(highlightedPrices)]
	fmt.Println("highlightedPrices array is:", highlightedPrices)
	fmt.Println("length is:", len(highlightedPrices), "\ncapacity is:", cap(highlightedPrices))
}
