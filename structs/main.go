package main

import (
	"fmt"
)

/*
	structs!
	data types comprised of fields
*/
// custom data types are created using structs!

type Customer struct { //struct type called Customer with 3 fields
	name string
	age  int
}

type Hacker struct {
	name, language string
	age            int
}

// defining embedded fields
type Person struct {
	name string
	age  int
}

type Nurse struct {
	Person //embedded field
	job    string
}

type Developer struct {
	name, language string
}

type Frontend struct {
	Developer
	frontDev Developer
	position string
}

func main() {
	fmt.Println("*")

	john := Customer{name: "john", age: 45}
	amos := Customer{name: "amos", age: 23}
	fmt.Println(john, amos)

	// creating a value
	jac43 := Hacker{"john lewis", "golang", 56}
	fmt.Println(jac43)
	jac43.language = "c++"
	fmt.Println(jac43)

	var thompson Customer
	fmt.Println(thompson.age == 0, thompson.name == "")
	thompson.age = 45
	thompson.name = "thompson"
	fmt.Println(thompson)

	// using the new function to create struct values!
	var aaron = new(Customer)
	fmt.Println(aaron)
	var aaron2 = &Customer{}
	fmt.Println(aaron2)

	anna := Nurse{
		Person: Person{name: "anna shaw", age: 23},
		job:    "nurse",
	}
	fmt.Println(anna)
	fmt.Println(anna.name, anna.Person.name)

	joyce := Frontend{
		Developer: Developer{name: "joyce", language: "c#"},
		frontDev:  Developer{name: "joyce", language: "typescript"},
		position:  "frontend developer",
	}
	fmt.Println(joyce.Developer.language, joyce.frontDev.language, joyce.position)

	// Anonymous struct types -> defined without using a name
	type Dog struct {
		name, color string
	}
	jackDog := Dog{"jack", "brown"}
	printDog(jackDog)
	printDog(Dog{"sheppy", "white"})

	type Person struct {
		name string
		age  int
	}

	type Farmer struct {
		Person
		products []string
	}

	farmersArray := [1]Farmer{
		{Person: Person{"andrew", 54}, products: []string{"pees", "maize"}},
	}
	fmt.Println(farmersArray)
	fmt.Println(farmersArray[0].Person.name, farmersArray[0].products[0])

	farmersSlice := []Farmer{
		{Person: Person{"jack", 23}, products: []string{"tomatoes"}},
	}
	fmt.Println(farmersSlice[0].Person.name, farmersSlice[0].products[0])
	fmt.Println(farmersSlice)

	farmersMap := map[string]Farmer{
		"alice": {
			Person:   Person{"alice", 45},
			products: []string{"beans", "corn"},
		},
	}
	if alice, ok := farmersMap["alice"]; ok != false {
		fmt.Println(alice.Person.name, alice.products[0])
	}

	type Car struct {
		name, color string
	}

	car1 := Car{"nissan", "red"}
	car2 := car1
	car2.color = "blue"
	fmt.Println(car1, car2)
	car3 := &car1
	car1.color = "black"
	fmt.Println(car1.color, (*car3).color, car3.color)

	teachers := [2]*Teacher{
		newTeacher("mr johns", 54),
		newTeacher("mrs lea", 52),
	}
	for _, t := range teachers {
		fmt.Println(t.name, t.age)
	}

}

func printDog(value struct{ name, color string }) {
	fmt.Println(value.name, value.color)
}

type Teacher struct {
	name string
	age  int
}

func newTeacher(name string, age int) *Teacher {
	return &Teacher{name, age}
}
