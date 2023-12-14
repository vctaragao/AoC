package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("error opening input file", err)
	}
	defer input.Close()

	finalValues := []int64{}
	reader := bufio.NewReader(input)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln("error reading line from file: ", err)
		}

		last := ""
		first := ""
		digits := []string{}

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if first == "" {
					first = string(line[i])
					digits = append(digits, string(line[i]))
				}

				last = string(line[i])
			}
		}

		digits = append(digits, last)

		// join the first digit and the last digit into a single value
		value, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			log.Fatalln("error joning digits into a value: ", err)
		}

		finalValues = append(finalValues, int64(value))
	}

	// add those values to get the final answer
	result := int64(0)
	for _, v := range finalValues {
		result += v
	}

	fmt.Println("Result: ", result)
}
