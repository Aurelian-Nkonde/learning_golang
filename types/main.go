package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		int(whole numbers), uint(positive whole numbers), byte(aliase for uint8),
		float32, float64(numbers with a fraction), bool(true or false), string(sequence of characters),
		rune(simply a representation of a single character)
	*/

	// using literal values
	fmt.Println(100 + 56)
	fmt.Println("this is a good message")

	// constants -> names for specific values
	// typed constants(type is defined) and untyped constants(no type defined)

	const teacherSalary float64 = 2000.56
	const emmaFriend string = "alice"
	fmt.Println(teacherSalary, emmaFriend)

	const jackSalary float32 = 3400.45
	const annaSalary float32 = 3500.44
	const months = 4
	fmt.Println((jackSalary + annaSalary) * months)

	// iota keyword -> used to create a series of successive untyped integer constants
	const (
		oslo = iota
		berlin
		brasilia
	)
	fmt.Println(oslo, berlin, brasilia)
	const boxAHeight, boxAWeight float32 = 55.90, 100.50
	fmt.Println(boxAHeight, boxAWeight)

	//variables
	var favCountry string = "brazil"
	var favNumber int = 456
	fmt.Println("Fav country:", favCountry, ", fav number:", favNumber)

	var johnFavFood = "rice and beans"
	fmt.Println(johnFavFood)

	// zero value
	var defaultInt int
	var defaultString string
	var defaultBool bool
	fmt.Println(defaultBool, defaultInt, defaultString == "")

	// short hand syntax, used only within functions
	goalKeeper := "thomas luke"
	canFly := false
	fmt.Println(goalKeeper, canFly)

	// blank identifier _
	firstGirl, _, thirdGirl := "alice", "jack", "pride"
	fmt.Println(firstGirl, thirdGirl)

	capitalCity, _ := "harare", "nothing"
	fmt.Println(capitalCity)

	// understanding pointers
	value1 := 100
	value2 := value1
	value1++
	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)

	// pointer -> a variable whose value is a memory address
	houseHeight := 4
	var houseHeightPointer *int //pointer type *type eg *int, *string, *float64
	houseHeightPointer = &houseHeight
	fmt.Println(houseHeightPointer)
	// using the address operator &

	// following a pointer -> reading the value at the memory address it refers to
	cowsAtTheFarm := 340
	cowsAtTheFarmP := &cowsAtTheFarm
	fmt.Println("cows at the farm", cowsAtTheFarm)
	*cowsAtTheFarmP++
	fmt.Println("added, cows at the farm:", cowsAtTheFarm)

	students := 790
	studentsP := &students
	*studentsP += 5
	fmt.Println("students:", students, studentsP)
	//dereferencing -> *pointerV

	goats := 800
	var goatsP *int
	fmt.Println("goats pointer:", goatsP)
	goatsP = &goats
	fmt.Println(goatsP, *goatsP)

	fmt.Println()

	colors := [3]string{"red", "blue", "green"}
	fmt.Println(colors)
	secondColor := &colors[1]
	fmt.Println("second color:", *secondColor)
	sort.Strings(colors[:])
	fmt.Println("second color:", *secondColor)
	fmt.Println(colors)

}
