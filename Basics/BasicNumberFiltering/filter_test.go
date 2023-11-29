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
		t.Errorf("filter prime numbers from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestFilterMultipleOfFive(t *testing.T) {
	sample := []int{1, 2, 3, 5, 7, 9, 10, 13}
	got := FilterMultipleOfFive(sample)
	want := []int{5, 10}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter multiples of five from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestFilterMultipleOfThree(t *testing.T) {
	sample := []int{1, 2, 3, 5, 7, 9, 10, 13}
	got := FilterMultipleOfThree(sample)
	want := []int{3, 9}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter multiples of three from %v; wanted %v; got %v", sample, want, got)
	}
}

/**** Composite Tests ****/

func TestFilterPrimeOdd(t *testing.T) {
	sample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	got := FilterOddPrime(sample)
	want := []int{3, 5, 7, 11}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter odd prime numbers from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestFilterEvenMultiplesOfFive(t *testing.T) {
	sample := []int{2, 5, 10, 15, 20}
	got := FilterEvenMultiplesOfFive(sample)
	want := []int{10, 20}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter even multiple of five from %v; wanted %v; got %v", sample, want, got)
	}
}

func TestOddMultiplesOfThreeGreaterThanTen(t *testing.T) {
	sample := []int{2, 5, 10, 15, 21, 32}
	got := FilterOddMultiplesOfThreeGreaterThanTen(sample)
	want := []int{15, 21}

	if !CompareSlicesInt(got, want) {
		t.Errorf("filter odd multiples of three greater than ten from %v; wanted %v; got %v", sample, want, got)
	}
}

/**** Match All and Match Any Tests ****/

func TestFilterByFuncsMatchAll(t *testing.T) {
	t.Run("match all numbers greater than 10 and prime", func(t *testing.T) {
		sample := []int{2, 5, 10, 11, 23, 15, 21, 32}
		got := FilterByFuncsMatchAll(sample, isGreaterThanTen, isPrime)
		want := []int{11, 23}

		if !CompareSlicesInt(got, want) {
			t.Errorf("filter numbers greater than 10 AND prime from %v; wanted %v; got %v", sample, want, got)
		}
	})
}

func TestFilterByFuncsMatchAny(t *testing.T) {
	t.Run("match all numbers greater than 10 or prime", func(t *testing.T) {
		sample := []int{1, 4, 8, 5, 10, 11, 23, 15, 21, 32}
		got := FilterByFuncsMatchAny(sample, isGreaterThanTen, isPrime)
		want := []int{5, 11, 23, 15, 21, 32}

		if !CompareSlicesInt(got, want) {
			t.Errorf("filter numbers greater than 10 OR prime from %v; wanted %v; got %v", sample, want, got)
		}
	})
}
