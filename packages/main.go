package main

import (
	"fmt"
	"strconv"

	"learning.packages/client"
	"learning.packages/customer"
	"learning.packages/products"
	proj "learning.packages/project"
)

// myModuleName/packageName

/*
	defined by creating code folders and files and using they keyword "package"

	define a package?
	-> create a folder and code files with package statements

	use a package?
	-> add import statement that specifies a path to the package

*/

/*
	Module file/? -> to make it easy to install packages that have been published
*/

/*
	using the custom package
	-> custom packages are declared using the import statement

	Uppercase type, function, method names -> exported
	Lowercase -> Unexported
*/

/*
	using a package alias -> to deal with package conflicts

	aliasName "moduleName/packageName"

	aliasName.ExportedMethod()
*/

/*
	Dot import!
	special alias that allows a package features to be used without using a prefix
*/

func main() {
	fmt.Println("packages!")

	terry := customer.Customer{Name: "terry afrika", Gender: "male", Age: 45}
	terry.PrintCustomerDetails()

	fmt.Println()

	myProducts := []*products.Product{
		{Name: "phone", Price: 350},
		{Name: "earphones", Price: 45},
	}
	for _, v := range myProducts {
		fmt.Println(v.Name + " is selling for:$ " + strconv.Itoa(v.Price))
	}

	fmt.Println()

	client1 := client.NewClient("thompson transport", "transportation", 12000)
	fmt.Println(client1)
	fmt.Println(client1.GetName())
	fmt.Println(client1.GetBusiness())
	fmt.Println(client1.GetPrice())
	fmt.Println(client1.GetClient())

	fmt.Println()

	projects := []*proj.Project{
		{Name: "mega city", Company: "2 rivers"},
		{Name: "online-store", Company: "brave-orbit"},
	}
	for _, v := range projects {
		fmt.Println(v.Name + " is done by a company called " + v.Company)
	}

}
