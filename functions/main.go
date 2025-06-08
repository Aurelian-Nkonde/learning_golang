package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("*")
	greet()
	hello("sam")
	hello("jane")
	sum(100, 34)
	printSuppliers("meat", "james$sons", "the watsons")
	printProducts("james")
	printProducts("the slayers", "beef", "pork")
	students := []string{"jelly", "alice"}
	printStudents("grade 4", students...)

	first := 35
	second := 103
	fmt.Println(first, second)
	swap(first, second)
	fmt.Println(first, second)
	fmt.Println()

	firstP, secondP := 230, 400
	fmt.Println("before pointer swap", firstP, secondP)
	swapWithPointers(&firstP, &secondP)
	fmt.Println("after pointer swap", firstP, secondP)
	hacker := helloHacker("james34")
	fmt.Println(hacker)

	doubled, uppercased := doubledOrUppercased("thousand the creator", 453)
	fmt.Println(doubled, uppercased)
	fmt.Println()

	namedSum, namedWord := namedReturnValues(10, 44, "ola", "mundo")
	fmt.Println(namedSum, namedWord)
	fmt.Println()

	usingDeferKeyword()
}

func greet() {
	fmt.Println("hello user!")
}

func hello(name string) {
	fmt.Println("hello,", name)
}

func sum(a, b int) {
	fmt.Println(a + b)
}

// defining variadic parameters
// variadic parameter accepts a variable number of values

func printSuppliers(product string, suppliers ...string) {
	for _, supp := range suppliers {
		fmt.Println(product, "supplier", supp)
	}
}

// value... -> spread operator
// value ...type -> rest parameter

func printProducts(supplier string, products ...string) {
	if len(products) == 0 {
		fmt.Println("Supplier", supplier, "has got 0 products")
		return
	}
	for _, v := range products {
		fmt.Println(supplier, " supplied", v)
	}
}

func printStudents(class string, students ...string) {
	if students == nil {
		fmt.Println(class, "Class has got zero students")
		return
	}
	for _, v := range students {
		fmt.Println(class, "student name is", v)
	}
}

func swap(first, second int) {
	fmt.Println("before swap func runs!", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("after the swap func runs!", first, second)
}

func swapWithPointers(first, second *int) {
	temp := *first
	*first = *second
	*second = temp
}

// functions define results
func helloHacker(hacker string) string {
	return "hello" + hacker
}

// returning multiple function results

func doubledOrUppercased(word string, numb int) (string, int) {
	return strings.ToUpper(word), numb * 2
}

// using named results
func namedReturnValues(a, b int, c, d string) (total int, joined string) {
	total = a + b
	joined = c + " " + d
	return
}

// using the defer keyword
func usingDeferKeyword() {
	fmt.Println("function started!")
	defer fmt.Println("defer function called()")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
	fmt.Println("function about to return")
	defer fmt.Println("second function is called()")
}
