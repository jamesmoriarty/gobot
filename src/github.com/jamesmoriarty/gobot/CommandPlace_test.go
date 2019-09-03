package main

import "testing"

func TestCommandPlace(t *testing.T) {
	tokens := []string{"place", "1", "2", "West"}
	got, err := CommandPlace(tokens)

	if err != nil {
		t.Errorf("*%v != nil", err)
	}

	if got == nil {
		t.Errorf("*%v != nil", got)
	}
}

func TestCommandPlaceXError(t *testing.T) {
	tokens := []string{"place", "zoo", "2", "West"}
	got, err := CommandPlace(tokens)

	if got != nil {
		t.Errorf("*%v != nil", got)
	}

	if err.Error() != "invalid x" {
		t.Errorf("*%v != err", err)
	}
}

func TestCommandPlaceYError(t *testing.T) {
	tokens := []string{"place", "1", "y", "West"}
	got, err := CommandPlace(tokens)

	if got != nil {
		t.Errorf("*%v != nil", got)
	}

	if err.Error() != "invalid y" {
		t.Errorf("*%v != err", err)
	}
}

func TestCommandPlaceDirectionError(t *testing.T) {
	tokens := []string{"place", "1", "1", "zoo"}
	got, err := CommandPlace(tokens)

	if got != nil {
		t.Errorf("*%v != nil", got)
	}

	if err.Error() != "invalid direction" {
		t.Errorf("*%v != err", err)
	}
}
