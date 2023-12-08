package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var input string

var stringToNumberMap = map[string]uint8{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var NUMBER_MAP map[string]int = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

//--------------------------------------------------------------//

// --------------------------- HELPERS ---------------------------//
// finds ascii digits fron byte values
func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}

func updateDigits(firstDigit *int, lastDigit *int, val int) {
	if *firstDigit == -1 {
		*firstDigit = val
	}
	*lastDigit = val
}

func stringToNumber(s string) int8 {
	if dig, err := strconv.Atoi(s); err == nil {
		return int8(dig)
	}

	dig, ok := stringToNumberMap[s]
	// If the key exists
	if ok {
		return int8(dig)
	}

	return -1
}

func RoundTwoProblem() int {
	// input = os.Getenv("INPUT")

	// inputLines := strings.Split(input, "\n")

	// var sum uint = 0

	// for _, line := range inputLines {
	// 	var dig1, dig2 int8 = -1, -1

	// 	for i := 0; i <= len(line)-1 && dig1 == -1; i++ {
	// 		for j := i; dig1 == -1 && j <= len(line)-1 && j-i <= 5; j++ {
	// 			window := line[i:j]
	// 			dig1 = stringToNumber(window)
	// 		}
	// 	}

	// 	for k := len(line) - 1; k >= 0 && dig2 == -1; k-- {
	// 		for m := k; dig2 == -1 && m >= 0 && k-m <= 5; m-- {
	// 			window := line[m : k+1]
	// 			dig2 = stringToNumber(window)
	// 		}
	// 	}

	// 	digsCombined := fmt.Sprintf("%s%s", strconv.Itoa(int(dig1)), strconv.Itoa(int(dig2)))

	// 	lineValues, err := strconv.Atoi(digsCombined)

	// 	if err != nil {
	// 		panic("Unable to parse the line")
	// 	}

	// 	fmt.Println("Line: ", line, "val: ", lineValues)

	// 	sum += uint(lineValues)
	// }

	// fmt.Printf("Result: %v \n", sum)

	// return sum
	input = os.Getenv("INPUT")

	// my solution is the same but it's not working for some edge case so imma copy paste
	inputLines := strings.Split(input, "\n")
	sum := 0
	maxWinLen := 5 // window can't exceed this length if number is in words
	for _, line := range inputLines {
		fmt.Println(line)
		firstDigit, lastDigit := -1, -1
		lineLen := len(line)
		for winSt, winEnd := 0, 0; winSt < lineLen; winSt++ {
			winEnd = winSt

			// if starting index contains a number, just use that as is and continue
			if isDigit(line[winSt]) {
				num, _ := strconv.Atoi(string(line[winSt]))
				updateDigits(&firstDigit, &lastDigit, num)
				continue
			}

			// else use sliding window to find the word if possible
			for winEnd < lineLen && winEnd-winSt < maxWinLen {
				if val, ok := NUMBER_MAP[string(line[winSt:winEnd+1])]; ok {
					updateDigits(&firstDigit, &lastDigit, val)
					break
				}
				winEnd++
			}
		}
		sum += firstDigit*10 + lastDigit
	}

	return sum
}

func main() {
	fmt.Println("Day 1 lesgooooo Go i missed you 🥺")

	godotenv.Load()

	result := RoundTwoProblem()

	fmt.Println(result)
}
