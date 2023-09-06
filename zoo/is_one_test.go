package main

import (
	"testing"
)

func TestIsOne(t *testing.T) {
	n := 1
	b := IsOne(n)
	if b != true {
		t.Errorf("%d is not one.", n)
	}
}

func TestIsOneFalse(t *testing.T) {
	n := 100
	b := IsOne(n)
	if b == true {
		t.Errorf("%d is not one.", n)
	}
}
