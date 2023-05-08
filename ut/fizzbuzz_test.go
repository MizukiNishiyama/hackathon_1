package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "-5~20の値でFizzBuzzテスト",
			args: args{
				numbers: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			want: []string{"-5", "-4", "-3", "-2", "-1", "0", "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
