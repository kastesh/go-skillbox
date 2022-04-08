package main

import "testing"

func TestgetEvenAndOdd(t *testing.T) {
	b := [10]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}
	evenB, oddB := getEvenAndOdd(b)
	if evenB != 4 && oddB != 6 {
		t.Error("Expected even=4 and odd=6; got ", evenB, " and ", oddB, ".")
	}
}
