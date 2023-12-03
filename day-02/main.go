package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	s := bufio.NewScanner(f)

	sumOfPossibleGames := 0
	cumulativePower := 0
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
		game := ParseGame(line)
		fmt.Printf("%v\n", game)
		if game.IsPossible() {
			sumOfPossibleGames += game.Id
		}
		cumulativePower += game.Power()
	}
	fmt.Println("Sum of possible games:", sumOfPossibleGames)
	fmt.Println("Cumulative power:", cumulativePower)
}

type BagPull map[string]int

type Game struct {
	Id int
	Pulls []BagPull
}

func ParseGame(line string) *Game {
	var gameId int
	fmt.Sscanf(line, "Game %d", &gameId)
	i := strings.Index(line, ": ")
	rawPulls := strings.Split(line[i + 2:], "; ")
	pulls := make([]BagPull, len(rawPulls))
	for i, rawPull := range rawPulls {
		pulls[i] = ParsePull(rawPull)
	}
	return &Game{gameId, pulls}
}

func (g Game) Power() int {
	minNeededRed := 0
	minNeededBlue := 0
	minNeededGreen := 0
	for _, pull := range g.Pulls {
		if pull["red"] > minNeededRed {
			minNeededRed = pull["red"]
		}
		if pull["blue"] > minNeededBlue {
			minNeededBlue = pull["blue"]
		}
		if pull["green"] > minNeededGreen {
			minNeededGreen = pull["green"]
		}
	}
	return minNeededRed * minNeededBlue * minNeededGreen
}

func ParsePull(s string) BagPull {
	pull := make(BagPull, 0)
	for _, colorPull := range strings.Split(s, ", ") {
		var color string
		var numCubes int
		fmt.Println(colorPull)
		fmt.Sscanf(colorPull, "%d %s", &numCubes, &color)
		pull[color] = numCubes
	}
	return pull
}

func (g Game) IsPossible() bool {
	return (
		g.LargestPullFor("red") <= 12 &&
		g.LargestPullFor("green") <= 13 &&
		g.LargestPullFor("blue") <= 14)
}

func (g Game) LargestPullFor(color string) int {
	largestPull := 0
	for _, pull := range g.Pulls {
		numCubes, exists := pull[color]
		if exists && numCubes > largestPull {
			largestPull = numCubes
		}
	}
	return largestPull
}