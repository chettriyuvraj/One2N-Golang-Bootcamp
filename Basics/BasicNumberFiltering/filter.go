package main

/**** Filters ****/

func FilterEven(inp []int) []int {
	return FilterByFuncs(inp, isEven)
}

func FilterOdd(inp []int) []int {
	return FilterByFuncs(inp, isOdd)
}

func FilterPrime(inp []int) []int {
	return FilterByFuncs(inp, isPrime)
}

/**** Composite Filters ****/

func FilterOddPrime(inp []int) []int {
	return FilterByFuncs(inp, isOdd, isPrime)
}

func FilterEvenMultiplesOfFive(inp []int) []int {
	return FilterByFuncs(inp, isEven, isMultipleOfFive)
}

func FilterOddMultiplesOfThreeGreaterThanTen(inp []int) []int {
	return FilterByFuncs(inp, isOdd, isMultipleOfThree, isGreaterThanTen)
}

/**** Filters not technically needed, only for testing ****/

func FilterMultipleOfFive(inp []int) []int {
	return FilterByFuncs(inp, isMultipleOfFive)
}

func FilterMultipleOfThree(inp []int) []int {
	return FilterByFuncs(inp, isMultipleOfThree)
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

func isGreaterThanTen(num int) bool {
	return num > 10
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

func FilterByFuncs(inp []int, fList ...func(n int) bool) []int {
	res := []int{}
	for _, num := range inp {
		satisfiesAllFuncs := true
		for _, f := range fList {
			if !f(num) {
				satisfiesAllFuncs = false
				break
			}
		}
		if satisfiesAllFuncs {
			res = append(res, num)
		}
	}
	return res
}
