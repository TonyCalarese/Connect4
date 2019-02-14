package main

// Segment represents a length of segmentLength pieces
type Segment [segmentLength]Piece

// CalculateScore calculates the score of the given segment
func (segment Segment) CalculateScore(player Piece) (score float32) {
	pieceCount := 0
	scoredPiece := Empty

	for _, currentPiece := range segment {
		if currentPiece != Empty && scoredPiece == Empty {
			scoredPiece = currentPiece
			pieceCount++
		} else if currentPiece != scoredPiece && currentPiece != Empty {
			return 0.0
		} else if currentPiece != Empty {
			pieceCount++
		}
	}

	if pieceCount == 0 {
		score = 0.0
	} else if pieceCount == 1 {
		score = 1.0
	} else if pieceCount == 2 {
		score = 5.0
	} else if pieceCount == 3 {
		score = 50.0
	} else {
		score = 5000.0
	}

	if scoredPiece != player {
		score = -score
	}

	return
}

func (segment Segment) Equivalent() bool {
	basePiece := segment[0]
	for _, piece := range segment {
		if piece != basePiece || piece == Empty {
			return false
		}
	}

	return true
}
