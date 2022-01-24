package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum a collection of 5 numbers", func(t *testing.T) 	{
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)

		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
	t.Run("sum using as agument a collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got: %d want: %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum the elements of 2 slices into one new slice", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3,9}

		if  !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

}

func TestSumAllTalis(t *testing.T) {
	  CheckSumsAfterTest := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	}
	t.Run ("Sum the tals of slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2,4}, []int{0,9,1})
		want := []int{6,10}
	
		CheckSumsAfterTest(t, got, want)
		
	})
	t.Run("passing an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0, 9}
		CheckSumsAfterTest(t, got, want)
		
	})
	
}



