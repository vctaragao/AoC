package main

type Color int

const (
	Blue Color = iota
	Green
	Red
)

func NewColor(color string) Color {
	switch color {
	case "blue":
		return Blue
	case "green":
		return Green
	case "red":
		return Red
	default:
		return 0
	}
}

func (c Color) String() string {
	return []string{"blue", "green", "red"}[c]
}
