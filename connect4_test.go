// Connect 4 for CSI 380 by David Kopec
// This file tests the final implementation
// with test positions as well as a couple of the
// methods of C4Board.

// Note that in board literals, columns go from left to right,
// not up and down.

package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetSegments1(t *testing.T) {
	// Not really a test, i wanted a playground tbh
	tracker := SegmentTracker{make(map[C4Board][]Segment), &sync.Mutex{}}

	b1 := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 6, 2, 5, 0, 0},
		turn:     Red}

	s1 := []Segment{{1, 1, 1, 1}}

	b2 := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 0, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{3, 3, 1, 1, 0, 0, 0},
		turn:     Black}

	s2 := []Segment{{0, 2, 0, 0}}

	tracker.AddBoard(b1, s1)
	tracker.AddBoard(b2, s2)

	func(board C4Board) {
		s5 := tracker.GetSegments(b1)
		s5[0][2] = 2
		fmt.Printf("s5 = %v\n", s5)
	}(b1)
	s3 := tracker.GetSegments(b1)
	s4 := tracker.GetSegments(b2)

	fmt.Printf("s1 = %v, s2 = %v, s3 = %v, s4 = %v", s1, s2, s3, s4)
}

// Check if two slices of moves are equivalent
func checkEquivalent(p1s, p2s []Move) bool {
	if len(p1s) != len(p2s) { // same length
		return false
	}

	for _, p := range p1s { // same items
		if !contains(p2s, p) {
			return false
		}
	}

	return true
}

// Test if the right legal moves are generated
// blank board
func TestLegalMoves1(t *testing.T) {
	b := C4Board{turn: Black}
	expected := []Move{0, 1, 2, 3, 4, 5, 6}
	actual := b.LegalMoves()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test if the right legal moves are generated
// board with 1 filled column
func TestLegalMoves2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 6, 2, 5, 0, 0},
		turn:     Red}
	expected := []Move{0, 1, 3, 4, 5, 6}
	actual := b.LegalMoves()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

/*
func TestCheckHorizontal1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 1, 2, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	_, actual := b.CheckHorizontal()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCheckVertical1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 2, 2, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	_, actual := b.CheckVertical()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCheckDiagonal1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 1, 2, 2, 1},
		[6]Piece{2, 1, 2, 1, 1, 2},
		[6]Piece{1, 1, 1, 2, 1, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 2, 2, 1, 1},
		[6]Piece{1, 2, 2, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	_, actual := b.CheckDiagonal()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}
*/
// Test if a non-full board is called a draw
func TestDraw1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{1, 2, 1, 2, 1, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 6, 2, 5, 0, 0},
		turn:     Red}
	expected := false
	actual := b.IsDraw()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test if a full board is called a draw
func TestDraw2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsDraw()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// false
func TestWin1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := false
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 1, 1}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin3(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 1, 1, 1, 2, 2},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 1, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin4(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 1, 2, 1},
		[6]Piece{2, 1, 2, 1, 2, 2},
		[6]Piece{1, 2, 1, 2, 1, 2},
		[6]Piece{1, 2, 1, 2, 1, 1},
		[6]Piece{2, 2, 1, 2, 1, 2},
		[6]Piece{1, 1, 2, 1, 2, 1},
		[6]Piece{2, 2, 1, 1, 1, 2}},
		colCount: [7]uint{6, 6, 6, 6, 6, 6, 6},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test win position
// true
func TestWin5(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 2, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{4, 3, 1, 0, 0, 0, 0},
		turn:     Black}
	expected := true
	actual := b.IsWin()
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestMove1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 0, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{3, 3, 1, 1, 0, 0, 0},
		turn:     Black}

	expected := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 2, 0, 0, 0},
		[6]Piece{1, 1, 1, 1, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{3, 4, 1, 1, 0, 0, 0},
		turn:     Red}

	actual := b.MakeMove(1)
	if expected != actual {
		t.Errorf("Test failed: expected\n%v to be\n%v", actual, expected)
	}
}

func TestFindBestMove1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 3, 1, 1, 0, 0, 0},
		turn:     Red}
	expected := Move(1)
	actual := FindBestMove(b, 4)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestFindBestMove2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{2, 2, 1, 0, 0, 0},
		[6]Piece{2, 1, 2, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 2, 2, 3, 3, 0, 0},
		turn:     Black}
	expected := Move(4)
	actual := FindBestMove(b, 3)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestFindBestMove3(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 1, 2, 0, 0, 0},
		turn:     Red}
	actual := FindBestMove(b, 5)
	if (actual != Move(1)) && (actual != Move(4)) { // two save moves
		t.Errorf("Test failed: expected %v to be 1 or 4", actual)
	}
}

// Test Concurrent Find Best Move
// stop the win
func TestConcurrentFindBestMove1(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 1, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{2, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 3, 1, 1, 0, 0, 0},
		turn:     Red}
	expected := Move(1)
	actual := ConcurrentFindBestMove(b, 4)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test Concurrent Find Best Move
// take the win
func TestConcurrentFindBestMove2(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{2, 2, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{1, 1, 0, 0, 0, 0},
		[6]Piece{2, 2, 1, 0, 0, 0},
		[6]Piece{2, 1, 2, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{2, 2, 2, 3, 3, 0, 0},
		turn:     Black}
	expected := Move(4)
	actual := ConcurrentFindBestMove(b, 3)
	if expected != actual {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

// Test Concurrent Find Best Move
// block a future win
func TestConcurrentFindBestMove3(t *testing.T) {
	b := C4Board{position: [7][6]Piece{
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{1, 0, 0, 0, 0, 0},
		[6]Piece{1, 2, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0},
		[6]Piece{0, 0, 0, 0, 0, 0}},
		colCount: [7]uint{0, 0, 1, 2, 0, 0, 0},
		turn:     Red}
	actual := ConcurrentFindBestMove(b, 5)
	if (actual != Move(1)) && (actual != Move(4)) { // two save moves
		t.Errorf("Test failed: expected %v to be 1 or 4", actual)
	}
}

func TestCalculateScore1(t *testing.T) {
	var s Segment = [4]Piece{Red, Black, Red, Red}
	var expected float32 = 0.0
	var actual float32 = s.CalculateScore(Red)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore2(t *testing.T) {
	var s Segment = [4]Piece{Red, Empty, Red, Red}
	var expected float32 = 50.0
	var actual float32 = s.CalculateScore(Red)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore3(t *testing.T) {
	var s Segment = [4]Piece{Red, Empty, Red, Red}
	var expected float32 = -50.0
	var actual float32 = s.CalculateScore(Black)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore4(t *testing.T) {
	var s Segment = [4]Piece{Empty, Empty, Red, Red}
	var expected float32 = -5.0
	var actual float32 = s.CalculateScore(Black)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore5(t *testing.T) {
	var s Segment = [4]Piece{Red, Empty, Empty, Red}
	var expected float32 = -5.0
	var actual float32 = s.CalculateScore(Black)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore6(t *testing.T) {
	var s Segment = [4]Piece{Black, Red, Red, Red}
	var expected float32 = 0.0
	var actual float32 = s.CalculateScore(Black)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}

func TestCalculateScore7(t *testing.T) {
	var s Segment = [4]Piece{Red, Red, Red, Red}
	var expected float32 = 5000.0
	var actual float32 = s.CalculateScore(Red)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", actual, expected)
	}
}
