package commands

import (
	"reflect"
	"testing"

	"../directions"
	"../robots"
)

func TestCommandMove(t *testing.T) {
	tokens := []string{"Move"}
	got, err := CommandMove(tokens,  &robots.Robot{1, 1, directions.West})

	if !reflect.DeepEqual(*got, robots.Robot{0, 1, directions.West}) {
		t.Errorf("*%v != ", got)
	}

	if err != nil {
		t.Errorf("*%v != nil", err)
	}
}