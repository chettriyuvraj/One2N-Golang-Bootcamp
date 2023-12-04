package main

import "testing"

func TestStrategy(t *testing.T) {

	tc := []struct {
		turnValues   []int /* List of dice values encountered in a given turn */
		turnTotals   []int /* Cumulative sum of turn values encountered so far */
		totals       []int /* Overall total of all turns */
		strategyfunc Strategy
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

func TestParseRange(t *testing.T) {
	tests := []struct {
		tc   []string
		want [][]int
	}{
		{[]string{"10", "15"}, [][]int{{10}, {15}}},
		{[]string{"10-100", "15"}, [][]int{{10, 100}, {15}}},
		{[]string{"1-2", "3-15"}, [][]int{{1, 2}, {3, 15}}},
		{[]string{"3-2", "3"}, [][]int{}},
		{[]string{"2-2", "3"}, [][]int{}},
	}

	for _, test := range tests {
		got, _ := ParseHoldCounts(test.tc)
		if !CompareIntSlices(got, test.want) {
			t.Errorf("Error parsing holdcount from stdin, got %v want %v", got, test.want)
		}
	}
}
