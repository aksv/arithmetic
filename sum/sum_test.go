package arithmetic_test

import (
	arithmetic "github.com/aksv/arithmetic/v2/sum"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "It should sum to numbers",
			args: args{a: 2, b: 2, c: 3},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmetic.Add(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
