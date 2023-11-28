package main

func FilterEven(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if (num & 1) == 0 {
			res = append(res, num)
		}
	}
	return res
}

func FilterOdd(inp []int) []int {
	res := []int{}
	for _, num := range inp {
		if (num & 1) != 0 {
			res = append(res, num)
		}
	}
	return res
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
