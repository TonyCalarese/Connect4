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
	var col Move

	fmt.Println("Enter a Column you would like to insert in(1-7): ")
	for {
		if _, err := fmt.Scanln(&col); err == nil && (col <= 7 || col >= 1) {
			break
		}
	}
	col--
	return col
}

func getDepth() uint {
	var depth uint

	fmt.Println("Enter a difficulty (1-5): ")
	for {
		if _, err := fmt.Scanln(&depth);
		 err == nil && (depth <= 5 || depth >= 1) {
			break
		}
	}
	return depth
}
// Main game loop
func main() {
	//var c chan Move = make(chan Move)


	//Set difficulty
	var depth uint = getDepth()

	
	for !gameBoard.IsDraw() && !gameBoard.IsWin() {
		fmt.Printf("%s", gameBoard.String())

		legal := false
		col := getPlayerMove()
		for legal == false {
			legalMoves := gameBoard.LegalMoves()
			for _, move := range legalMoves {
				if col == move {
					legal = true
					break
				}
			}
			if legal == false {
				col = getPlayerMove()
			}
		}
		gameBoard = gameBoard.MakeMove(col)
		//gameBoard = gameBoard.MakeMove(FindBestMove(gameBoard, 3))

		 if gameBoard.IsWin() {
			 //Player has won 
			 //Need to check the win after every Move
			 fmt.Println("!!!!!!!!!Congradulations You Won!!!!!!!!!!!!!")
			 break
		 }
		gameBoard = gameBoard.MakeMove(ConcurrentFindBestMove(gameBoard, depth))

		if gameBoard.IsWin() {
			//Player has won 
			//Need to check the win after every Move
			fmt.Println("!!!!!!!!!The Computer Won!!!!!!!!!!!!")
			break
		}
		/*
		fmt.Printf("Black: %f", gameBoard.Evaluate(Black))
		fmt.Printf("Red: %f", gameBoard.Evaluate(Red))
		fmt.Println()
		*/
	}
	fmt.Printf("%s", gameBoard.String())
}
