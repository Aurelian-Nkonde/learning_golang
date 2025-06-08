package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("control flow!")
	busFare := 20
	if busFare >= 15 {
		fmt.Println("Prices are really up!")
	}

	childName := "lorrain"
	if len(childName) > 5 {
		message := "this is a long name!"
		fmt.Println(message)
	} else {
		message := "a good short name!"
		fmt.Println(message)
	}
	fmt.Println()

	// Using an Initialization Statement with an if Statement
	schoolFees := "2500o"
	if fees, err := strconv.Atoi(schoolFees); err == nil {
		fmt.Println("student fees:", fees)
	} else {
		fmt.Println("fees are in wrong format")
	}
	fmt.Println()

	counter := 0
	for {
		fmt.Println("Magic count is:", counter)
		counter++
		if counter == 4 {
			fmt.Println("counter hits the target number!")
			break
		}
	}

	eggs := 0
	for eggs < 4 {
		fmt.Println("counting eggs:", eggs)
		eggs++
	}

	fmt.Println()

	for chickens := 0; chickens < 3; chickens++ {
		fmt.Println("Counting chickens:", chickens)
	}

	for counter := 0; counter < 5; counter++ {
		if counter == 3 {
			fmt.Println("3 is hit!")
			fmt.Println("bad number hit, exiting now..!")
			break
		}
		fmt.Println("number hit:", counter)
	}

	fmt.Println()

	for counter := 0; counter < 5; counter++ {
		if counter == 3 {
			fmt.Println("3 is hit!")
			continue
		}
		fmt.Println("printing number:", counter)
	}

	// Enumerating Sequences
	// using a range keyword
	word := "hello"
	for i, v := range word {
		fmt.Println(i, ":", string(v))
	}

	fmt.Println()

	for _, v := range "oops" {
		fmt.Println("character value:", string(v))
	}

	//the blank identifier _ is used for ommiting either the index or value

	// using a switch statement
	day := time.Monday
	switch day {
	case time.Sunday:
		fmt.Println("a day for church")
	case time.Wednesday:
		fmt.Println("a day for movies")
	case time.Monday:
		fmt.Println("a day for school")
	default:
		fmt.Println("day is not found")
	}
}
