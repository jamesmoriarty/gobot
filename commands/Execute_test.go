package commands

import (
	"reflect"
	"testing"

	"../directions"
	"../robots"
)

func TestExecuteCommandPlace(t *testing.T) {
	tokens := []string{"Place", "1", "1", "West"}
	got, err := Execute(tokens, nil)

	if !reflect.DeepEqual(*got, robots.Robot{1, 1, directions.West}) {
		t.Errorf("*%v != ", got)
	}

	if err != nil {
		t.Errorf("*%v != nil", err)
	}
}

func TestExecuteCommandMove(t *testing.T) {
	tokens := []string{"Move"}
	got, err := Execute(tokens, &robots.Robot{1, 1, directions.West})

	if !reflect.DeepEqual(*got, robots.Robot{0, 1, directions.West}) {
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
