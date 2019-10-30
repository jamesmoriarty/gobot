package robots

import "github.com/jamesmoriarty/gobot/directions"

type Robot struct {
	X         int64
	Y         int64
	Direction directions.Direction
}
