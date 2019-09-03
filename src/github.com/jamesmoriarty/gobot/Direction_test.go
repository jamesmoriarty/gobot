package main

import (
	"testing"
)

func TestNorthStringToDirection(t *testing.T) {
	got, _ := StringToDirection("North")

	if *got != North {
		t.Errorf("*%v != North", got)
	}
}

func TestEastStringToDirection(t *testing.T) {
	got, _ := StringToDirection("East")

	if *got != East {
		t.Errorf("*%v != East", got)
	}
}

func TestSouthStringToDirection(t *testing.T) {
	got, _ := StringToDirection("South")

	if *got != South {
		t.Errorf("*%v != South", got)
	}
}

func TestWestStringToDirection(t *testing.T) {
	got, _ := StringToDirection("West")

	if *got != West {
		t.Errorf("*%v != West", got)
	}
}
