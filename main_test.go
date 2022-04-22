package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_say_hello(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "simple test",
			want: "Hello Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := say_hello(); got != tt.want {
				t.Errorf("say_hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:  "good path with valid numbers",
			input: "13,41",
			want:  54.0,
		},
		{
			name:    "invalid numbers as input with string in it",
			input:   "13a,41",
			wantErr: true,
		},
		{
			name:    "only one number as input and nothing after comma",
			input:   "13,",
			want:    13.0,
			wantErr: false,
		},
		{
			name:    "invalid second number as input with string in it",
			input:   "13,41b",
			wantErr: true,
		},
		{
			name:    "invalid second number as input with string in it",
			input:   "13 ,41",
			wantErr: true,
		},
		{
			name:    "passing float values as input",
			input:   "13.9,41.2",
			want:    55.1,
			wantErr: false,
		},
		{
			name:    "passing float values as input",
			input:   "1.9,4.2,5",
			want:    11.1,
			wantErr: false,
		},
		{
			name:    "passing float values as input",
			input:   "1.9,4.2,-5",
			want:    1.1,
			wantErr: false,
		},
		{
			name:    "passing float values as input with negative",
			input:   "1.9,4.25467,-5",
			want:    1.1547,
			wantErr: false,
		},
		{
			name:    "passing float values as input",
			input:   "1.9,356770,-5",
			want:    356766.9,
			wantErr: false,
		},
		{
			name:    "passing float values as input plus one",
			input:   fmt.Sprintf("%v,1.0e+300", math.MaxFloat64),
			wantErr: true,
		},
		{
			name:    "passing float values as input plus one 565",
			input:   "565,1.8E+308",
			wantErr: true,
		},
		{
			name:    "passing exact max float values as input plus one 565",
			input:   fmt.Sprintf("565,%v", math.MaxFloat64),
			want:    math.MaxFloat64,
			wantErr: false,
		},
		{
			name:    "playing with boundaries",
			input:   fmt.Sprintf("1,%v", math.MaxFloat64/100),
			want:    math.MaxFloat64/100 + 1,
			wantErr: false,
		},
		{
			name:    "passing max float and 33 but wont do add operation",
			input:   fmt.Sprintf("%v,33", math.MaxFloat64),
			want:    37 + math.MaxFloat64,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add(tt.input)
			if err == nil {
				fmt.Printf("Value of *got is %v\n", *got)
				assert.Equal(t, tt.want, *got, "these should be equal")
			}
			assert.Equal(t, err != nil, tt.wantErr)

		})
	}
}

func Test_subtract(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:  "2 valid numbers",
			input: "40,20",
			want:  20.0,
		},
		{
			name:  "valid numbers with negative int",
			input: "-40,20",
			want:  -60.0,
		},
		{
			name:  "3 valid numbers with negative ints",
			input: "-40,20,-30",
			want:  -30.0,
		},
		{
			name:    "going below minimum Float64",
			input:   fmt.Sprintf("-%v,%v", math.MaxFloat64, math.MaxFloat64),
			wantErr: true,
		},
		{
			name:    "Insert non numeric value",
			input:   "!,7",
			wantErr: true,
		},
		{
			name:    "Insert non numeric value to subtract",
			input:   "7,a",
			wantErr: true,
		},
		{
			name:    "passing float values as input",
			input:   "1.9,4.2,-5",
			want:    2.7,
			wantErr: false,
		},
		{
			name:  "should get zero subtracting a big number from itself",
			input: fmt.Sprintf("%v,%v", math.MaxFloat64, math.MaxFloat64),
			want:  0,
		},
		{
			name:  "should get less than zero subtracting a big number from itself and then 33",
			input: fmt.Sprintf("%v,%v,33", math.MaxFloat64, math.MaxFloat64),
			want:  0,
		},
		{
			name:    "playing with boundaries",
			input:   fmt.Sprintf("1,%v", math.MaxFloat64/100),
			want:    1 - math.MaxFloat64/100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subtract(tt.input)
			assert.Equal(t, err != nil, tt.wantErr)
			if err != nil {
				fmt.Printf("Error: %s)", err.Error())
			}
			if err == nil {
				fmt.Printf("Value of *got is %v\n", *got)
				assert.Equal(t, tt.want, *got, "these should be equal")
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:    "simple good path with decimal fraction",
			input:   "1,0.5",
			want:    0.5,
			wantErr: false,
		},
		{
			name:    "simple good path",
			input:   "1,2",
			want:    2,
			wantErr: false,
		},
		{
			name:    "multiply by zero case",
			input:   "0,3.5",
			want:    0,
			wantErr: false,
		},
		{
			name:    "multiply by non digit case",
			input:   "4,d",
			wantErr: true,
		},
		{
			name:    "i know we are still concerned about boundaries breakpoint",
			input:   fmt.Sprintf("%v,1.0000000000000001", math.MaxFloat64),
			want:    math.MaxFloat64,
			wantErr: false,
		},
		{
			name:    "i know we are still concerned about boundaries, finally!!!",
			input:   fmt.Sprintf("%v,1.000000000000001", math.MaxFloat64),
			wantErr: true,
		},
		{
			name:    "i know we are still concerned about boundaries but not with 1",
			input:   fmt.Sprintf("%v,1", math.MaxFloat64),
			want:    math.MaxFloat64,
			wantErr: false,
		},
		{
			name:    "more than two digits",
			input:   "4,3,2",
			want:    24,
			wantErr: false,
		},
		{
			name:    "negative numbers",
			input:   "4,-2",
			want:    -8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := multiply(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) && (*got != tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:  "simple int division",
			input: "3,2",
			want:  1.5,
		},
		{
			name:  "simple int division whole number answer",
			input: "10,2",
			want:  5,
		},
		{
			name:  "three number division",
			input: "10,2,5",
			want:  1,
		},
		{
			name:  "single number input",
			input: "10",
			want:  10,
		},
		{
			name:    "divide by zero",
			input:   "10,0",
			wantErr: true,
		},
		{
			name:  "single number input trailing with a comma",
			input: "10,",
			want:  10,
		},
		{
			name:  "multiple numbers with starting zero",
			input: "0,10,5,8,4",
			want:  0,
		},
		{
			name:    "multiple numbers with a zero inbetween",
			input:   "6,2,0,10,5,8,4",
			wantErr: true,
		},
		{
			name:    "divide numbers from less than 1",
			input:   fmt.Sprintf("%v,0.05", math.MaxFloat64),
			wantErr: true,
		},
		{
			name:    "invalid input",
			input:   "0,&",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) && *got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
