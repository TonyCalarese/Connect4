package main

// segmentEquivalent checks if all of the pieces in the segment
// are of the same kind and non empty. Returns true if all pieces
// in the segment are of the same kind, false otherwise.
func segmentEquivalent(segment Segment) bool {
	basePiece := segment[0]
	for _, piece := range segment {
		if piece != basePiece || piece == Empty {
			return false
		}
	}

	return true
}

// CheckVertical checks if there is a winning vertical segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckVertical() (segments []Segment, win bool) {
	win = false
	var segment Segment

	var i, j uint
	for i = 0; i < numCols; i++ {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i][j+1],
				board.position[i][j+2],
				board.position[i][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
				return
			}
		}
	}

	return
}

// VerticalWin makes checking for winning conditions
// on the board look cleaner for vertical checking
func (board C4Board) VerticalWin() bool {
	_, verticalWin := board.CheckVertical()

	return verticalWin
}

// CheckHorizontal checks if there is a winning vertical segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckHorizontal() (segments []Segment, win bool) {
	win = false
	var segment Segment

	var i, j uint

	for i = 0; i < numRows; i++ {
		for j = 0; j < numCols-3; j++ {
			segment = Segment{
				board.position[j][i],
				board.position[j+1][i],
				board.position[j+2][i],
				board.position[j+3][i],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
				return
			}
		}
	}

	return
}

// HorizontalWin makes checking for winning conditions
// on the board look cleaner for horizontal checking
func (board C4Board) HorizontalWin() bool {
	_, horizontalWin := board.CheckHorizontal()

	return horizontalWin
}

// CheckDiagonal checks if there is a winning diagonal segment
// it will return immediately if a win is found. If a win is not found
// it will return all the segments tested and a win status of "false".
func (board C4Board) CheckDiagonal() (segments []Segment, win bool) {
	win = false
	var segment Segment

	var i, j uint
	// Left to right diagonal checking
	for i = 0; i < numCols-3; i++ {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i+1][j+1],
				board.position[i+2][j+2],
				board.position[i+3][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
				return
			}
		}
	}

	// Right to left diagonal checking
	for i = numCols - 1; i > 2; i-- {
		for j = 0; j < numRows-3; j++ {
			segment = Segment{
				board.position[i][j],
				board.position[i-1][j+1],
				board.position[i-2][j+2],
				board.position[i-3][j+3],
			}
			segments = append(segments, segment)

			if segmentEquivalent(segment) {
				win = true
				return
			}
		}
	}

	return
}

// DiagonalWin makes checking for winning conditions
// on the board look cleaner for diagonal checking
func (board C4Board) DiagonalWin() bool {
	_, diagonalWin := board.CheckDiagonal()

	return diagonalWin
}

// CalculateDirection calculates the score in the direction of segments
func CalculateDirection(segments []Segment, player Piece) (score float32) {

	for _, segment := range segments {
		score += calculateScore(segment, player)
	}

	return
}

func calculateScore(segment Segment, player Piece) float32 {
	pieceCount := 0
	pieceToCount := Empty
	for _, piece := range segment {
		if piece != Empty && pieceToCount == Empty {
			pieceToCount = piece
			pieceCount++
		} else if piece != pieceToCount && piece != Empty {
			return 0.0
		} else if piece != Empty {
			pieceCount++
		}
	}

	score := func() float32 {
		if pieceCount == 0 {
			return 0.0
		} else if pieceCount == 1 {
			return 1.0
		} else if pieceCount == 2 {
			return 5.0
		} else if pieceCount == 3 {
			return 50.0
		} else {
			return 5000.0
		}
	}

	if pieceToCount != player && pieceToCount != Empty {
		return -score()
	}

	return score()
}
