package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/// A for Rock, B for Paper, and C for Scissors.
/// X for Rock, Y for Paper, and Z for Scissors.
/// 1 for Rock, 2 for Paper, and 3 for Scissors
/// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
/// 1 defeats 3, 3 defeats 2, 2 defeats 1
/// X lose, Y draw, Z win

var objectValues = map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

func extra_score(p1 int, p2 int) (int, int) {
	diff := p1 - p2
	if diff == 0 {
		return 3, 3
	}
	if diff == 1 {
		return 6, 0
	}
	if diff == -2 {
		return 6, 0
	}
	return 0, 6
}

func rps_1(p1 string, p2 string) (int, int) {
	scoreP1 := objectValues[p1]
	scoreP2 := objectValues[p2]
	extra1, extra2 := extra_score(scoreP1, scoreP2)
	return extra1 + scoreP1, extra2 + scoreP2
}

func rps(p1 string, p2 string) (int, int) {
	scoreP1 := objectValues[p1]
	if p2 == "X" {
		ours := scoreP1 - 1
		if ours == 0 {
			ours = 3
		}
		return scoreP1 + 6, ours
	}
	if p2 == "Y" {
		return scoreP1 + 3, scoreP1 + 3
	}
	if p2 == "Z" {
		ours := scoreP1 + 1
		if ours == 4 {
			ours = 1
		}
		return scoreP1, ours + 6
	}
	return 0, 0
}

func main() {
	inputFile, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	input := string(inputFile)
	lines := strings.Split(input, "\n")
	our_score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		plays := strings.Split(line, " ")
		p1, p2 := rps(plays[0], plays[1])
		fmt.Println(line, p1, p2)
		our_score += p2
	}
	fmt.Println(our_score)
}
