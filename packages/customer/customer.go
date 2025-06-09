package customer //the package name

import (
	"fmt" // the name specified by the package should match the name of the containing folder
	"strconv"
)

type Customer struct {
	Name, Gender string
	Age          int
}

func (customer Customer) PrintCustomerDetails() {
	fmt.Println(customer.Name + " is a " + customer.Gender + ", is also " + strconv.Itoa(customer.Age) + " years old!")
}
