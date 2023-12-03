package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numberString string

	for scanner.Scan() {
		numberString = scanner.Text()
		fmt.Printf("%q.\n", numberString)
		numArray := []int{}
		for _, v := range numberString {
			// fmt.Println(string(v))
			if intv, err := strconv.Atoi(string(v)); err == nil {
				numArray = append(numArray, intv)
				// fmt.Printf("%q looks like a number.\n", v)
			}
		}
		fmt.Println(numArray)
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
