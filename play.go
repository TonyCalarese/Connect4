/*
	Author: Tony Calarese, Adam Decosta
	Class: CSI-380
	Assignment: Assignment 3
	Due Date: February 14, 2019 11:59 PM
	Certification of Authenticity:
	We certify that this is entirely our own work, except where we have given
	fully-documented references to the work of others. We understand the definition
	and consequences of plagiarism and acknowledge that the assessor of this
	assignment may, for the purpose of assessing this assignment:
	- Reproduce this assignment and provide a copy to another member of academic
	- staff; and/or Communicate a copy of this assignment to a plagiarism checking
	- service (which may then retain a copy of this assignment on its database for
	- the purpose of future plagiarism checking)
*/
// main.go for CSI 380 Assignment 3
// This file includes the main game loop
// that actually creates a human vs computer game.

package main

import (
	"fmt"
)

var gameBoard Board = C4Board{turn: Black}

// Find the user's next move
//Source of Reference: https://golang.org/pkg/fmt/#Scanf
//https://golang.org/src/fmt/scan.go?s=2653:2699#L53
func getPlayerMove() Move {
	// YOUR CODE HERE
	var col Move

	fmt.Println("Enter a Column you would like to insert in: ")
	fmt.Scan(&col)

	return col
}

// Main game loop
func main() {
	// YOUR CODE HERE
	//Only looping this a few times to just get a few inserted for testing purposes
	for i := 0; i < 2; i++ {
		fmt.Printf("%s", gameBoard.String())
		col := getPlayerMove()
		gameBoard = gameBoard.MakeMove(col)
	}
	fmt.Printf("%s", gameBoard.String())
}
