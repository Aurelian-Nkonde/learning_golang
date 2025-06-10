package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	Error is an interface!

	type error interface {
		Error() string
	}
*/

/*
	Generating an error!
	functions, methods can express exceptional or unexpected outcomes by producing error responses.
*/

type CategoryError struct {
	requestedCategory string
}

func (e *CategoryError) Error() string {
	return "Category " + e.requestedCategory + " does not exist!"
}

type YearSpendings struct {
	project, catgory string
	amount           float64
}

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

var spendings = []*YearSpendings{
	{project: "ground improvement", catgory: "infrastructure", amount: 3000.10},
	{project: "office toilets", catgory: "infrastructure", amount: 3500.200},
	{project: "new cars", catgory: "transport", amount: 12000.00},
	{project: "new bus", catgory: "transport", amount: 34000.00},
}

func totalPrice(catergory string) (total float64, err *CategoryError) {
	productionCount := 0
	for _, s := range spendings {
		if s.catgory == catergory {
			total += s.amount
			productionCount++
		}
	}
	if productionCount == 0 {
		err = &CategoryError{requestedCategory: catergory}
	}
	return
}

var categories = []string{"transport", "infrastructure", "bills"}

/*
	Using the error convinience functions!

	errors pkg, New function (it creates simple errors)
*/

type ShopProduct struct {
	name  string
	price float64
}

var shopProducts = []*ShopProduct{
	{name: "sugar", price: 4.00},
	{name: "salt", price: 1.50},
	{name: "rice", price: 4.50},
}

func doubleShopPrice(product string) (value float64, err error) {
	found := 0
	for _, sp := range shopProducts {
		if sp.name == product {
			value += sp.price * 2
			found++
		}
	}
	if found == 0 {
		err = errors.New(product + " does not exist!")
	}
	return
}

/*
	Errors with more complex string contents

	fmt.Errorf
*/

type Student struct {
	name, subject string
	marks         int
}

var studentsAndMarks = []*Student{
	{name: "jack", subject: "maths", marks: 78},
	{name: "jack", subject: "science", marks: 56},
	{name: "jack", subject: "shona", marks: 43},
	{name: "anna", subject: "shona", marks: 65},
	{name: "anna", subject: "science", marks: 76},
	{name: "anna", subject: "geography", marks: 34},
	{name: "tom", subject: "english", marks: 80},
}

func studentFinalMarks(student string) (total int, err error) {
	found := 0
	for _, s := range studentsAndMarks {
		if s.name == student {
			total += s.marks
			found++
		}
	}
	if found == 0 {
		err = fmt.Errorf("%s is not found", student)
	}
	return
}

/*
	Dealing with unrecoverable errors!
	panicking -> a serious error that should lead to the immediate termination of the application
*/

type Framework struct {
	name  string
	stars int
}

var frameworks = []*Framework{
	{name: "react", stars: 24000},
	{name: "nodejs", stars: 19000},
	{name: "laravel", stars: 16000},
}

func PrintFrameworkDetails(framework string) {
	for _, f := range frameworks {
		if f.name == framework {
			fmt.Println(f.name + " has " + strconv.Itoa(f.stars) + " on github!")
			return
		} else {
			panic(fmt.Errorf("%s does not exist", framework)) //panic with an error
		}
	}
}

func main() {
	for _, c := range categories {
		total, err := totalPrice(c)
		if err == nil {
			fmt.Println(c, "total:", ToCurrency(total))
		} else {
			fmt.Println(c, "(No such category!)")
		}
	}

	fmt.Println()

	doubledPriceRice, err := doubleShopPrice("rice")
	if err == nil {
		fmt.Println("rice doubled price:", doubledPriceRice)
	} else {
		fmt.Println(err)
	}

	doubledPriceOil, err := doubleShopPrice("oil")
	if err == nil {
		fmt.Println("oil doubled price:", doubledPriceOil)
	} else {
		fmt.Println(err)
	}

	fmt.Println()
	tomMarks, err := studentFinalMarks("jack")
	if err == nil {
		fmt.Println("tom final marks:", tomMarks)
	} else {
		fmt.Println(err)
	}
	aliceMarks, err := studentFinalMarks("alice")
	if err == nil {
		fmt.Println("alice final marks:", aliceMarks)
	} else {
		fmt.Println(err)
	}

	fmt.Println()
	PrintFrameworkDetails("django")
	PrintFrameworkDetails("react")

}
