package main

import "testing"

func TestFilterEven(t *testing.T) {
	sample := []int{1, 2, 3, 4}
	got := FilterEven(sample)
	want := []int{2, 4}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter even numbers from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestFilterOdd(t *testing.T) {
	sample := []int{1, 2, 3, 4}
	got := FilterOdd(sample)
	want := []int{1, 3}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter odd numbers from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestFilterPrime(t *testing.T) {
	sample := []int{1, 2, 3, 5, 7, 9, 10, 13}
	got := FilterPrime(sample)
	want := []int{2, 3, 5, 7, 13}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter odd numbers from %v; wanted %v; got %v", sample, want, got)
	}
}
