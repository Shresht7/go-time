package calendar

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
)

type Model struct {
	t time.Time

	keys KeyMaps
	help help.Model
}

func New() Model {
	return Model{
		t:    time.Now(),
		keys: DefaultKeyMaps,
		help: help.New(),
	}
}
