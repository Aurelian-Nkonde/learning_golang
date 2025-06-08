package main

import (
	"fmt"
)

func main() {
	fmt.Println("*")
	// 1
	array1 := [5]int{}
	fmt.Println(array1)
	// array 2
	array2 := [3]string{"banana", "apple", "orange"}
	fmt.Println(array2)
	array3 := [4]bool{true, false, true, false}
	fmt.Println(array3[0], array3[len(array3)-1])
	// array 4
	array4 := [5]float64{23.65, 7.00, 55.67, 33.010, 12.234}
	fmt.Print(array4, " ")
	array4[2] = 55.002
	fmt.Print(array4, "\n")
	// array 5
	array5 := "hello!"
	for _, v := range array5 {
		fmt.Print(string(v), " ")
	}
	fmt.Print("\n")
	// array 6
	array6 := [5]int{10, 2, 3, 5, 1}
	array6Sum := 0
	for _, v := range array6 {
		array6Sum += v
	}
	fmt.Println(array6Sum)
	// array 7
	//array 8
	array8Old := [3]int{100, 400, 589}
	array8ONew := array8Old
	fmt.Println(array8ONew)
	// array 9
	array9 := [3]string{"ray", "lucy", "mark"}
	array9Result := nameExists("riay", array9)
	fmt.Println(array9Result)

	// slice 1
	slice1 := []int{}
	fmt.Println(slice1)

}

func nameExists(name string, names [3]string) bool {
	var numb int = 0
	for i := 0; i < len(names); i++ {
		if names[i] == name {
			numb += 1
		}
	}
	return numb > 0
}
