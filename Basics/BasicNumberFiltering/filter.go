package main

func FilterEven(inp []int) []int {
	return FilterByFunc(inp, isEven)
}

func FilterOdd(inp []int) []int {
	return FilterByFunc(inp, isOdd)
}

func FilterPrime(inp []int) []int {
	return FilterByFunc(inp, isPrime)
}

func FilterOddPrime(inp []int) []int {
	return FilterByFunc(inp, isOddPrime)
}

func FilterEvenMultipleOfFive(inp []int) []int {
	return FilterByFunc(inp, isEvenMultipleOfFive)
}

/**** Helpers ****/

func isEven(num int) bool {
	return (num & 1) == 0
}

func isOdd(num int) bool {
	return !isEven(num)
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i <= num/2; i++ { // You can even go only till sqrt(num)
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isMultipleOfFive(num int) bool {
	return num%5 == 0
}

/**** Composite Helpers ****/

func isOddPrime(num int) bool {
	return isOdd(num) && isPrime(num)
}

func isEvenMultipleOfFive(num int) bool {
	return isEven(num) && isMultipleOfFive(num)
}

/**** Misc Helpers ****/

func CompareSlicesInt(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, n1 := range slice1 {
		n2 := slice2[i]
		if n1 != n2 {
			return false
		}
	}
	return true
}

func FilterByFunc(inp []int, f func(n int) bool) []int {
	res := []int{}
	for _, num := range inp {
		if f(num) {
			res = append(res, num)
		}
	}
	return res
}
