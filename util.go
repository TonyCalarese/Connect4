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
