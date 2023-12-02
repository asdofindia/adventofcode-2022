package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")
	inputFile, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	input := string(inputFile)
	lines := strings.Split(input, "\n")
	calories := []int{}
	current := 0
	for _, line := range lines {
		if line == "" {
			calories = append(calories, current)
			current = 0
			continue
		}
		calorie, _ := strconv.Atoi(line)
		current += calorie
	}
	slices.Sort(calories)
	fmt.Println(calories[len(calories)-1])
	lastThree := calories[len(calories)-3 : len(calories)]
	fmt.Println(lastThree)
	fmt.Println(lastThree[0] + lastThree[1] + lastThree[2])
}
