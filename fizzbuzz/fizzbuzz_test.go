package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("test function FizzBuzz case input [1] should be return 1", func(t *testing.T) {
		input := []int{1}
		got := FizzBuzz(input)
		want := "1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [1,2] should be return 1,2", func(t *testing.T) {
		input := []int{1, 2}
		got := FizzBuzz(input)
		want := "1,2"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [1,1,1,1,1,1,1] should be return 1,1,1,1,1,1,1", func(t *testing.T) {
		input := []int{1, 1, 1, 1, 1, 1, 1}
		got := FizzBuzz(input)
		want := "1,1,1,1,1,1,1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [1,2,3,4,5,6,7,8,9,10] should be return 1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := FizzBuzz(input)
		want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [11,12,13,14,15] should be return 11,Fizz,13,14,FizzBuzz", func(t *testing.T) {
		input := []int{11, 12, 13, 14, 15}
		got := FizzBuzz(input)
		want := "11,Fizz,13,14,FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [100,101,102,103,104,105] should be return Buzz,Fizz,13,14,FizzBuzz", func(t *testing.T) {
		input := []int{100, 101, 102, 103, 104, 105}
		got := FizzBuzz(input)
		want := "Buzz,101,Fizz,103,104,FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [3,5,6,9,10,15] should be return Fizz,Buzz,Fizz,Fizz,Buzz,FizzBuzz", func(t *testing.T) {
		input := []int{3, 5, 6, 9, 10, 15}
		got := FizzBuzz(input)
		want := "Fizz,Buzz,Fizz,Fizz,Buzz,FizzBuzz"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [] should be return Empty Array", func(t *testing.T) {
		input := []int{}
		got := FizzBuzz(input)
		want := "Empty Array"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [0,1,2] should be return Error", func(t *testing.T) {
		input := []int{0, 1, 2}
		got := FizzBuzz(input)
		want := "Error"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("test function FizzBuzz case input [-1,2,3] should be return Error", func(t *testing.T) {
		input := []int{-1, 2, 3}
		got := FizzBuzz(input)
		want := "Error"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
