package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Add(input string) (*float64, error) {
	var number, add float64
	var err error
	numbers := strings.Split(input, ",")
	for _, numberString := range numbers {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 64); err != nil {
			return nil, fmt.Errorf("Invalid input. Try adding numbers next time!")
		}
		add = add + number
		if add < math.MaxFloat64/10000 {
			add = math.Round(add*10000) / 10000
		}
	}
	if checkForInfinity(add) {
		return nil, fmt.Errorf("Better get a super computer! Out of bound exception")
	}
	return &add, nil
}

func Subtract(input string) (*float64, error) {
	var sub float64
	var err error
	numbers := strings.SplitN(input, ",", 2)
	sub, err = strconv.ParseFloat(numbers[0], 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid input. Try subtracting numbers next time!")
	}
	addResult, err := Add(numbers[1])
	if err != nil {
		return nil, err
	}
	sub = sub - *addResult

	if checkForInfinity(sub) {
		return nil, fmt.Errorf("Better get a super computer! Out of bound exception")
	}
	return &sub, nil
}

func Multiply(input string) (*float64, error) {
	var number, product float64
	var err error
	product = 1
	numbers := strings.Split(input, ",")
	for _, numberString := range numbers {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 64); err != nil {
			return nil, fmt.Errorf("Invalid input. Try multiplying numbers next time!")
		}
		product = product * number
		if product < math.MaxFloat64/10000 {
			product = math.Round(product*10000) / 10000
		}
	}
	if checkForInfinity(product) {
		return nil, fmt.Errorf("Better get a super computer! Out of bound exception")
	}
	return &product, nil
}

func Divide(input string) (*float64, error) {
	var err error
	var number, result float64
	numbers := strings.SplitN(input, ",", 2)
	if result, err = strconv.ParseFloat(numbers[0], 64); err != nil {
		return nil, fmt.Errorf("Invalid input. Try dividing numbers next time!")
	}
	fmt.Println(numbers)
	if len(numbers) == 1 {
		return &result, nil
	}
	for _, numberString := range strings.Split(numbers[1], ",") {
		if numberString == "" {
			continue
		}
		if number, err = strconv.ParseFloat(numberString, 64); err != nil {
			return nil, fmt.Errorf("Invalid input. Try dividing numbers next time!")
		}
		if number == 0 {
			return nil, fmt.Errorf("don't divide by zero, it won't work")
		}
		result = result / number
	}
	if checkForInfinity(result) {
		return nil, fmt.Errorf("Better get a super computer! Out of bound exception")
	}
	return &result, nil
}

func checkForInfinity(number float64) bool {
	return math.IsInf(number, 0)
}
