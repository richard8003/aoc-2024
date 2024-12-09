package main

import (
	"fmt"
	"os"
	"strings"
)

var Up string = "^"
var Right string = ">"
var Down string = "v"
var Left string = "<"

type Node struct {
	x   int
	y   int
	val string
}

type NodeArray []Node

type Karta struct {
	Nodes    []NodeArray
	GuardPos Node
}

/*


....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...


*/

func main() {
	karta := makeMap()

	for i := 0; i <= 5347; i++ {
		guard := karta.guardPosition()
		x := guard.x
		y := guard.y

		if x > 129 || y > 129 {
			// if x == len(karta.Nodes)-1 || y > len(karta.Nodes[0])-1 {
			break
		}

		karta.Move()
	}

	karta.Calc()
	// karta.Print()
	karta.guardPosition()

}

func (k Karta) Calc() {
	counter := 1

	for i := range k.Nodes {
		for _, j := range k.Nodes[i] {
			if j.val == "X" {
				counter += 1
			}
		}

	}

	fmt.Println(counter)
}

func (k Karta) Move() {

	guard := k.guardPosition()

	x := guard.x
	y := guard.y

	// CAN go up
	if guard.val == "^" && k.Nodes[x-1][y].val != "#" {
		k.Nodes[x-1][y].val = "^"
		k.Nodes[x][y].val = "X"
	}

	// CAN'T go up
	if guard.val == "^" && k.Nodes[x-1][y].val == "#" {
		k.Nodes[x][y].val = ">"
	}

	// CAN go right
	if guard.val == ">" && k.Nodes[x][y+1].val != "#" {
		k.Nodes[x][y+1].val = ">"
		k.Nodes[x][y].val = "X"
	}

	// CAN'T go right
	if guard.val == ">" && k.Nodes[x][y+1].val == "#" {
		k.Nodes[x][y].val = "v"
		// k.PrintRaw()
	}

	// CAN go down
	if guard.val == "v" && k.Nodes[x+1][y].val != "#" {
		k.Nodes[x+1][y].val = "v"
		k.Nodes[x][y].val = "X"
	}

	// CAN'T go down
	if guard.val == "v" && k.Nodes[x+1][y].val == "#" {
		k.Nodes[x][y].val = "<"
	}

	// CAN go left
	if guard.val == "<" && k.Nodes[x][y-1].val != "#" {
		k.Nodes[x][y-1].val = "<"
		k.Nodes[x][y].val = "X"
	}

	// CAN'T go left
	if guard.val == "<" && k.Nodes[x][y-1].val == "#" {
		k.Nodes[x][y].val = "^"
	}

	// anders
	// k.Print()

}

func (k Karta) Print() {
	for i := range k.Nodes {
		for _, j := range k.Nodes[i] {
			fmt.Print(j.val)
		}
		fmt.Println("")
	}
}

func (k Karta) PrintRaw() {
	for i := range k.Nodes {
		fmt.Println(k.Nodes[i])
	}
}

func (k Karta) guardPosition() Node {
	for i := range k.Nodes {
		for _, j := range k.Nodes[i] {
			if j.val == "^" || j.val == ">" || j.val == "v" || j.val == "<" {
				// fmt.Printf(" \nThe guard is at X:%d Y:%d, facing %s\n\n", j.x, j.y, j.val)
				k.GuardPos.y = j.y
				k.GuardPos.x = j.x
				k.GuardPos.val = j.val
				return j
			}
		}
	}
	return Node{}
}

func makeMap() Karta {
	rawMap, err := os.ReadFile("./input.txt")
	// rawMap, err := os.ReadFile("./map.txt")
	handleError(err)

	var array []NodeArray

	stringLine := strings.Split(strings.TrimRight(string(rawMap), "\n\n"), "\n")

	for i := 0; i < len(stringLine); i++ {

		var line NodeArray
		var newNode Node
		for j := 0; j < len(stringLine); j++ {
			newNode.x = i
			newNode.y = j
			newNode.val = string(stringLine[i][j])
			line = append(line, newNode)
		}
		array = append(array, line)
	}

	var result Karta
	result.Nodes = array

	return result
}
