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
	myAnts           int
	oppAnts          int
	_type            int
	initialResources int
	neigh0           int
	neigh1           int
	neigh2           int
	neigh3           int
	neigh4           int
	neigh5           int
}

type Base struct {
	myBaseIndex  int64
	oppBaseIndex int64
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

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
		cell._type = _type
		cell.initialResources = initialResources
		cell.neigh0 = neigh0
		cell.neigh1 = neigh1
		cell.neigh2 = neigh2
		cell.neigh3 = neigh3
		cell.neigh4 = neigh4
		cell.neigh5 = neigh5
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
			myMap.cells[i].initialResources = resources
			myMap.cells[i].myAnts = myAnts
			myMap.cells[i].oppAnts = oppAnts
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
			fmt.Fprintf(os.Stderr, "Cell: %d, type: %d, resources: %d, neigh0: %d, neigh1: %d, neigh2: %d, neigh3: %d, neigh4: %d, neigh5: %d\n", cell, cell._type, cell.initialResources, cell.neigh0, cell.neigh1, cell.neigh2, cell.neigh3, cell.neigh4, cell.neigh5)
		}

		// WAIT | LINE <sourceIdx> <targetIdx> <strength> | BEACON <cellIdx> <strength> | MESSAGE <text>
		// fmt.Println("WAIT")
		// var actions []string
		// for _, neighbour := range myMap.bases[0]
	}
}
