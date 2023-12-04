package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type coordinate struct {
	number int
	y      int
	xStart int
	xEnd   int
}
type symbol struct {
	symbol string
	y      int
	x      int
}

func isDigit(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func isSymbol(r rune) bool {
	return !isDigit(r) && !isDot(r)
}

func isDot(r rune) bool {
	return r == '.'
}

func isAdjacent(coord coordinate, symbols []symbol) bool {
	startX := coord.xStart - 1
	endX := coord.xEnd + 1
	for _, symbol := range symbols {
		yMatches := symbol.y >= coord.y-1 && symbol.y <= coord.y+1
		xMatches := symbol.x >= startX && symbol.x <= endX
		if yMatches && xMatches {
			return true
		}
	}
	return false
}

func parse(ss string, yCoordinate int) (coordinates []coordinate, symbols []symbol) {
	var bufferedNumber = ""
	var xCoordinate = 0
	for _, char := range ss {
		if (isSymbol(char) || isDot(char)) && len(bufferedNumber) > 0 {
			number, err := strconv.Atoi(bufferedNumber)
			if err != nil {
				panic("Unknown number")
			}
			coordinates = append(
				coordinates,
				coordinate{
					number: number,
					y:      yCoordinate,
					xStart: xCoordinate - len(bufferedNumber),
					xEnd:   xCoordinate - 1,
				})
			bufferedNumber = ""
		}
		if isDigit(char) {
			bufferedNumber += string(char)
		}
		if isSymbol(char) {
			symbols = append(symbols, symbol{symbol: string(char), y: yCoordinate, x: xCoordinate})
		}
		xCoordinate += 1
	}
	if len(bufferedNumber) > 0 {
		number, err := strconv.Atoi(bufferedNumber)
		if err != nil {
			panic("Unknown number")
		}
		coordinates = append(
			coordinates,
			coordinate{
				number: number,
				y:      yCoordinate,
				xStart: xCoordinate - len(bufferedNumber),
				xEnd:   xCoordinate - 1,
			})
	}
	return
}

func main() {
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum1 = 0
	var sum2 = 0
	var allCoordinates = []coordinate{}
	var allSymbols = []symbol{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		coordinates, symbols := parse(line, y)
		y += 1
		allCoordinates = append(allCoordinates, coordinates...)
		allSymbols = append(allSymbols, symbols...)
	}

	// Part 1
	for _, coord := range allCoordinates {
		if isAdjacent(coord, allSymbols) {
			sum1 += coord.number
		}
	}

	// Part 2
	for _, symb := range allSymbols {
		if symb.symbol != "*" {
			continue
		}

		adjacentNumbers := []int{}
		for _, coord := range allCoordinates {
			if isAdjacent(coord, []symbol{symb}) {
				adjacentNumbers = append(adjacentNumbers, coord.number)
			}
		}
		if len(adjacentNumbers) == 2 {
			sum2 += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("First step:", sum1)
	log.Println("Second step:", sum2)
}
