package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type schemaIndex struct {
	index [][]int
}

type schemaNumbers struct {
	number       int
	indexNumbers []schemaIndex
}

type engineSchema struct {
	schema [140][140]string
}

func (es *engineSchema) indexNumbers() {
	for i := 0; i < 140; i++ {
		fmt.Println()
		var numberArray = []int{}
		for j := 0; j < 140; j++ {
			IdInt, err := strconv.Atoi(es.schema[i][j])
			if err == nil {
				numberArray = append(numberArray, IdInt)
			} else {

			}
			fmt.Print(es.schema[i][j])
			fmt.Print(numberArray)
		}
	}
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
	// for i := 0; i < 140; i++ {
	// 	fmt.Println()
	// 	for j := 0; j < 140; j++ {
	// 		fmt.Print(engineSchematic.schema[i][j])
	// 	}
	// }

	fmt.Println(engineSchematic.schema[139][139])
}

func main() {
	part1()
}
