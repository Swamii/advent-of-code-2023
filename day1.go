package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var charsToDigits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func charsToDigit(ss string) (val string, startIndex int) {
	for charDigit := range charsToDigits {
		if idx := strings.Index(ss, charDigit); idx != -1 {
			return charsToDigits[charDigit], idx
		}
	}
	return "", -1
}

func filterDigits(ss string) (ret []string) {

	var buf = ""
	for _, s := range ss {
		char := string(s)
		buf += char

		var digit = ""
		if _, err := strconv.Atoi(char); err == nil {
			digit = char
			buf = ""
		} else if converted, index := charsToDigit(buf); index != -1 {
			digit = converted
			buf = buf[index+1:]
		}

		if _, err := strconv.Atoi(digit); err == nil {
			ret = append(ret, digit)
		}
	}
	return ret
}

func main() {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Reading line:", line)
		digits := filterDigits(line)
		if len(digits) == 0 {
			continue
		}
		log.Println("Found filtered digits:", strings.Join(digits, ", "))
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]
		number, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			panic("shouldn't happen")
		}
		sum += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println("Final result:", sum)
}
