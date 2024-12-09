package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lists []List

type List struct {
	Row       []int
	Direction string
	Status    string
}

func main() {
	lists := parseData()
	lists.Set_direction()
	lists.Set_status()
	// lists.Print()

	fmt.Print("Part 1: ")
	lists.Count_safe_lists()

	for i, n := range lists {
		if n.Status == "unsafe" {

			if n.ListChecks() {

				n.Status = "safe"
				lists[i] = n
			}

		}
	}
	// lists.Print()
	fmt.Print("Part 2: ")
	lists.Count_safe_lists()

}

func (l List) ListChecks() bool {

	for i := range l.Row {
		var temp List
		for j, m := range l.Row {
			if i == j {
				continue
			}
			temp.Row = append(temp.Row, m)

		}
		if temp.is_it_safe() {
			l.Status = "safe"
			return true
		}

	}

	return false
}

func (l Lists) Count_safe_lists() {
	sum := 0

	for i := range l {
		if l[i].Status == "safe" {
			sum += 1
		}
	}

	fmt.Println("\nNumber of safe lists:", sum)
}

func (l List) Make_modified_list(indx int) []int {
	var mod_list []int

	row := l.Row

	for i, n := range row {
		if i == indx {
			continue
		}
		mod_list = append(mod_list, n)

	}

	return mod_list
}

func (l List) is_it_safe() bool {

	increasing := false
	for i, n := range l.Row {
		if i == 0 {
			if n < l.Row[i+1] {
				increasing = true
			}
		}

		// if I'm on the last number there is no need to continue...
		if i == len(l.Row)-1 {
			break
		}

		current := n
		next := l.Row[i+1]

		if increasing {
			if current > next || current == next || next-current > 3 {
				return false
			}
		}

		if !increasing {
			if current < next || current == next || current-next > 3 {
				return false
			}
		}
	}

	return true
}

func additional_checks(l List) bool {

	fmt.Println("Original list:", l.Row)
	for i := range l.Row {
		modified := l.Make_modified_list(i)
		fmt.Println("Modified list: ", modified)

		for x := range modified {

			if x == len(modified)-1 {
				break
			}

			current := modified[x]
			next := modified[x+1]

			if l.Direction == "increasing" {
				if current > next || current == next || next-current > 3 {
					continue
				}
			}

			if l.Direction == "decreasing" {
				if current < next || current == next || current-next > 3 {
					continue
				}
			}
		}

	}

	fmt.Println("==================")

	return false
}

func (l List) BACKUP_is_it_safe() bool {

	for i, n := range l.Row {

		// if I'm on the last number there is no need to continue...
		if i == len(l.Row)-1 {
			break
		}

		current := n
		next := l.Row[i+1]

		if l.Direction == "increasing" {
			if current > next || current == next || next-current > 3 {
				return false
			}
		}

		if l.Direction == "decreasing" {
			if current < next || current == next || current-next > 3 {
				return false
			}
		}
	}

	return true
}

func (l Lists) Set_direction() {
	for indx := range l {

		for i := range l[indx].Row {
			// make sure I'm not on the last number
			if i == len(l[indx].Row)-1 {
				break
			}

			current := l[indx].Row[i]
			next := l[indx].Row[i+1]

			if current == next {
				continue
			}

			if current > next {
				l[indx].Direction = "decreasing"

			}
			if current < next {
				l[indx].Direction = "increasing"
			}
		}
	}
}

func (l Lists) Set_status() {
	for indx := range l {
		if l[indx].is_it_safe() {
			l[indx].Status = "safe"

		} else {
			l[indx].Status = "unsafe"
		}
	}

}

func (l Lists) Print() {
	for _, list := range l {
		fmt.Println(list)
	}

}

func parseData() Lists {
	dat, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("There was an error reading the file")
	}
	test := strings.Trim(string(dat), "\n\n")

	step_one := strings.Split(test, "\n")

	var lists Lists

	for _, n := range step_one {
		modifiedLists := parseData_makeRows(n)
		lists = append(lists, modifiedLists)
	}

	return lists
}

func parseData_makeRows(s string) List {
	var result List
	string_array := strings.Split(s, " ")

	// make a slice of numbers and add it to result
	for _, char := range string_array {
		n, err := strconv.Atoi(char)
		if err != nil {
			fmt.Println("Error converting character to int")
		}
		result.Row = append(result.Row, n)
	}
	return result
}
