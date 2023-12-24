package main

import (
	"strconv"
	"strings"
)

type Game struct {
	Number int
	Draws  []Draw
}

func NewGame(gameStr string) Game {
	strGame := strings.Split(gameStr, ":")

	game := Game{
		Number: parseNumber(strGame[0]),
		Draws:  parseDraws(strGame[1]),
	}

	return game
}

func parseNumber(strGame string) int {
	gameInfo := strings.Split(strGame, " ")
	number, err := strconv.Atoi(gameInfo[1])
	if err != nil {
		panic(err)
	}

	return number
}

func parseDraws(drawsStr string) []Draw {
	strDraws := strings.Split(strings.TrimSpace(drawsStr), ";")

	draws := make([]Draw, 0, len(strDraws))
	for _, strDraw := range strDraws {
		draws = append(draws, NewDraw(strDraw))
	}

	return draws
}

func (g *Game) IsPossible(maxRedSize, maxBlueSize, maxGreenSize int) bool {
	for _, draw := range g.Draws {
		if !draw.IsPossible(maxRedSize, maxBlueSize, maxGreenSize) {
			return false
		}
	}

	return true
}
