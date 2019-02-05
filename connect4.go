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
// connect4.go for CSI 380 Assignment 3
// The struct C4Board should implement the Board
// interface specified in Board.go
// Note: you will almost certainly need to add additional
// utility functions/methods to this file.

package main

// size of the board
const numCols uint = 7
const numRows uint = 6

// size of a winning segment in Connect 4
const segmentLength uint = 4

// The main struct that should implement the Board interface
// It maintains the position of a game
// You should not need to add any additional properties to this struct, but
// you may add additional methods
type C4Board struct {
	position [numCols][numRows]Piece // the grid in Connect 4
	colCount [numCols]uint           // how many pieces are in a given column (or how many are "non-empty")
	turn     Piece                   // who's turn it is to play
}

// Who's turn is it?
func (board C4Board) Turn() Piece {
	return board.turn
}

// Put a piece in column col.
// Returns a copy of the board with the move made.
// Does not check if the column is full (assumes legal move).
func (board C4Board) MakeMove(col Move) Board {
	// YOUR CODE HERE
	return nil
}

// All of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {

	legalMoves := make([]Move, 0, 7)

	var i uint
	for i = 0; i < numCols; i++ {
		if board.colCount[i] < numRows {
			legalMoves = append(legalMoves, Move(i))
		}
	}

	return legalMoves
}

// Is it a win?
func (board C4Board) IsWin() bool {
	// YOUR CODE HERE
	var i, j uint
	for i = 0; i < numCols; i++ {
		for j = 0; j < numRows; j++ {
			// if Check Horizontal || Check Vertical || Check Diagonal
			// 		return true
		}
	}
	return false

	// We theoretically don't even need to have two iterations there as isWin() and
	// isDraw() doesn't check what segment wins, just if there is a win, so we don't need
	// to go from a specific point.
}

// Is it a draw?
func (board C4Board) IsDraw() bool {
	// YOUR CODE HERE
	// if LegalMoves() is empty
	//		return true
	//
	var i, j uint
	for i = 0; i < numCols; i++ {
		for j = 0; j < numRows; j++ {

		}
	}
	return true
}

// Who is winning in this position?
// This function scores the position for player
// and returns a numerical score
// When player is doing well, the score should be higher
// When player is doing worse, player's returned score should be lower
// Scores mean nothing except in relation to one another; so you can
// use any scale that makes sense to you
// The more accurately Evaluate() scores a position, the better that minimax will work
// There may be more than one way to evaluate a position but an obvious route
// is to count how many 1 filled, 2 filled, and 3 filled segments of the board
// that the player has (that don't include any of the opponents pieces) and give
// a higher score for 3 filleds than 2 filleds, 1 filleds, etc.
// You may also need to score wins (4 filleds) as very high scores and losses (4 filleds
// for the opponent) as very low scores
func (board C4Board) Evaluate(player Piece) float32 {
	// YOUR CODE HERE
	return 0.0
}

// Nice to print board representation
// This will be used in play.go to print out the state of the position
// to the user
func (board C4Board) String() string {
	// YOUR CODE HERE

	return ""
}
