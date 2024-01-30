package calendar

import (
	"time"
)

type Model struct {
	t time.Time
}

func New() Model {
	return Model{
		t: time.Now(),
	}
}
