package robots

import (
	"testing"

	"../directions"
)

func TestRobotX(t *testing.T) {
	got := Robot{-1, 2, directions.West}.X

	if got != -1 {
		t.Errorf("*%v != X", got)
	}
}

func TestRobotY(t *testing.T) {
	got := Robot{-1, 2, directions.West}.Y

	if got != 2 {
		t.Errorf("*%v != Y", got)
	}
}

func TestRobotDirection(t *testing.T) {
	got := Robot{-1, 2, directions.West}.Direction

	if got != directions.West {
		t.Errorf("*%v != Direction", got)
	}
}
