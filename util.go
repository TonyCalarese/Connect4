package main

// CheckHorizontal checks if there is a winning horizontal segment
func (board C4Board) checkHorizontal() bool {	
	if board.position[i][j] && board.position[i+1][j] && board.position[i+2][j] == board.position[i+3][j] {
	return true
	}
	else {
	return false
	}
}

// CheckVertical checks if there is a winning vertical segment
func (board C4Board) checkVerticle bool {
	if board.position[i][j] && board.position[i][j+1] && board.position[i][j+2] == board.position[i][j+3] {
	return true
	}
	else {
	return false
	}
}

// CheckDiagonal checks if there is a winning diagonal segment
func (board C4Board) checkDiagnal() bool {
		
	//Check left to right
	//Check right to left
	return
}
