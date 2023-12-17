package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("error opening input file", err)
	}
	defer input.Close()

	finalValues := []int{}
	reader := bufio.NewReader(input)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln("error reading line from file: ", err)
		}

		mappedDigits := findDigits(line)
		mappedStrDigits := strDigits(line)

		if len(mappedDigits) == 0 && len(mappedStrDigits) == 0 {
			continue
		}

		minDigitIndex, maxDigitIndex := getMaxAndMin(mappedDigits)
		minStrDigitIndex, maxStrDigitIndex := getMaxAndMin(mappedStrDigits)

		first := ""
		last := ""

		if minStrDigitIndex < minDigitIndex {
			first = mappedStrDigits[minStrDigitIndex]
		} else {
			first = mappedDigits[minDigitIndex]
		}

		if maxStrDigitIndex > maxDigitIndex {
			last = mappedStrDigits[maxStrDigitIndex]
		} else {
			last = mappedDigits[maxDigitIndex]
		}

		value, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalln("error joining digits into a value: ", err)
		}

		finalValues = append(finalValues, value)
	}

	result := 0
	for _, v := range finalValues {
		result += v
	}

	fmt.Println("Result: ", result)
}

func getMaxAndMin(mappedValues map[int]string) (int, int) {
	maxValue := math.MinInt
	minValue := math.MaxInt

	if len(mappedValues) == 0 {
		return minValue, maxValue
	}

	digitIndexes := maps.Keys(mappedValues)
	return slices.Min(digitIndexes), slices.Max(digitIndexes)
}

func strDigits(line string) map[int]string {
	strDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	srtDigitsFound := map[int]string{}
	for _, strDigit := range strDigits {
		first := strings.Index(line, strDigit)
		if first == -1 {
			continue
		}

		last := strings.LastIndex(line, strDigit)

		srtDigitsFound[first] = converStrDigitToDigit(strDigit)
		srtDigitsFound[last] = converStrDigitToDigit(strDigit)
	}

	return srtDigitsFound
}

func findDigits(line string) map[int]string {
	digits := map[int]string{}
	for i, c := range line {
		if unicode.IsDigit(c) {
			digits[i] = string(c)
		}
	}

	return digits
}

func converStrDigitToDigit(strDigit string) string {
	switch strDigit {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}
