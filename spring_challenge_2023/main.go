package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	cells []Cell
	bases []Base
}

type Cell struct {
	index            int64
	myAnts           int64
	oppAnts          int64
	_type            int64
	initialResources int64
	neighbours       [6]int64
}

type Base struct {
	myBaseIndex  int64
	oppBaseIndex int64
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
var actions []string // actions to do
var myBase Cell      // my base

func main() {
	var myMap Map
	myMap.cells = make([]Cell, 0)
	myMap.bases = make([]Base, 0)
	var cristalCells []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	// numberOfCells: amount of hexagonal cells in this myMap
	var numberOfCells int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &numberOfCells)

	for i := 0; i < numberOfCells; i++ {
		// _type: 0 for empty, 1 for eggs, 2 for crystal
		// initialResources: the initial amount of eggs/crystals on this cell
		// neigh0: the index of the neighbouring cell for each direction
		var _type, initialResources, neigh0, neigh1, neigh2, neigh3, neigh4, neigh5 int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &_type, &initialResources, &neigh0, &neigh1, &neigh2, &neigh3, &neigh4, &neigh5)
		var cell Cell
		cell.index = int64(i)
		cell._type = int64(_type)
		cell.initialResources = int64(initialResources)
		cell.neighbours[0] = int64(neigh0)
		cell.neighbours[1] = int64(neigh1)
		cell.neighbours[2] = int64(neigh2)
		cell.neighbours[3] = int64(neigh3)
		cell.neighbours[4] = int64(neigh4)
		cell.neighbours[5] = int64(neigh5)
		myMap.cells = append(myMap.cells, cell)
		if cell.initialResources > 0 {
			cristalCells = append(cristalCells, strconv.Itoa(i))
		}
	}
	var numberOfBases int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &numberOfBases)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < numberOfBases; i++ {
		myBaseIndex, _ := strconv.ParseInt(inputs[i], 10, 32)
		myMap.bases = append(myMap.bases, Base{myBaseIndex: myBaseIndex})
	}
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < numberOfBases; i++ {
		oppBaseIndex, _ := strconv.ParseInt(inputs[i], 10, 32)
		myMap.bases = append(myMap.bases, Base{oppBaseIndex: oppBaseIndex})
	}
	for {
		for i := 0; i < numberOfCells; i++ {
			// resources: the current amount of eggs/crystals on this cell
			// myAnts: the amount of your ants on this cell
			// oppAnts: the amount of opponent ants on this cell
			var resources, myAnts, oppAnts int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &resources, &myAnts, &oppAnts)
			myMap.cells[i].initialResources = int64(resources)
			myMap.cells[i].myAnts = int64(myAnts)
			myMap.cells[i].oppAnts = int64(oppAnts)
		}
		/*
								       ----------- RULES -----------
			            - BEACON `index` `strength`: place une balise de puissance `strength` sur la cellule `index`.
			            - LINE `index1` `index2` `strength`: place des balises le long d'un chemin entre la cellule index1 et la cellule index2. Toutes les balises placées sont de puissance strength. Le chemin le plus court est choisi automatiquement.
			            - WAIT: ne rien faire.
			            - MESSAGE text. Affiche un texte sur votre côté du HUD.

			                ----------- PRIORITY -----------
			            1.Les actions LINE sont calculées.
			            2.Les actions BEACON sont calculées.
			            3.Les fourmis se déplacent.
			            4.Les cristaux sont récoltés et les points sont marqués.
		*/

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Fprintf(os.Stderr, "Number of cells: %d\n", numberOfCells)
		fmt.Fprintf(os.Stderr, "My base index: %d\n", myMap.bases[0].myBaseIndex)
		fmt.Fprintf(os.Stderr, "Opp base index: %d\n", myMap.bases[0].oppBaseIndex)
		fmt.Fprintf(os.Stderr, "Cells with cristals: %s\n", cristalCells)
		for _, cell := range myMap.cells {
			fmt.Fprintf(os.Stderr, "Cell: %d, type: %d, resources: %d, neigh0: %d, neigh1: %d, neigh2: %d, neigh3: %d, neigh4: %d, neigh5: %d\n", cell, cell._type, cell.initialResources, cell.neighbours[0], cell.neighbours[1], cell.neighbours[2], cell.neighbours[3], cell.neighbours[4], cell.neighbours[5])
		}

		// WAIT | LINE <sourceIdx> <targetIdx> <strength> | BEACON <cellIdx> <strength> | MESSAGE <text>
		// fmt.Println("WAIT")

		myBase = getBase(myMap, myMap.bases[0].myBaseIndex)
		actions = make([]string, 0)
		actions = append(actions, searchPath(myMap, myBase))
	}
}

func getBase(myMap Map, BaseIndex int64) Cell {
	for _, cell := range myMap.cells {
		if cell.index == BaseIndex {
			return cell
		}
	}
	return Cell{}
}

// func to search for the shortest path between cell and my base and return action
func searchPath(myMap Map, cell Cell) string {
	for _, neighbour := range cell.neighbours {
		if getBase(myMap, neighbour).initialResources > 0 {
			return "BEACON " + strconv.Itoa(int(neighbour)) + " 1"
		}
	}
	return "WAIT"
}
