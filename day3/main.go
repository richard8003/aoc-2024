package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Data []string

func main() {
	input := parse_input()

	// partOne(input)
	partTwo(input)
}

func partTwo(input string) {
	data := extractTwo(input)
	data.Trim()
	data.PartOne()
	fmt.Println("\n")

}

func partOne(input string) {
	data := extractOne(input)
	data.Trim()
	data.PartOne()
	fmt.Println("\n")

}

// ====================== Methods

func (d Data) Print() {
	for _, n := range d {
		fmt.Println(n)
	}
}

func (d Data) Trim() {
	for i := range d {
		fu := strings.Trim(d[i], "mul(")
		bar := strings.Trim(fu, ")")
		d[i] = bar
	}
}

func (d Data) PartOne() {
	sum := 0
	for _, s := range d {
		arr := strings.Split(s, ",")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])

		sum += (a * b)

	}
	fmt.Println("Part 1:", sum)
}

// ====================== Functions

func extractTwo(input string) Data {
	var data Data

	pattern := regexp.MustCompile(`(?m)mul\(\d+,\d+\)|do\(|don't`)

	fu := pattern.FindAllString(input, -1)

	flag := true

	for _, s := range fu {

		if s == "do(" {
			flag = true
			continue
		}
		if s == "don't" {
			flag = false
			continue
		}

		if flag {
			data = append(data, s)
		}

	}

	data.Print()

	return data
}

func extractOne(input string) Data {
	var data Data

	pattern := regexp.MustCompile(`(?m)mul\(\d+,\d+\)`)

	fu := pattern.FindAllString(input, -1)

	for i := range fu {

		data = append(data, fu[i])
	}

	// return pattern.FindAllString(input, -1)
	return data
}

func parse_input() string {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("There was an error reading the file")
	}

	return (string(dat))
}
