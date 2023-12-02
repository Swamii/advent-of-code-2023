package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const LIMIT_RED = 12
const LIMIT_GREEN = 13
const LIMIT_BLUE = 14

type set struct {
	red   int
	green int
	blue  int
}
type game_t struct {
	id   int
	sets []set
}

func parse(ss string) game_t {
	gameSets := strings.Split(ss, ":")
	gameIDString := strings.ReplaceAll(gameSets[0], "Game ", "")
	gameID, err := strconv.Atoi(gameIDString)
	if err != nil {
		panic("Invalid game ID")
	}

	setsRaw := strings.Split(gameSets[1], ";")
	var sets = []set{}
	for _, setRaw := range setsRaw {
		var red = 0
		var green = 0
		var blue = 0
		setCubes := strings.Split(setRaw, ", ")
		for _, cubeRaw := range setCubes {
			amountColor := strings.Split(strings.Trim(cubeRaw, " "), " ")
			value, err := strconv.Atoi(amountColor[0])
			if err != nil {
				panic("Failed to parse cube amount")
			}
			switch amountColor[1] {
			case "blue":
				blue = value
			case "green":
				green = value
			case "red":
				red = value
			default:
				panic("Unknown value")
			}
		}
		sets = append(sets, set{red: red, green: green, blue: blue})
	}

	return game_t{id: gameID, sets: sets}
}

func main() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum1 = 0
	var sum2 = 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Reading line:", line)
		game := parse(line)
		log.Println("Parsed game:", game)
		var passes = true
		var highestRed = 0
		var highestGreen = 0
		var highestBlue = 0
		for _, gameSet := range game.sets {
			if gameSet.red > LIMIT_RED || gameSet.green > LIMIT_GREEN || gameSet.blue > LIMIT_BLUE {
				passes = false
			}
			if gameSet.red > highestRed {
				highestRed = gameSet.red
			}
			if gameSet.green > highestGreen {
				highestGreen = gameSet.green
			}
			if gameSet.blue > highestBlue {
				highestBlue = gameSet.blue
			}
		}
		if passes {
			sum1 += game.id
		}
		sum2 += highestRed * highestGreen * highestBlue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("First step:", sum1)
	log.Println("Second step:", sum2)
}
