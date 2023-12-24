package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatalln("error opening input file", err)
	}
	defer f.Close()

	possibleGames := []Game{}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln("error reading line", err)
		}

		game := NewGame(line)
		fmt.Println(game)
		fmt.Println(game.IsPossible(RedCubes, BlueCubes, GreenCubes))

		if game.IsPossible(RedCubes, BlueCubes, GreenCubes) {
			possibleGames = append(possibleGames, game)
		}

		totalGameNumber := 0
		for _, game := range possibleGames {
			totalGameNumber += game.Number
		}

		fmt.Println(totalGameNumber)
	}
}
