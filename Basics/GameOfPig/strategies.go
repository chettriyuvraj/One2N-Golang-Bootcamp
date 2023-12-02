package main

/* Custom Strategies */
var holdAtTenStrategy Strategy = getHoldAtValueStrategy(10)
var holdAfterTwoFivesStrategy Strategy = func(turnValues []int, turnTotal int, total int) bool {
	fiveCount := 0
	for _, turnValue := range turnValues {
		if turnValue == 5 {
			fiveCount++
		}
		if fiveCount == 2 {
			return true
		}
	}
	return false
}

/***** Helpers *****/
func getHoldAtValueStrategy(holdValue int) Strategy {
	return func(turnValues []int, turnTotal int, total int) bool {
		if turnTotal >= holdValue {
			return true
		}
		return false
	}
}
