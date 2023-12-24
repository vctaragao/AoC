package main

import (
	"strconv"
	"strings"
)

type Event struct {
	Size  int
	Color Color
}

func NewEvent(strEvent string) Event {
	d := strings.Split(strings.TrimSpace(strEvent), " ")

	size, err := strconv.Atoi(d[0])
	if err != nil {
		panic(err)
	}

	return Event{
		Size:  size,
		Color: NewColor(d[1]),
	}
}

func (e Event) GetColor() Color {
	return e.Color
}

func (e *Event) IsPossible(maxSize int) bool {
	return e.Size <= maxSize
}
