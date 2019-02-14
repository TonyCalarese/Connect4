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
// minimax.go for CSI 380 Assignment 3
// This file contains a working implementation of Minimax
// You will need to implement the FindBestMove() methods to
// actually evaluate a position by running MiniMax on each of the legal
// moves in a starting position and finding the move associated with the best outcome
package main

import (
	"math"
	"fmt"
)

// Find the best possible outcome evaluation for originalPlayer
// depth is initially the maximum depth
func MiniMax(board Board, maximizing bool, originalPlayer Piece, depth uint) float32 {
	// Base case — terminal position or maximum depth reached
	if board.IsWin() || board.IsDraw() || depth == 0 {
		return board.Evaluate(originalPlayer)
	}

	// Recursive case - maximize your gains or minimize the opponent's gains
	if maximizing {
		var bestEval float32 = -math.MaxFloat32 // arbitrarily low starting point
		for _, move := range board.LegalMoves() {
			result := MiniMax(board.MakeMove(move), false, originalPlayer, depth-1)
			if result > bestEval {
				bestEval = result
			}
		}
		return bestEval
	} else { // minimizing
		var worstEval float32 = math.MaxFloat32
		for _, move := range board.LegalMoves() {
			result := MiniMax(board.MakeMove(move), true, originalPlayer, depth-1)
			if result < worstEval {
				worstEval = result
			}
		}
		return worstEval
	}
}
// Eval represents a move evaluation
type Eval struct {
	m Move
	f float32
}

// ConcurrentFindBestMove finds the best possible move in
// the current position looking up to depth ahead.
// This version looks at each legal move from the starting position
// concurrently (runs minimax on each legal move concurrently)
func ConcurrentFindBestMove(board Board, depth uint) Move {
	var bestMove Move
	var bestScore float32 = -math.MaxFloat32
	legalMoves := board.LegalMoves()

	scores := make(chan Eval, len(legalMoves))
	//evals := make([]eval, len(legalMoves))

	for _, move := range legalMoves {
		go func(move Move) {
			var e Eval
			e.m = move
			e.f = MiniMax(board.MakeMove(move), true, board.Turn(), depth)
			scores <- e
		}(move)
	}

	for i := 0; i < len(legalMoves); i++ {
		eval := <-scores
		//fmt.Printf("m: %d, f: %f\n", eval.m, eval.f)
		if eval.f > bestScore {
			bestScore = eval.f
			bestMove = eval.m
		}
	}
	close(scores)

	return bestMove
}

// FindBestMove finds the best possible move in the current position
// looking up to depth ahead
// This is a non-concurrent version that you may want to test first
func FindBestMove(board Board, depth uint) Move {
	var bestMove Move
	var bestScore float32 = -math.MaxFloat32
	fmt.Println(board)
	for _, move := range board.LegalMoves() {
		// Should we pass board.Turn() or pass the next turn into this function?
		if score := MiniMax(board.MakeMove(move), true, board.Turn(), depth); score > bestScore {
			bestMove = move
			bestScore = score
		}
	}

	return bestMove
}
