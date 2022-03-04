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

func checkForInfinity(number float64) bool {
	return math.IsInf(number, 0)
}
func main() {

	greet := say_hello()
	fmt.Println(greet)

}
