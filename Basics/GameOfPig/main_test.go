package main

import "testing"

func TestStrategy(t *testing.T) {

	/* Define two custom strategies */
	holdAtTenStrategy := func(turnValues []int, turnTotal int, total int) bool {
		if turnTotal >= 10 {
			return true
		}
		return false
	}

	holdAfterTwoFivesStrategy := func(turnValues []int, turnTotal int, total int) bool {
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

	/* Test drive */
	tc := []struct {
		turnValues   []int /* List of dice values encountered in a given turn */
		turnTotals   []int /* Cumulative sum of turn values encountered so far */
		totals       []int /* Overall total of all turns */
		strategyfunc func(turnValues []int, turnTotal int, total int) bool
		want         []bool
	}{
		{
			[]int{1, 5, 3, 6},
			[]int{1, 6, 9, 15},
			[]int{32, 37, 40, 46},
			holdAtTenStrategy,
			[]bool{false, false, false, true},
		},

		{
			[]int{1, 5, 3, 5},
			[]int{1, 6, 9, 15},
			[]int{32, 37, 40, 46},
			holdAfterTwoFivesStrategy,
			[]bool{false, false, false, true},
		},
	}

	for _, test := range tc {
		for i, _ := range test.turnValues {
			curTurnValues := test.turnValues[:i+1]
			got := test.strategyfunc(curTurnValues, test.turnTotals[i], test.totals[i])
			want := test.want[i]
			if got != want {
				t.Errorf("strategy func error: got %v, want %v", got, want)
			}
		}
	}

}
