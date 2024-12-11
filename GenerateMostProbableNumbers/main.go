package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/creack/pty"
	"golang.org/x/term"
)

var (
	mainNumbers = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
		"13",
		"14",
		"15",
		"16",
		"17",
		"18",
		"19",
		"20",
		"21",
		"22",
		"23",
		"24",
		"25",
		"26",
		"27",
		"28",
		"29",
		"30",
		"31",
		"32",
		"33",
		"34",
		"35",
		"36",
		"37",
		"38",
		"39",
		"40",
		"41",
		"42",
		"43",
		"44",
		"45",
		"46",
		"47",
		"48",
		"49",
	}
	luckyNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	chunkSize    = 150000
)

func main() {
	// Setting up window sizing
	cmd := exec.Command("sh", "-c", "")
	master, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("Error opening PTY:", err)
	}
	defer master.Close()

	// Get terminal size
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
	}

	// Set terminal size
	size := pty.Winsize{Cols: uint16(width), Rows: uint16(height)}
	err = pty.Setsize(master, &size)
	if err != nil {
		fmt.Println("Error setting terminal size:", err)
	}

	// Timing
	startTime := time.Now()

	// Set up CSV writer
	fileCounter := 1
	csvFile := fmt.Sprintf("Lottoallcomb%d.csv", fileCounter)
	file, err := os.Create(csvFile)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Main loop
	// Pretty ugly huh
	totalLoopCounter := 1
	for _, a := range mainNumbers[:11] {
		aInt, _ := strconv.Atoi(a)
		for _, b := range mainNumbers[4:25] {
			bInt, _ := strconv.Atoi(b)
			for _, c := range mainNumbers[9:39] {
				cInt, _ := strconv.Atoi(c)
				for _, d := range mainNumbers[20:45] {
					dInt, _ := strconv.Atoi(d)
					for _, e := range mainNumbers[29:] {
						eInt, _ := strconv.Atoi(e)
						for _, f := range luckyNumbers {
							if a == b || a == c || a == d || a == e || b == c || b == d ||
								b == e ||
								c == d ||
								c == e ||
								d == e {
								break
							} else if (bInt == aInt+1 && cInt == aInt+2 && dInt == aInt+3) || (cInt == bInt+1 && cInt == dInt-1 && cInt == eInt-2) {
								break
							} else if ((aInt >= 1 && aInt <= 9) && (bInt >= 1 && bInt <= 9) && (cInt >= 1 && cInt <= 9) && (dInt >= 1 && dInt <= 9) && (eInt >= 1 && eInt <= 9)) ||
								((aInt >= 11 && aInt <= 19) && (bInt >= 11 && bInt <= 19) && (cInt >= 11 && cInt <= 19) && (dInt >= 11 && dInt <= 19) && (eInt >= 11 && eInt <= 19)) ||
								((aInt >= 21 && aInt <= 29) && (bInt >= 21 && bInt <= 29) && (cInt >= 21 && cInt <= 29) && (dInt >= 21 && dInt <= 29) && (eInt >= 21 && eInt <= 29)) ||
								((aInt >= 31 && aInt <= 39) && (bInt >= 31 && bInt <= 39) && (cInt >= 31 && cInt <= 39) && (dInt >= 31 && dInt <= 39) && (eInt >= 31 && eInt <= 39)) ||
								((aInt >= 41 && aInt <= 49) && (bInt >= 41 && bInt <= 49) && (cInt >= 41 && cInt <= 49) && (dInt >= 41 && dInt <= 49) && (eInt >= 41 && eInt <= 49)) {
								break
							} else if aInt%2 == 0 && bInt%2 == 0 && cInt%2 == 0 && dInt%2 == 0 && eInt%2 == 0 {
								break
							} else if aInt%2 != 0 && bInt%2 != 0 && cInt%2 != 0 && dInt%2 != 0 && eInt%2 != 0 {
								break
							} else if (aInt%3 == 0 && bInt%3 == 0 && cInt%3 == 0 && dInt%3 == 0 && eInt%3 == 0) || (aInt%4 == 0 && bInt%4 == 0 && cInt%4 == 0 && dInt%4 == 0 && eInt%4 == 0) ||
								(aInt%5 == 0 && bInt%5 == 0 && cInt%5 == 0 && dInt%5 == 0 && eInt%5 == 0) || (aInt%6 == 0 && bInt%6 == 0 && cInt%6 == 0 && dInt%6 == 0 && eInt%6 == 0) ||
								(aInt%7 == 0 && bInt%7 == 0 && cInt%7 == 0 && dInt%7 == 0 && eInt%7 == 0) || (aInt%8 == 0 && bInt%8 == 0 && cInt%8 == 0 && dInt%8 == 0 && eInt%8 == 0) ||
								(aInt%9 == 0 && bInt%9 == 0 && cInt%9 == 0 && dInt%9 == 0 && eInt%9 == 0) {
								break
							} else if a[1:] == b[1:] && a[1:] == c[1:] && a[1:] == d[1:] && a[1:] == e[1:] {
							} else {
								result := []string{a, b, c, d, e, f}
								err := writer.Write(result)
								if err != nil {
									fmt.Println("Error writing to CSV:", err)
									return
								}
								totalLoopCounter++
								if totalLoopCounter%chunkSize == 0 {
									fileCounter++
									csvFile = fmt.Sprintf("Lottoallcomb%d.csv", fileCounter)
									file, err = os.Create(csvFile)
									if err != nil {
										fmt.Println("Error creating CSV file:", err)
										return
									}
									defer file.Close()
									writer = csv.NewWriter(file)
									defer writer.Flush()
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\nTotal running time:", time.Since(startTime))
	fmt.Println("Total lines saved:", totalLoopCounter)
}

func terminalSize() (int, int, error) {
	info, err := os.Stdout.Stat()
	if err != nil {
		return 0, 0, err
	}
	return int(info.Size()), 0, nil
}
