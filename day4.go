package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game_t struct {
	number         int
	winningNumbers map[int]bool
	playerNumbers  []int
}

var digitsRe = regexp.MustCompile(`[0-9]+`)

func parse(ss string) (game game_t) {
	game.playerNumbers = []int{}
	game.winningNumbers = map[int]bool{}

	gameNumbers := strings.Split(ss, ":")
	gameNumber := digitsRe.FindString(gameNumbers[0])
	numbers := strings.Split(gameNumbers[1], "|")

	for _, winningNumber := range digitsRe.FindAllString(numbers[0], -1) {
		winningNum, _ := strconv.Atoi(winningNumber)
		game.winningNumbers[winningNum] = true
	}
	for _, playerNumber := range digitsRe.FindAllString(numbers[1], -1) {
		playerNum, _ := strconv.Atoi(playerNumber)
		game.playerNumbers = append(game.playerNumbers, playerNum)
	}

	game.number, _ = strconv.Atoi(gameNumber)
	return
}

func main() {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum1 = 0
	var sum2 = 0
	copies := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		game := parse(line)
		matchingNumbers := 0
		for _, playerNumber := range game.playerNumbers {
			if game.winningNumbers[playerNumber] {
				matchingNumbers += 1
			}
		}
		if matchingNumbers == 1 {
			sum1 += 1
		} else if matchingNumbers > 1 {
			sum1 += int(math.Pow(2.0, float64(matchingNumbers-1)))
		}
		copies[game.number] += 1 // The original copy
		for copyNum := 0; copyNum < copies[game.number]; copyNum += 1 {
			for gameNum := game.number + 1; gameNum < game.number+matchingNumbers+1; gameNum += 1 {
				copies[gameNum] += 1
			}
		}
	}

	log.Println("First step:", sum1)
	for _, value := range copies {
		sum2 += value
	}
	log.Println("Second step:", sum2)
}
