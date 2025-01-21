package main

import (
	"testing"
)

// TODO: Right test cases to test all possible branches

func TestGetDestinationMatchLeft(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 10,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      5,
			DestinationStart: 10,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 15 {
		t.Fatalf("incorrect origin %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 5 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapMatchRight(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 10,
		RangeLength: 10,
	}

	toInput := []MapEntry{
		{
			OriginStart:      8,
			DestinationStart: 12,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 14 {
		t.Fatalf("incorrect origin %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 8 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapMatchPerfect(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 5,
		RangeLength: 5,
	}

	toInput := []MapEntry{
		{
			OriginStart:      5,
			DestinationStart: 20,
			RangeLength:      5,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 20 {
		t.Fatalf("incorrect origin %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 5 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}
}

func TestGetDestinationMapMatchMiddle(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 0,
		RangeLength: 20,
	}

	toInput := []MapEntry{
		{
			OriginStart:      5,
			DestinationStart: 20,
			RangeLength:      5,
		},
	}

	/* Should include
	{
		OriginStart: 0,
		RangeLength: 5
	}
	*/
	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 20 {
		t.Fatalf("incorrect origin %d", actual[0].OriginStart)
	}

	if actual[0].RangeLength != 5 {
		t.Fatalf("incorrect range length %d", actual[0].RangeLength)
	}

	if len(actual) != 3 {
		t.Fatalf("incorrect  length %d", actual[0].RangeLength)
	}

}

func TestGetDestinationMapMatchMultiple(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 0,
		RangeLength: 100,
	}

	toInput := []MapEntry{
		{
			OriginStart:      5,
			DestinationStart: 20,
			RangeLength:      10,
		},
		{
			OriginStart:      40,
			DestinationStart: 55,
			RangeLength:      10,
		},
		{
			OriginStart:      200,
			DestinationStart: 215,
			RangeLength:      10,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if len(actual) != 5 {
		t.Fatalf("incorrect length %d", len(actual))
	}

	if actual[0].OriginStart != 20 {
		t.Fatalf("incorrect origin %d", actual[0].OriginStart)
	}

	if actual[1].OriginStart != 55 {
		t.Fatalf("incorrect origin %d", actual[1].OriginStart)
	}
}

func TestGetDestinationMapNoMatchLeft(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 10,
		RangeLength: 10,
	}

	toInput := []MapEntry{
		{
			OriginStart:      0,
			DestinationStart: 12,
			RangeLength:      2,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 10 {
		t.Fatalf("incorrect length %d", len(actual))
	}
}

func TestGetDestinationMapNoMatchRight(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 10,
		RangeLength: 10,
	}

	toInput := []MapEntry{
		{
			OriginStart:      100,
			DestinationStart: 12,
			RangeLength:      2,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 10 {
		t.Fatalf("incorrect length %d", len(actual))
	}
}

func TestGetDestinationMapRangeOne(t *testing.T) {

	fromInput := MapEntry{
		OriginStart: 1,
		RangeLength: 1,
	}

	toInput := []MapEntry{
		{
			OriginStart:      1,
			DestinationStart: 12,
			RangeLength:      2,
		},
	}

	actual := getDestinationMap(fromInput, toInput)

	if actual[0].OriginStart != 12 {
		t.Fatalf("incorrect length %d", len(actual))
	}

	if actual[0].RangeLength != 1 {
		t.Fatalf("incorrect length %d", len(actual))
	}
}
