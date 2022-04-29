package main

import (
	"action-test/api"
	"fmt"
)

func say_hello() string {

	return "Hello Hello"
}

func main() {

	greet := say_hello()
	fmt.Println(greet)
	api.Server()

}
