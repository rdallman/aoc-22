package main

import (
	"bufio"
	"fmt"
	"os"
)

// points
// 1 = rock
// 2 = paper
// 3 = scissors
//
// 0 = loss
// 3 = draw
// 6 = win
//
// A = rock
// B = paper
// C = scissors
//
// X = rock
// Y = paper
// Z = scissors

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanWords)
	var score int
	for fs.Scan() {
		opp := fs.Text()
		fs.Scan()
		me := fs.Text()
		score += points(opp, me)
	}

	readFile.Close()
	fmt.Println(score)
}

func points(opp, me string) int {
	var score int
	switch {
	case me == "X": // rock
		score += 1
		switch {
		case opp == "A": // rock
			score += 3
		case opp == "B": // paper
			score += 0
		case opp == "C": // scissors
			score += 6
		}
	case me == "Y": // paper
		score += 2
		switch {
		case opp == "A": // rock
			score += 6
		case opp == "B": // paper
			score += 3
		case opp == "C": // scissors
			score += 0
		}
	case me == "Z": // scissors
		score += 3
		switch {
		case opp == "A": // rock
			score += 0
		case opp == "B": // paper
			score += 6
		case opp == "C": // scissors
			score += 3
		}
	}
	return score
}
