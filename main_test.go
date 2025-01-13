package main

import (
	"testing"
)

// There is a perfect overlap

// The destination origin matches with left side of origin origin
// The desination origin mathces with right side of the origin origin
// No match. Destination origins are left side
// No match. Destination origins are right side

func TestGetDestinationMapLeft(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 0,
		RangeLength: 10,
	}

	toInput := []MapEntry{
		{
			OriginStart:      0,
			DestinationStart: 10,
			RangeLength:      5,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 10 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 5 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapSingleRight(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 2 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapSinglePerfect(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 2 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapSingleMiddle(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 2 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapNoMatchLeft(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 2 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapNoMatchRight(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect origin: %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 2 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}
