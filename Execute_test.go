package main

import (
	"reflect"
	"testing"
)

func TestExecuteCommandPlace(t *testing.T) {
	tokens := []string{"place", "1", "1", "West"}
	got, err := Execute(tokens, nil)

	if !reflect.DeepEqual(*got, Robot{1, 1, West}) {
		t.Errorf("*%v != ", got)
	}

	if err != nil {
		t.Errorf("*%v != nil", err)
	}
}

func TestExecuteError(t *testing.T) {
	tokens := []string{"zoo"}
	got, err := Execute(tokens, nil)

	if got != nil {
		t.Errorf("*%v != nil", got)
	}

	if err.Error() != "invalid command" {
		t.Errorf("*%v != nil", err)
	}
}
