package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numberString string
	var total int
	for scanner.Scan() {
		numberString = scanner.Text()
		// fmt.Printf("%q.\n", numberString)
		numArray := []int{}
		for _, v := range numberString {
			// fmt.Println(string(v))
			if intv, err := strconv.Atoi(string(v)); err == nil {
				numArray = append(numArray, intv)
				// fmt.Printf("%q looks like a number.\n", v)
			}
		}
		total = total + numArray[0]*10 + numArray[len(numArray)-1]
		// fmt.Println(numArray[0])
		// fmt.Println(numArray[len(numArray)-1])
		// fmt.Println(numArray[0]*10 + numArray[len(numArray)-1])
		// fmt.Println(total)

		// fmt.Println(scanner.Text())
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stringMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	var numberString string
	var total int
	for scanner.Scan() {
		numberString = scanner.Text()
		// fmt.Printf("%q.\n", numberString)
		for letter_string, num_string := range stringMap {
			numberString = strings.ReplaceAll(numberString, letter_string, num_string)
		}
		// fmt.Printf("%q.\n", numberString)
		numArray := []int{}
		for _, v := range numberString {
			// fmt.Println(string(v))
			if intv, err := strconv.Atoi(string(v)); err == nil {
				numArray = append(numArray, intv)
				// fmt.Printf("%q looks like a number.\n", v)
			}
		}
		total = total + numArray[0]*10 + numArray[len(numArray)-1]
		// fmt.Println(numArray)
		// // fmt.Println(numArray[len(numArray)-1])
		// fmt.Println(numArray[0]*10 + numArray[len(numArray)-1])
		// fmt.Println(total)

		// fmt.Println(scanner.Text())
	}
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var inputFile string = "input.txt"
	part1(inputFile)
	part2(inputFile)
}
