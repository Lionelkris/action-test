package main

import (
	"fmt"
	"strconv"
	"strings"
)

func say_hello() string {

	return "Hello Hello"
}

func add(input string) (*float64, error) {
	var number, add float64
	var err error
	numbers := strings.Split(input, ",")
	fmt.Println(numbers)
	fmt.Printf("Length of string is %v", len(numbers))
	for _, numberString := range numbers {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 32); err != nil {
			fmt.Println("Cannot convert string to float", err)
			return nil, err
		}
		fmt.Printf("Converted Number is %v\n", number)
		//fmt.Printf("Add value is %v\n", add)
		add = add + number
		//fmt.Printf("Add value is %v\n", add)
	}
	fmt.Printf("Add value is %v\n", add)
	return &add, nil
}

func main() {

	greet := say_hello()
	fmt.Println(greet)

}
