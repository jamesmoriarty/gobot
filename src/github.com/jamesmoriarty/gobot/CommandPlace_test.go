package main

import "testing"

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
