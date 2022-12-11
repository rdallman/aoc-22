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
// X = lose
// Y = draw
// Z = win

func points(opp, me string) int {
	var points int
	switch {
	case opp == "A": // rock
		switch {
		case me == "X": // lose
			points += 0
			points += 3
		case me == "Y": // draw
			points += 3
			points += 1
		case me == "Z": // win
			points += 6
			points += 2
		}
	case opp == "B": // paper
		switch {
		case me == "X": // lose
			points += 0
			points += 1
		case me == "Y": // draw
			points += 3
			points += 2
		case me == "Z": // win
			points += 6
			points += 3
		}
	case opp == "C": // scissors
		switch {
		case me == "X": // lose
			points += 0
			points += 2
		case me == "Y": // draw
			points += 3
			points += 3
		case me == "Z": // win
			points += 6
			points += 1
		}
	}

	return points
}

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
