// out here we will try to solve this with channels and go routines alone

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getRucksacks(channel chan string) {
	inputFile, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	input := string(inputFile)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line != "" {
			channel <- line
		}
	}
	close(channel)
}

func getCommonLetters(rucksacks chan string, channel chan rune) {
	for rucksack := range rucksacks {
		comparmentSize := len(rucksack) / 2
		left := rucksack[0:comparmentSize]
		right := rucksack[comparmentSize:]
		for _, c := range left {
			if strings.ContainsRune(right, c) {
				channel <- c
				break
			}
		}
	}
	close(channel)
}

func makePriorities() map[rune]int {
	priorities := map[rune]int{}
	for r, i := 'a', 1; r <= 'z'; r, i = r+1, i+1 {
		priorities[r] = i
	}
	for r, i := 'A', 27; r <= 'Z'; r, i = r+1, i+1 {
		priorities[r] = i
	}
	return priorities
}

var priorities = makePriorities()

func getPriorities(letters chan rune, channel chan int) {
	for letter := range letters {
		channel <- priorities[letter]
	}
	close(channel)
}

func main() {
	rucksacks := make(chan string)
	go getRucksacks(rucksacks)
	commonLetters := make(chan rune)
	go getCommonLetters(rucksacks, commonLetters)
	priorities := make(chan int)
	go getPriorities(commonLetters, priorities)
	sum := 0
	for priority := range priorities {
		sum += priority
	}
	fmt.Println(sum)
}
