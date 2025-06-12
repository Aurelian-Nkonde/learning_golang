package main

type Product struct {
	name  string
	price float64
}

var Products = []*Product{
	{name: "phone", price: 1200.00},
	{name: "toyota", price: 12000.40},
	{name: "laptop", price: 400.10},
}
