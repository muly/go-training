package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "happy path test",
			a: 1, b: 2, want: 3,
		},
		{name: "0 vale test",
			a: 1, b: -1, want: 0,
		},
		{name: "-ve  path test",
			a: 1, b: -4, want: -3,
		},
	}
	for _, tt := range tests {
		if got := add(tt.a, tt.b); got != tt.want {
			t.Errorf("add() test case %s = %v, want %v", tt.name, got, tt.want)
		}

	}

}

// test data:
// : 1, 2: 3
// : 1, -1, 0
// -ve vale test: 1, -4, -3

func Test_multiple(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiple(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
