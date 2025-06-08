package main

import (
	"fmt"
	"strconv"
)

/*
	methods -> functions that are involved on struct and have access to all fields by struct value
	interfaces -> defines set of methods that can implemented by a struct value
*/

type Farmer struct {
	name string
	age  int
}

func (farmer *Farmer) printDetails() {
	fmt.Println(farmer.name, "is", farmer.age, "years old!!")
}

func printFarmerDetails(farmer *Farmer) {
	fmt.Println(farmer.name, "is", farmer.age, "years old!")
}

type Dealer struct {
	deals []*Product
	name  string
	age   int
}

type Product struct {
	name  string
	price int
}

func (dealer *Dealer) dealerTotalDealsAmount() int {
	total := 0
	for _, v := range dealer.deals {
		total += v.price
	}
	return total
}

func (dealer *Dealer) dealerProfit(borrowed int) int {
	total := 0
	for _, v := range dealer.deals {
		total += v.price
	}
	return total - borrowed
}

func (dealer Dealer) dealerDetails() {
	fmt.Println(dealer.name, "is", dealer.age, "years old.")
}

type Student struct {
	name string
	mark int
}

func (student *Student) doubleMark() {
	fmt.Println(student.mark * 2)
}

func (student *Student) multiMark(by int) {
	fmt.Println(student.mark * by)
}

func main() {
	fmt.Println("**")
	farmers := []*Farmer{
		{"james", 34},
		{"alice", 74},
	}
	fmt.Println("Printing out farmers!")
	for _, v := range farmers {
		// printFarmerDetails(v)
		v.printDetails()
	}

	fmt.Println()

	dealers := []*Dealer{
		{name: "john", age: 34, deals: []*Product{
			{name: "phone", price: 340},
			{name: "car", price: 5400},
		}},
		{name: "joseph", age: 23, deals: []*Product{
			{name: "shoes", price: 23},
		}},
	}
	for _, v := range dealers {
		v.dealerDetails()
	}
	for _, v := range dealers {
		fmt.Println(v.name, "got this amount:$", v.dealerTotalDealsAmount(), "from all deals!")
	}
	for _, v := range dealers {
		if v.dealerProfit(300) > 0 {
			fmt.Println(v.name, "got profit:", v.dealerProfit(300), "yeah!")
		} else {
			fmt.Println(v.name, "did not get any profit:", v.dealerProfit(300))
		}
	}

	fmt.Println()

	sandra := Student{"sandra", 45}
	sandra.doubleMark()
	sandra.multiMark(3)

	fmt.Println()

	sean := Developer{name: "sean", age: 34, language: "c++"}
	printITWorkerDetails(sean)

	fmt.Println()

	jackDonkey := donkey{"jack"}
	fmt.Println(jackDonkey.sound())
	animalSounds(jackDonkey)

	fmt.Println()

	transports := []Transport{
		Cart{"ngoro"},
		Train{"gauten train"},
	}
	for _, v := range transports {
		fmt.Println(v.getName() + " is a good transport means!")
	}

	fmt.Println()

	pets := []Pet{
		Cat{"alana"},
		Dog{"halo", "black"},
	}
	for _, v := range pets {
		if cat, isCat := v.(Cat); isCat == true {
			fmt.Println(cat.name + " is amazing!")
		}
		if dog, isDog := v.(Dog); isDog == true {
			fmt.Println(dog.name + " is " + dog.color + " in color, woow!")
		}
	}
	fmt.Println()

	var markEngineer *Engineer = &Engineer{"mark", "civil engineer"}
	everything := []interface{}{
		markEngineer,
		"hello world",
		[]int{1, 2, 3, 9000},
		false,
	}
	for _, v := range everything {
		switch value := v.(type) {
		case Engineer:
			fmt.Println(value.name + "is a " + value.proffession)
		case string:
			fmt.Println(value + "is a string!")
		case []int:
			fmt.Println("the length of the slice is:", len(value))
		case bool:
			fmt.Println("the value is a boolean:", value)
		}
	}

}

// the way the methods are invoked is what makes functions and methods different

// go does not support method overloading
// -> multiple methods defined with the same name but different parameters

// interfaces -> describes a set of methods without specifying the implementation of those methods
/*
	if a types implements all the methods defined by the interface, then a value of that type
	can be used wherever the interface is permitted
*/

type ITWorker interface {
	getName() string
	getDetails() string
}

type Developer struct {
	name, language string
	age            int
}

type Hacker struct {
	name string
	age  int
}

func (dev Developer) getName() string {
	return dev.name
}

func (dev Developer) getDetails() string {
	return dev.name + " is " + strconv.Itoa(dev.age) + " years, and code in " + dev.language
}

func printITWorkerDetails(worker ITWorker) {
	name := worker.getName()
	details := worker.getDetails()
	fmt.Println(name)
	fmt.Println(details)
}

/*
	to implement an interface, all the methods defined by the interface must be defined for a struct type!
*/

type Animal interface {
	sound() string
}

type donkey struct {
	name string
}

func (donk donkey) sound() string {
	return donk.name + " makes the sound: Oio ooi ooi!"
}

func animalSounds(animal Animal) {
	fmt.Println(animal.sound())
}

type Transport interface {
	getName() string
}

type Cart struct {
	name string
}

func (car Cart) getName() string {
	return car.name
}

type Train struct {
	name string
}

func (train Train) getName() string {
	return train.name
}

/*
	performing type assetations
	accessing the dynamic type directly[type narrowing]

	type assertion -> is used to access the dynamic value of an interface value
*/

type Pet interface {
	getName() string
}

type Cat struct {
	name string
}

type Dog struct {
	name, color string
}

func (cat Cat) getName() string {
	return cat.name
}

func (dog Dog) getName() string {
	return dog.name
}

// empty interface
/*
	an interface that defines no methods, represent an type
*/

type Engineer struct {
	name, proffession string
}
