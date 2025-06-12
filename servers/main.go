package main

import (
	"fmt"
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, sh.message)
}

/*
	functions for setting up a server
	ListenAndServe --> listening for HTTP request
	ListenAndServeTLS --> listening for HTTPS requests

	when request arrives it is passed on to the handler
	request -> handler

	handler is an interface

	type handler interface {
		ServeHTTP(writer, request)
	}
*/

func main() {
	err := http.ListenAndServe(":5000", StringHandler{message: "ola, mundo!"})
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
}

/*
	Large-scale web application -> application should be scalable
	Scalable -> Able to easily and quickly increase the capacity of the app to take on a bigger volume of request

	Scaling in 2 ways!
	Vertical scaling -> increasing the amount of CPU's or capacity in a single machine
	Horizontal scaling -> increasing the number of machines to expand capacity

	Go scales well vertically

*/
