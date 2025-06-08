package main

import (
	"fmt"
)

func main() {
	fmt.Println("*")
	products := map[string]int{
		"phone": 2000,
		"car":   8000,
	}
	for pro, price := range products {
		var printProduct func(string)
		if price > 5000 {
			printProduct = printExpensiveProduct
		} else {
			printProduct = printCheapProduct
		}
		printProduct(pro)
	}
}

// understanding function types
// functions have a data type in go!, meaning they can be assigned to variables, used as parameters etc

func printExpensiveProduct(product string) {
	fmt.Println(product, "is expensive, yeah!")
}

func printCheapProduct(product string) {
	fmt.Println(product, "is cheapas, Oops!")
}
