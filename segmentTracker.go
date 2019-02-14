package main

import "sync"

// SegmentTracker stores and keeps track of segments for
// various instances of the board type
type SegmentTracker struct {
	boards map[C4Board][]Segment
	mutex  *sync.Mutex
}

// AddBoard tracks a new board
func (tracker SegmentTracker) AddBoard(board C4Board, segments []Segment) {
	tracker.mutex.Lock()
	tracker.boards[board] = segments
	tracker.mutex.Unlock()
}

// RemoveBoard removes a board from the tracker
func (tracker SegmentTracker) RemoveBoard(board C4Board) {
	tracker.mutex.Lock()
	delete(tracker.boards, board)
	tracker.mutex.Unlock()
}

// GetSegments returns the segments for the specified board instance
func (tracker SegmentTracker) GetSegments(board C4Board) []Segment {
	tracker.mutex.Lock()
	segments := tracker.boards[board]
	tracker.mutex.Unlock()

	return segments
}
