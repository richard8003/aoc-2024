package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumList []int

type Fubar struct {
	a NumList
	b NumList
}

func main() {
	data := parseData()

	partOne(data)
	partTwo(data)
}

func partTwo(f Fubar) {
	var temp []int
	var sum int

	for _, n := range f.a {
		times := 0
		for _, j := range f.b {
			if j == n {
				times += 1
			}
		}
		temp = append(temp, n*times)
		sum += n * times
	}

	fmt.Println("Part 2:", sum)
}

func partOne(f Fubar) {
	n := calcDiff(f)
	var r int

	for _, m := range n {
		r += m
	}

	fmt.Println("Part 2:", r)
}

func parseData() Fubar {
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("There was an error reading the file")
	}
	test := strings.Trim(string(dat), "\n\n")

	step_one := strings.Split(test, "\n")

	var list_a, list_b NumList

	for _, row := range step_one {
		row_array := strings.Split(row, "  ")

		first_num_string := strings.Trim(row_array[0], " ")
		first_num_int, _ := strconv.Atoi(first_num_string)

		second_num_string := strings.Trim(row_array[1], " ")
		second_num_int, _ := strconv.Atoi(second_num_string)

		list_a = append(list_a, first_num_int)
		list_b = append(list_b, second_num_int)
	}

	sort.Ints(list_a)
	sort.Ints(list_b)

	fubar := Fubar{
		a: list_a,
		b: list_b,
	}

	return fubar

}

func calcDiff(f Fubar) []int {
	var difs []int

	a := f.a
	b := f.b

	for i := range a {

		if a[i] > b[i] {
			difs = append(difs, a[i]-b[i])
		} else {
			difs = append(difs, b[i]-a[i])
		}
	}

	return difs
}
