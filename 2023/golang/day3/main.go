package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type schemaNumbers struct {
	number       int
	indexNumbers [][2]int
}
type schemaGear struct {
	gearIndex   [2]int
	gearNumbers []int
}

type engineSchema struct {
	schema [140][140]string
}

func (es *engineSchema) indexNumbers() []schemaNumbers {
	var numberString = ""
	var indexAray = [][2]int{}
	var schemaNumbersArray = []schemaNumbers{}
	for i := 0; i < 140; i++ {
		// fmt.Println()
		for j := 0; j < 140; j++ {
			_, err := strconv.Atoi(es.schema[i][j])
			if err == nil {
				numberString = numberString + es.schema[i][j]
				siValue := [2]int{i, j}
				indexAray = append(indexAray, siValue)
				// fmt.Println(es.schema[i][j])
			} else {
				// we got a new number!
				if numberString != "" {
					// fmt.Println(numberString)
					stringNum, err := strconv.Atoi(numberString)
					if err == nil {
						snum := schemaNumbers{number: stringNum, indexNumbers: indexAray}
						schemaNumbersArray = append(schemaNumbersArray, snum)
						// fmt.Println(snum)
					}

				}
				numberString = ""
				indexAray = [][2]int{}
			}
		}
	}
	if numberString != "" {
		// fmt.Println(numberString)
		stringNum, err := strconv.Atoi(numberString)
		if err == nil {
			snum := schemaNumbers{number: stringNum, indexNumbers: indexAray}
			schemaNumbersArray = append(schemaNumbersArray, snum)
			// fmt.Println(snum)
		}
	}
	return schemaNumbersArray
}

func (es *engineSchema) indexGears(schemaNums []schemaNumbers) []schemaGear {
	var schemaGearArray = []schemaGear{}
	for i := 0; i < 140; i++ {
		// fmt.Println()
		for j := 0; j < 140; j++ {
			schemaString := es.schema[i][j]
			if schemaString == "*" {
				matchNum := 0
				matchArray := []int{}
				for _, schemaNum := range schemaNums {
					fmt.Println(slices.Contains(schemaNum.indexNumbers, [2]int{i, j}))
					if slices.Contains(schemaNum.indexNumbers, [2]int{i, j}) {
						fmt.Println(schemaString, [2]int{i, j}, schemaNum.indexNumbers)
						matchNum += 1
						matchArray = append(matchArray, schemaNum.number)
					}
				}
				scGear := schemaGear{gearIndex: [2]int{i, j}, gearNumbers: matchArray}
				if matchNum == 2 {
					fmt.Println(schemaString, [2]int{i, j}, matchArray)
					schemaGearArray = append(schemaGearArray, scGear)
				}
			}
		}
	}
	return schemaGearArray
}

func newEngineSchema(inputFile string) *engineSchema {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var newSchema [140][140]string
	scanner := bufio.NewScanner(file)
	rowNum := 0
	for scanner.Scan() {
		row := scanner.Text()
		for colNum, colValue := range row {
			newSchema[rowNum][colNum] = string(colValue)
		}
		rowNum += 1
	}

	return &engineSchema{schema: newSchema}

}

func isValidNumber() {

}

func findValidIndex(inIndex [2]int, es *engineSchema) bool {
	x := inIndex[0]
	y := inIndex[1]

	colMax := len(es.schema[0]) - 1
	rowMax := len(es.schema) - 1

	neighbourIndices := [][2]int{}

	// 1 2 3
	// 4 X 5
	// 6 7 8
	if (x == 0) && (y == 0) {
		// X 5
		// 7 8
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y + 1}) // 8
	} else if (x == 0) && (y == rowMax) {
		// 2 3
		// X 5
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y - 1}) // 3
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
	} else if (x == 0) && (y != 0) {
		// 2 3
		// X 5
		// 7 8
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y - 1}) // 3
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y + 1}) // 8
	} else if (x == colMax) && (y == rowMax) {
		// 1 2
		// 4 X
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y - 1}) // 1
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
	} else if (x == colMax) && (y == 0) {
		// 4 X
		// 6 7
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y + 1}) // 6
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
	} else if (x == colMax) && (y != 0) {
		// 1 2
		// 4 X
		// 6 7
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y - 1}) // 1
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y + 1}) // 6
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
	} else if (x != 0) && (y == rowMax) {
		// 1 2 3
		// 4 X 5
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y - 1}) // 1
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y - 1}) // 3
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
	} else if (x != 0) && (y == 0) {
		// 4 X 5
		// 6 7 8
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y + 1}) // 6
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y + 1}) // 8
	} else {
		// 1 2 3
		// 4 X 5
		// 6 7 8
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y - 1}) // 1
		neighbourIndices = append(neighbourIndices, [2]int{x, y - 1})     // 2
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y - 1}) // 3
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y})     // 4
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y})     // 5
		neighbourIndices = append(neighbourIndices, [2]int{x - 1, y + 1}) // 6
		neighbourIndices = append(neighbourIndices, [2]int{x, y + 1})     // 7
		neighbourIndices = append(neighbourIndices, [2]int{x + 1, y + 1}) // 8
	}
	for _, indexNum := range neighbourIndices {
		// fmt.Println(colMax, rowMax)
		schemaString := es.schema[indexNum[0]][indexNum[1]]
		// neighbourStrings = append(neighbourStrings, schemaString)
		_, err := strconv.Atoi(schemaString)
		if err != nil && schemaString != "." {
			fmt.Println(inIndex, indexNum)
			fmt.Println(schemaString)
			return true
		}
		// fmt.Println(neighbourStrings)
	}
	return false
}

func part1() {
	filePath := "input.txt"
	engineSchematic := newEngineSchema(filePath)
	schemaNums := engineSchematic.indexNumbers()
	totalNum := 0
	for _, schemaNum := range schemaNums {
		for _, indexNum := range schemaNum.indexNumbers {
			validIndex := findValidIndex(indexNum, engineSchematic)
			if validIndex {
				totalNum += schemaNum.number
				break
			}
			// fmt.Println(findNeighbours(indexNum, engineSchematic))
		}
	}
	fmt.Println(totalNum)
	// for i := 0; i < 140; i++ {
	// 	fmt.Println()
	// 	for j := 0; j < 140; j++ {
	// 		fmt.Print(engineSchematic.schema[i][j])
	// 	}
	// }

	// fmt.Println(engineSchematic.schema[139][139])
}
func part2() {
	filePath := "input.txt"
	engineSchematic := newEngineSchema(filePath)
	schemaNums := engineSchematic.indexNumbers()
	totalNum := 0
	gears := engineSchematic.indexGears(schemaNums)
	for _, gear := range gears {
		gearNum1 := gear.gearNumbers[0]
		gearNum2 := gear.gearNumbers[1]
		gearRatio := gearNum1 * gearNum2
		totalNum += gearRatio
	}
	fmt.Println(totalNum)
	// for i := 0; i < 140; i++ {
	// 	fmt.Println()
	// 	for j := 0; j < 140; j++ {
	// 		fmt.Print(engineSchematic.schema[i][j])
	// 	}
	// }

	// fmt.Println(engineSchematic.schema[139][139])
}

func main() {
	// part1()
	part2()
}
