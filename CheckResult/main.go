package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"time"
)

// Initial bet
// TODO: ask the user
var bet = []int{1, 2, 3, 4, 5}

func main() {
	// os.Open() opens file in read-only mode
	// and return a pointer of type os.File
	file, err := os.Open("loto_201911.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// Closes the file when over
	defer file.Close()

	// csv.NewReader take a os.File as parameter
	// and return a new csv.Reader
	reader := csv.NewReader(file)

	// Specify delimiter
	reader.Comma = ';'

	// Read the first line, corresponding to the headers
	if _, err := reader.Read(); err != nil {
		panic(err)
	}

	// ReadAll reads all the draws from the CSV file
	// and returns them as slice of slices of string
	draws, err := reader.ReadAll()
	if err != nil {
		fmt.Sprintln("Error reading draws", err)
	}

	// Will hold numbers based on the last time their were draw
	statistics := make(map[int]time.Time)

	for _, draw := range draws {
		firstNumber, _ := strconv.Atoi(draw[1])
		secondNumber, _ := strconv.Atoi(draw[2])
		thirdNumber, _ := strconv.Atoi(draw[3])
		fourthNumber, _ := strconv.Atoi(draw[4])
		fifthNumber, _ := strconv.Atoi(draw[5])
		// luckyNumber, _ := strconv.Atoi(draw[6])
		// bet := append(bet, luckyNumber)
		firstDraw := []int{firstNumber, secondNumber, thirdNumber, fourthNumber, fifthNumber}

		secondDrawFirstNumber, _ := strconv.Atoi(draw[7])
		secondDrawSecondNumber, _ := strconv.Atoi(draw[8])
		secondDrawThirdNumber, _ := strconv.Atoi(draw[9])
		secondDrawFourthNumber, _ := strconv.Atoi(draw[10])
		secondDrawFifthNumber, _ := strconv.Atoi(draw[11])
		secondDraw := []int{
			secondDrawFirstNumber,
			secondDrawSecondNumber,
			secondDrawThirdNumber,
			secondDrawFourthNumber,
			secondDrawFifthNumber,
		}

		drawDate, err := time.Parse("02/01/2006", draw[0])
		if err != nil {
			panic(err)
		}

		statistics[firstNumber] = drawDate
		statistics[secondNumber] = drawDate
		statistics[thirdNumber] = drawDate
		statistics[fourthNumber] = drawDate
		statistics[fifthNumber] = drawDate

		// To compare
		sort.Ints(firstDraw)
		sort.Ints(secondDraw)
		sort.Ints(bet)

		// Did we win ???
		if slices.Equal(firstDraw, bet) {
			fmt.Println("They are equal yeah")
			fmt.Println(firstDraw, bet)
		}
		if slices.Equal(secondDraw, bet) {
			fmt.Println("They are equal on the second draw youhou")
			fmt.Println(secondDraw, bet)
		}

		bet = CalculateNextBet(statistics)
		fmt.Println(bet)
	}
}

// Return a slice of int
// The ints are the 5 numbers who were draw the oldest
func CalculateNextBet(statistics map[int]time.Time) []int {
	// Create a slice of int containing the keys
	keys := make([]int, 0, len(statistics))
	for key := range statistics {
		keys = append(keys, key)
	}
	// Sort this slice by statistics time value
	sort.Slice(
		keys,
		func(i, j int) bool { return statistics[keys[i]].Before(statistics[keys[j]]) },
	)

	nextBet := make([]int, 5)
	copy(nextBet, keys[:5])

	return nextBet
}
