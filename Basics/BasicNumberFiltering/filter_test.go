package main

import "testing"

func TestFilter(t *testing.T) {
	sample := []int{1, 2, 3, 4}
	t.Run("test filtering even numbers from a list", func(t *testing.T) {
		got := FilterEven(sample)
		want := []int{2, 4}

		if !CompareSlicesInt(got, want) {
			t.Errorf("filter even numbers from %v; wanted %v; got %v", sample, want, got)
		}
	})

	t.Run("test filtering odd numbers from a list", func(t *testing.T) {
		got := FilterOdd(sample)
		want := []int{1, 3}
		if !CompareSlicesInt(got, want) {
			t.Errorf("filter odd numbers from %v; wanted %v; got %v", sample, want, got)
		}
	})

}
