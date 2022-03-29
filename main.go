package main

import (
	"fmt"
	"math"
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
	for _, numberString := range numbers {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 64); err != nil {
			return nil, err
		}
		add = add + number
		if add < math.MaxFloat64/10000 {
			add = math.Round(add*10000) / 10000
		}
	}
	if checkForInfinity(add) {
		fmt.Println("error from here")
		return nil, fmt.Errorf("out of bound")
	}
	return &add, nil
}

func subtract(input string) (*float64, error) {
	var sub float64
	var err error
	numbers := strings.SplitN(input, ",", 2)
	sub, err = strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		return nil, err
	}
	addResult, err := add(numbers[1])
	if err != nil {
		return nil, err
	}
	sub = sub - *addResult

	if checkForInfinity(sub) {
		return nil, fmt.Errorf("out of bound")
	}
	return &sub, nil
}

func multiply(input string) (*float64, error) {
	var number, product float64
	var err error
	product = 1
	numbers := strings.Split(input, ",")
	for _, numberString := range numbers {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 64); err != nil {
			return nil, err
		}
		product = product * number
		if product < math.MaxFloat64/10000 {
			product = math.Round(product*10000) / 10000
		}
	}
	if checkForInfinity(product) {
		fmt.Println("error from here")
		return nil, fmt.Errorf("out of bound")
	}
	return &product, nil
}

func checkForInfinity(number float64) bool {
	return math.IsInf(number, 0)
}
func main() {

	greet := say_hello()
	fmt.Println(greet)

}
