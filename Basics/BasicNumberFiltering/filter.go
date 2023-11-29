package main

/**** Filters ****/

func FilterEven(inp []int) []int {
	return FilterByFunc(inp, isEven)
}

func FilterOdd(inp []int) []int {
	return FilterByFunc(inp, isOdd)
}

func FilterPrime(inp []int) []int {
	return FilterByFunc(inp, isPrime)
}

/**** Filters not technically needed, only for testing ****/

func FilterMultipleOfFive(inp []int) []int {
	return FilterByFunc(inp, isMultipleOfFive)
}

func FilterMultipleOfThree(inp []int) []int {
	return FilterByFunc(inp, isMultipleOfThree)
}

/**** Composite Filters ****/

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

func isMultipleOfThree(num int) bool {
	return num%3 == 0
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
