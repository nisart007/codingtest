package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	PrintNumbers()
}

func TestAreaOfTraingleSuccess(t *testing.T) {
	var expected float64 = 25
	actual := AreaOfTraingle(5, 10)
	if actual != expected {
		t.Errorf("Area Of Traingle, actual: %f, expected: %f.", actual, expected)
	}
}
