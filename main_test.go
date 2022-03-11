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
			name:    "playing with boundaries",
			input:   fmt.Sprintf("1,%v", math.MaxFloat64/100),
			want:    math.MaxFloat64/100 + 1,
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
			name:  "going beloew minFloat64",
			input: "-40,20,-30",
			want:  -30.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subtract(tt.input)
			assert.Equal(t, err != nil, tt.wantErr)
			if err == nil {
				fmt.Printf("Value of *got is %v\n", *got)
				assert.Equal(t, tt.want, *got, "these should be equal")
			}
		})
	}
}
