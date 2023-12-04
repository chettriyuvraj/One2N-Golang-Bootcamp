package main

import "fmt"

/* Custom Strategies */
var holdAtTenStrategy Strategy = getHoldAtValueStrategy(10)
var holdAfterTwoFivesStrategy Strategy = Strategy{desc: "Hold only after two fives are rolled",
	strategyfunc: func(turnValues []int, turnTotal int, total int) bool {
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
	},
}

/***** Helpers *****/
func getHoldAtValueStrategy(holdValue int) Strategy {
	return Strategy{desc: fmt.Sprintf("Hold at or after %d", holdValue),
		strategyfunc: func(turnValues []int, turnTotal int, total int) bool {
			if turnTotal >= holdValue {
				return true
			}
			return false
		},
	}
}
