package main

func FilterEven(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if isEven(num) {
			res = append(res, num)
		}
	}
	return res
}

func FilterOdd(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if isOdd(num) {
			res = append(res, num)
		}
	}
	return res
}

func FilterPrime(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if isPrime(num) {
			res = append(res, num)
		}
	}
	return res
}

func FilterOddPrime(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if isPrime(num) && isOdd(num) {
			res = append(res, num)
		}
	}
	return res
}

func FilterEvenMultipleOfFive(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if isEven(num) && isMultipleOfFive(num) {
			res = append(res, num)
		}
	}
	return res
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
