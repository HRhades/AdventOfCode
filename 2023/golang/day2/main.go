package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cubeSet struct {
	redCubes   int
	blueCubes  int
	greenCubes int
}

func newSet(redcubes int, bluecubes int, greencubes int) *cubeSet {
	set := cubeSet{redCubes: redcubes, blueCubes: bluecubes, greenCubes: greencubes}
	return &set
}

func (s *cubeSet) isValid(cs *cubeSet) bool {
	// fmt.Printf("%d - %d red cubes   %t\n", cs.redCubes, s.redCubes, cs.redCubes < s.redCubes)
	// fmt.Printf("%d - %d blue cubes  %t\n", cs.blueCubes, s.blueCubes, cs.blueCubes < s.blueCubes)
	// fmt.Printf("%d - %d green cubes %t\n", cs.greenCubes, s.greenCubes, cs.greenCubes < s.greenCubes)
	if cs.redCubes < s.redCubes {
		// fmt.Println("returning set as NOT Valid")
		return false
	}
	if cs.blueCubes < s.blueCubes {
		// fmt.Println("returning set as NOT Valid")
		return false
	}
	if cs.greenCubes < s.greenCubes {
		// fmt.Println("returning set as NOT Valid")
		return false
	}
	// fmt.Println("returning set as Valid")
	return true
}

type cubeGame struct {
	gameId   int
	cubeSets []*cubeSet
}

func newGame(id int) *cubeGame {
	cg := cubeGame{gameId: id}
	return &cg
}

func (g *cubeGame) addSet(cs *cubeSet) []*cubeSet {
	g.cubeSets = append(g.cubeSets, cs)
	return g.cubeSets
}

func (g *cubeGame) isValid(cs *cubeSet) bool {
	for _, s := range g.cubeSets {
		if !s.isValid(cs) {
			return false
		}
	}
	return true
}

func parseGameString(gstring string) *cubeGame {
	// fmt.Println(gstring)
	gameSplit := strings.Split(gstring, ":")
	// fmt.Println(gameSplit)
	gameIdString := gameSplit[0]
	IdInt, err := strconv.Atoi(strings.Split(gameIdString, " ")[1])
	if err != nil {
		log.Fatal("Cannot convert game ID")
	}
	cg := newGame(IdInt)
	// fmt.Printf("Game ID: %d\n", IdInt)

	cubeSetsString := gameSplit[1]
	cubeGameStrings := strings.Split(cubeSetsString, ";")
	// fmt.Println(cubeGameStrings)

	for _, cubeSetString := range cubeGameStrings {

		cubeColors := strings.Split(cubeSetString, ",")
		// fmt.Println(cubeColors)
		numRed := 0
		numBlue := 0
		numGreen := 0
		for _, colors := range cubeColors {
			colorsSplit := strings.Split(strings.TrimSpace(colors), " ")
			// fmt.Println(colorsSplit)
			colorName := strings.TrimSpace(colorsSplit[1])
			colorCountString := strings.TrimSpace(colorsSplit[0])
			// fmt.Printf("%q", colorCountString)
			colorCountNum, err := strconv.Atoi(colorCountString)
			if err != nil {
				log.Fatal("Cannot convert cube number")
			}
			if colorName == "red" {
				numRed = colorCountNum
			}
			if colorName == "blue" {
				numBlue = colorCountNum
			}
			if colorName == "green" {
				numGreen = colorCountNum
			}
		}
		cs := newSet(numRed, numBlue, numGreen)
		cg.addSet(cs)
	}
	return cg
}

func part1(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	checkSet := newSet(12, 14, 13)
	var total int
	for scanner.Scan() {
		gameString := scanner.Text()

		cgame := parseGameString(gameString)
		if cgame.isValid(checkSet) {

			fmt.Printf("%v\n", cgame)
			total += cgame.gameId
		}
		// fmt.Println(cgame.isValid(&cubeSet{redCubes: 12, blueCubes: 14, greenCubes: 13}))
	}
	fmt.Println(total)
}

func main() {
	filePath := "input.txt"
	part1(filePath)
}
