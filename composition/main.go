package main

import (
	"fmt"
	"strconv"
)

/*
	Go does not use inheritance, instead it relies on an approach called composition

	what is composition?
	-> a process by which new types are created by compining structs and interfaces
*/

/*
	Understanding type composition!
	composition can be used to create hierarchies
*/

type Product struct {
	name, category string
	price          int
}

func NewProduct(name, category string, price int) *Product { //constructor function
	return &Product{name, category, price}
}

type Car struct {
	*Product
	color string
	year  int
}

func newCar(name, category string, price int, color string, year int) *Car {
	return &Car{
		NewProduct(name, category, price),
		color,
		year,
	}
}

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

type Doctor struct {
	*Person
	hospital string
}

func NewDoctor(name, hospital string, age int) *Doctor {
	return &Doctor{
		NewPerson(name, age),
		hospital,
	}
}

type Teacher struct {
	*Person
	school string
	class  int
}

func NewTeacher(name, school string, class, age int) *Teacher {
	return &Teacher{
		NewPerson(name, age),
		school,
		class,
	}
}

func (teacher Teacher) TeacherDetails() {
	fmt.Println(teacher.Person.name + " is " + strconv.Itoa(teacher.age) + " years old. teaches at " + teacher.school + " a class " + strconv.Itoa(teacher.class))
}

func main() {
	fmt.Println("composition!")

	iphone := NewProduct("iphone 20", "phone", 1200)
	msi := NewProduct("msi laptop", "laptop", 3000)
	for _, v := range []*Product{iphone, msi} {
		fmt.Println(v.name + " of cateory " + v.category + " costs: $ " + strconv.Itoa(v.price))
	}

	toyota := newCar("toyota wish", "car", 12000, "red", 2004)
	fmt.Println(toyota.Product.name)

	fmt.Println()

	mrJones := NewTeacher("mr aaron johns", "st anna", 4, 65)
	mrJones.TeacherDetails()
}
