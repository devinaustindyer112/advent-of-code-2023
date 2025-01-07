package main

import (
	"testing"
)

func TestNothing(t *testing.T) {

	actual := MapEntry{}

	if actual.OriginStart != 0 {
		t.Error("Nah playah")
	}

	if actual.RangeLength != 0 {
		t.Error("Naaaah")
	}
}
