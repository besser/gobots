package main

import "testing"

var testCasesIsNeighbour = []struct {
	x        int
	y        int
	expected bool
}{
	{1, 1, false},
	{0, 1, true},
	{1, 0, true},
	{3, 4, false},
}

func TestIsNeighbour(t *testing.T) {
	player := position{1, 1}

	for _, test := range testCasesIsNeighbour {
		observed := player.isNeighbour(test.x, test.y)
		if observed != test.expected {
			t.Fatalf("The result expected of test was '%t' for PlayerXY(%d,%d) and the new position XY(%d,%d)!",
				test.expected, player.x, player.y, test.x, test.y)
		}
	}
}
