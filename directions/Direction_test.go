package directions

import (
	"testing"
)

func TestNorthDirectionToString(t *testing.T) {
	got := North.String()

	if got != "North" {
		t.Errorf("*%v != North", got)
	}
}

func TestStringToDirectionNorth(t *testing.T) {
	got, _ := StringToDirection("North")

	if *got != North {
		t.Errorf("*%v != North", got)
	}
}

func TestStringToDirectionEast(t *testing.T) {
	got, _ := StringToDirection("East")

	if *got != East {
		t.Errorf("*%v != East", got)
	}
}

func TestStringToDirectionSouth(t *testing.T) {
	got, _ := StringToDirection("South")

	if *got != South {
		t.Errorf("*%v != South", got)
	}
}

func TestStringToDirectionWest(t *testing.T) {
	got, _ := StringToDirection("West")

	if *got != West {
		t.Errorf("*%v != West", got)
	}
}

func TestStringToDirectionError(t *testing.T) {
	got, err := StringToDirection("zoo")

	if got != nil {
		t.Errorf("*%v != nil", got)
	}

	if err.Error() != "invalid direction" {
		t.Errorf("*%v != err", err)
	}
}
