package main

import "testing"

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
