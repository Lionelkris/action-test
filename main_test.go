package main

import (
	"fmt"
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
		}, {
			name:    "passing float values as input",
			input:   "13.9,41.2",
			want:    55.1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add(tt.input)
			if err == nil {
				fmt.Printf("Value of *got is %v\n", *got)
				assert.Equal(t, *got, tt.want, "these should be equal")
			}
			assert.Equal(t, err != nil, tt.wantErr)

		})
	}
}
