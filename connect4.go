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

import "sync"

// size of the board
const numCols uint = 7
const numRows uint = 6

// Segment represents an array of moves of length segmentLength
//type Segment [segmentLength]Piece

// size of a winning segment in Connect 4
const segmentLength uint = 4

// C4Board is the main struct that should implement the Board interface
// It maintains the position of a game
// You should not need to add any additional properties to this struct, but
// you may add additional methods
type C4Board struct {
	position [numCols][numRows]Piece // the grid in Connect 4
	colCount [numCols]uint           // how many pieces are in a given column (or how many are "non-empty")
	turn     Piece                   // who's turn it is to play
}

// tracker tracks the segments for board instances
var tracker SegmentTracker = SegmentTracker{make(map[C4Board][]Segment), &sync.Mutex{}}

// Turn returns who's turn it is.
func (board C4Board) Turn() Piece {
	return board.turn
}

// MakeMove puts a piece in column col.
// Returns a copy of the board with the move made.
// Does not check if the column is full (assumes legal move).
func (board C4Board) MakeMove(col Move) Board {
	b := board
	piece := b.Turn()

	// board.colCount[col] will be the empty space in the column
	// technically this can error however it shouldn't be called if
	// it isn't a legal move
	b.position[col][board.colCount[col]] = piece
	b.colCount[col]++

	b.turn = b.Turn().opposite()
	b.LoadSegments()

	return b
}

//LoadSegments loads all the segments into the board
func (board C4Board) LoadSegments() {
	segments := make([]Segment, 1)
	// Loads Vertical Segments
	var i, j uint
	var segment Segment
	for i = 0; i < numCols; i++ {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i][j+1],
				board.position[i][j+2],
				board.position[i][j+3],
			}
			segments = append(segments, segment)
		}
	}

	// Loads Horizontal Segments
	for i = 0; i < numRows; i++ {
		for j = 0; j < numCols-3; j++ {
			segment = Segment{
				board.position[j][i],
				board.position[j+1][i],
				board.position[j+2][i],
				board.position[j+3][i],
			}
			segments = append(segments, segment)
		}
	}

	// Diagonal Segments
	for i = 0; i < numCols-3; i++ {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i+1][j+1],
				board.position[i+2][j+2],
				board.position[i+3][j+3],
			}
			segments = append(segments, segment)
		}
	}

	for i = numCols - 1; i > 2; i-- {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i-1][j+1],
				board.position[i-2][j+2],
				board.position[i-3][j+3],
			}
			segments = append(segments, segment)
		}
	}

	tracker.AddBoard(board, segments)
}

// LegalMoves returns all of the current legal moves.
// Remember, a move is just the column you can play.
func (board C4Board) LegalMoves() []Move {
	// Creates a slice with the capacity of the max amount of possible moves
	legalMoves := make([]Move, 0, 7)

	// Appends a possible move if it isn't full
	for i := uint(0); i < numCols; i++ {
		if board.colCount[i] < numRows {
			legalMoves = append(legalMoves, Move(i))
		}
	}

	return legalMoves
}

// IsWin calculates if the board is in a winning position
// if it is, then returns true, else returns false.
func (board C4Board) IsWin() bool {
	segments := tracker.GetSegments(board)
	if segments == nil {
		board.LoadSegments()
		segments = tracker.GetSegments(board)
	}

	for _, segment := range segments {
		if segment.Equivalent() {
			return true
		}
	}

	return false
}

// IsDraw determines if the board is currently in a draw state
func (board C4Board) IsDraw() bool {

	// If there are no legal moves AND it isn't currently a win, then
	// its a draw. Theoretically, IsDraw is never called before IsWin, therefore
	// we know the board isn't in a winning state and don't neccesarily need that check.
	if legalMoves := board.LegalMoves(); len(legalMoves) == 0 && !board.IsWin() {
		return true
	}

	return false
}

// Evaluate returns the value of the piece's board
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
	segments := tracker.GetSegments(board)
	if segments == nil {
		board.LoadSegments()
		segments = tracker.GetSegments(board)
	}
	var totalScore float32

	for _, segment := range segments {
		totalScore += segment.CalculateScore(player)
	}

	return totalScore

}

// Nice to print board representation
// This will be used in play.go to print out the state of the position
// to the user
func (board C4Board) String() string {
	b := ""

	var j uint
	for i := int(numRows) - 1; i >= 0; i-- {
		b += "|"
		for j = 0; j < numCols; j++ {
			b += board.position[j][i].String() + "|"
		}
		b += "\n"
	}

	return b
}
