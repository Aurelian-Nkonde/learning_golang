package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("*")
	word1, word2 := "hello", "world"
	fmt.Println(word1, ",", word2)
	fmt.Println(100 * 2)

	// values must be of the same type

	// allows integer values to overflow by wrapping around
	intValue := math.MaxInt64
	floatValue := math.MaxFloat64
	fmt.Println(intValue+1, floatValue*2)
	fmt.Println()

	// go remainder operator %
	fmt.Println(3 % 2)
	fmt.Println(-3 % 2)
	fmt.Println()

	// explicity conversion
	brownGoats := 60
	idealChickens := 40.50
	totalAnimals := brownGoats + int(idealChickens)
	fmt.Println("total animals:", totalAnimals)
	// syntax is: T(value) eg float64(int), int(float32)

	height := 10.70
	fmt.Println(int(height))
	fmt.Println(math.Ceil(height))
	fmt.Println(math.Floor(height))
	fmt.Println(math.Floor(height))

	fmt.Println()

	// parsing from strings
	// strconv -> provides methods for converting string values to oother simple data types
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseFloat("50.54", 64))
	fmt.Println(strconv.ParseInt("105", 10, 64))
	fmt.Println(strconv.Atoi("105"))

	eggs := "346"
	eggCount, err := strconv.Atoi(eggs)
	if err == nil {
		fmt.Println("total eggs:", eggCount)
	} else {
		fmt.Println("error getting eggs count")
	}

	studentsAtSchool := "4500"
	if studentsCount, err := strconv.Atoi(studentsAtSchool); err == nil {
		fmt.Println("number of students:", studentsCount)
	} else {
		fmt.Println("error getting students total number")
	}

	// formatting values as strings
	/*
		format -> other data types to strings
		parsing -> string to other data types
	*/
	fmt.Println(strconv.FormatInt(102, 10))
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.Itoa(653))
}
