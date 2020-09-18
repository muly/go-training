package main

import (
	"testing"
)

// func Test_xxx(t *testing.T) {
// 	// -ve input
// 	a, b := -5, -10
// 	expected := -15
// 	actual := add(a, b)
// 	if actual != expected {
// 		t.Errorf("for -ve input case: wanted %d, but got %d", expected, actual)
// 	}

// }

// integer overflow:
// -ve input: -5,-10 -> -15
// 0 inputs: 1,0 -> 1
// result is 0: -5,5 -> 0

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-ve input",
			args: args{a:-5 ,
				b:-10},
				want: -15,
		},{
			name: "0 inputs",
			args: args{a:1 ,
				b:0},
				want: 1,
		},{
			name: "result is 0",
			args: args{a:-5 ,
				b:5},
				want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Benchmark_add(b *testing.B){
	// b.N=5
	for i:=0; i<= b.N; i++{
		add(3,4)
	}
}


// pkg: github.com/muly/go-training/20200914/4/unittests
// Benchmark_add-8   	1000000000	         0.299 ns/op	       0 B/op	       0 allocs/op