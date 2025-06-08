package main

import (
	"fmt"
)

func main() {
	fmt.Println("collection types")

	// arrays -> fixed length, contains items of the same type
	var breadCustomers [3]string
	fmt.Println(breadCustomers, len(breadCustomers))
	breadCustomers[0] = "mr lew"
	breadCustomers[1] = "anesu"
	breadCustomers[2] = "james"
	fmt.Println(breadCustomers)

	var luckyNumbers [2]int
	fmt.Println(luckyNumbers)
	luckyNumbers[0] = 324
	luckyNumbers[1] = 211

	// using array literal syntax
	grade1Students := [3]string{"alice", "luke", "jenifer"}
	fmt.Println(grade1Students)
	fmt.Println()

	// multidimensional array
	multi := [3][5]int{
		{1, 2, 4, 8, 8},
		{3, 44, 11, 785, 33},
		{7, 90, 98, 23, 55},
	}
	fmt.Println(multi[1][3]) //758

	// type of an array -> combination of its size and the underlying type
	arrayTyp := [2]string{"red", "blue"}
	fmt.Println(arrayTyp) //type of an array is [2]string

	// letting the compiler determine the array length
	animals := [...]string{"rabbit", "cow", "bird"}
	fmt.Println(len(animals))

	studentsA := [3]string{"james", "amos", "jack"}
	studentsB := studentsA
	studentsA[0] = "anna"
	fmt.Println(studentsA)
	fmt.Println(studentsB)

	// comparing arrays
	words1 := [2]string{"red", "food"}
	words2 := [2]string{"red", "food"}
	fmt.Println(words1 == words2)

	// enumerating an array
	studentsC := [3]string{"alice", "anne", "anna"}
	for _, v := range studentsC {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// slices
	// variable length array
	teachersNames := make([]string, 3)
	teachersNames[0] = "mrs smith"
	teachersNames[1] = "mr amos"
	teachersNames[2] = "mr more"
	fmt.Println(teachersNames)

	topStudents := []string{"james", "aaron"}
	fmt.Println(topStudents)
	/*
		slice -> contains 3 values [ pointer, length, capacity ]
		pointer -> a pointer to the array
		length -> the length of the slice [the number of elements a slice can store]
		capacity -> the capacity of the slice [the number of elements an array can store]
	*/

	phones := []string{"car"}
	phones = append(phones, "phone")
	phones = append(phones, "dog", "cat")
	fmt.Println(phones)

	goodNames := []string{"alan", "jenifer"}
	moreGoodNames := []string{"jessica", "paida"}
	allGoodNames := append(goodNames, moreGoodNames...)
	fmt.Println(goodNames)
	fmt.Println(allGoodNames)

	// creating slices from existing arrays
	numbersA := [5]int{77, 54, 34, 12, 78}
	numbersB := numbersA[:2]
	fmt.Println(numbersB)

	// working with maps
	// associates data values with keys!

	products := make(map[string]string, 3)
	products["best"] = "sadza"
	products["worst"] = "porridge"
	fmt.Println(len(products), products)
	fmt.Println(products["worst"])

	// using map literal syntax
	uncleSamProducts := map[int]string{
		1: "sandals",
		2: "hand bags",
	}
	uncleSamProducts[3] = "shoes"
	fmt.Println(uncleSamProducts)
	fmt.Println(uncleSamProducts[1])

	// checking for items in a map
	// using coma ok!
	bestRunners := map[int]string{
		1: "thomas",
	}
	value, ok := bestRunners[1]
	if ok {
		fmt.Println("runner is:", value)
	}

	if value, ok := bestRunners[1]; ok {
		fmt.Println("hello runner:", value)
	} else {
		fmt.Println("runner is not found!")
	}
	fmt.Println()
	// deleting items from a map
	cars := map[int]string{
		1: "toyota",
		2: "nissan",
	}
	fmt.Println(cars)

	if car, ok := cars[1]; ok {
		fmt.Println("deleting car:", car)
		delete(cars, 1)
	}
	fmt.Println(cars)

	// enumerating the contents of a map
	// using for and range

	singers := map[string]string{
		"best singer": "wink d",
		"best vocals": "marren morris",
		"best voice":  "passenger",
	}
	for key, singer := range singers {
		fmt.Println(singer, ":", key)
	}

	studentResults := map[int]string{
		43: "john",
		87: "lewis",
		24: "anna",
		98: "jack",
	}
	fmt.Println(studentResults)
}
