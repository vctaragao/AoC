package main

import (
	"strings"
)

type Draw struct {
	Events []Event
}

func NewDraw(strDraws string) Draw {
	strEvents := strings.Split(strings.TrimSpace(strDraws), ",")

	events := make([]Event, 0, len(strEvents))
	for _, strEvent := range strEvents {
		events = append(events, NewEvent(strEvent))
	}

	return Draw{
		Events: events,
	}
}

func (d *Draw) IsPossible(maxRedSize, maxBlueSize, maxGreenSize int) bool {
	for _, event := range d.Events {
		switch event.GetColor() {
		case Blue:
			if !event.IsPossible(maxBlueSize) {
				return false
			}
		case Red:
			if !event.IsPossible(maxRedSize) {
				return false
			}
		case Green:
			if !event.IsPossible(maxGreenSize) {
				return false
			}
		}
	}

	return true
}
