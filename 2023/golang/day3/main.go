package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type schemaNumbers struct {
	number       int
	indexNumbers [][2]int
}

type engineSchema struct {
	schema [140][140]string
}

func (es *engineSchema) indexNumbers() []schemaNumbers {
	var numberString = ""
	var indexAray = [][2]int{}
	var schemaNumbersArray = []schemaNumbers{}
	for i := 0; i < 140; i++ {
		fmt.Println()
		for j := 0; j < 140; j++ {
			_, err := strconv.Atoi(es.schema[i][j])
			if err == nil {
				numberString = numberString + es.schema[i][j]
				siValue := [2]int{i, j}
				indexAray = append(indexAray, siValue)
				fmt.Println(es.schema[i][j])
			} else {
				// we got a new number!
				if numberString != "" {
					fmt.Println(numberString)
					stringNum, err := strconv.Atoi(numberString)
					if err == nil {
						snum := schemaNumbers{number: stringNum, indexNumbers: indexAray}
						schemaNumbersArray = append(schemaNumbersArray, snum)
						fmt.Println(snum)
					}

				}
				numberString = ""
				indexAray = [][2]int{}
			}
		}
	}
	if numberString != "" {
		fmt.Println(numberString)
		stringNum, err := strconv.Atoi(numberString)
		if err == nil {
			snum := schemaNumbers{number: stringNum, indexNumbers: indexAray}
			schemaNumbersArray = append(schemaNumbersArray, snum)
			fmt.Println(snum)
		}
	}
	return schemaNumbersArray
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

func part1() {
	filePath := "input.txt"
	engineSchematic := newEngineSchema(filePath)
	engineSchematic.indexNumbers()
	// for i := 0; i < 140; i++ {
	// 	fmt.Println()
	// 	for j := 0; j < 140; j++ {
	// 		fmt.Print(engineSchematic.schema[i][j])
	// 	}
	// }

	// fmt.Println(engineSchematic.schema[139][139])
}

func main() {
	part1()
}
