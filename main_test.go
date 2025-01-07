package main

import (
	"testing"
)

func TestGetDestinationValue(t *testing.T) {

	actual1 := getDestinationMap(
		MapEntry{
			OriginStart: 0,
			RangeLength: 10,
		},
		MapEntry{
			OriginStart:      5,
			DestinationStart: 10,
			RangeLength:      10,
		})

	if actual1.OriginStart != 10 {
		t.Error("Nah playah")
	}

	if actual1.RangeLength != 5 {
		t.Error("Naaaah")
	}
}

func TestGetDestinationValueSingleMatch(t *testing.T) {

	// Need to double check that these tests are accurate.
	actual := getDestinationMap(
		MapEntry{
			OriginStart: 0,
			RangeLength: 10,
		},
		MapEntry{
			OriginStart:      9,
			DestinationStart: 30,
			RangeLength:      10,
		})

	if actual.OriginStart != 30 {
		t.Error("Nah playah")
	}

	if actual.RangeLength != 1 {
		t.Error("Naaaah")
	}
}

func TestGetDestinationValueNoMatch(t *testing.T) {

	// Need to double check that these tests are accurate.
	actual := getDestinationMap(
		MapEntry{
			OriginStart: 0,
			RangeLength: 10,
		},
		MapEntry{
			OriginStart:      20,
			DestinationStart: 30,
			RangeLength:      10,
		})

	if actual.OriginStart != 0 {
		t.Errorf("Nah playah %d", actual.OriginStart)
	}

	if actual.RangeLength != 1 {
		t.Error("Naaaah")
	}
}
