package chess

import (
	"fmt"
	"strconv"
	"sync"
)

// Board ...
type Board struct {
	size  int
	board [][]int
	count int
}

// SetQueen ...
func (br *Board) SetQueen(a int) {
	for i := 0; i < br.size; i++ {
		if br.tryQueen(a, i) {
			br.board[a][i] = 1
			thr := new(Thread)
			thr.board = br
			thr.numColumn = a + 1
			thr.Start()
			thr.Join()
			br.board[a][i] = 0
		}
	}
}

func (br *Board) tryQueen(a, b int) bool {
	for i := 0; i < a; i++ {
		if br.board[i][b] == 1 {
			return false
		}
	}

	for i := 1; i <= a && b-i >= 0; i++ {
		if br.board[a-i][b-i] == 1 {
			return false
		}
	}

	for i := 1; i <= a && b+i < br.size; i++ {
		if br.board[a-i][b+i] == 1 {
			return false
		}
	}

	return true
}

func (br *Board) showBoard() {
	for a := 0; a < br.size; a++ {
		for b := 0; b < br.size; b++ {
			if br.board[a][b] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Thread ...
type Thread struct {
	mux      sync.Mutex
	board    *Board
	numColumn int
}

// Start ...
func (thread *Thread) Start() {
	thread.mux.Lock()
	go thread.run()
}

// Join ...
func (thread *Thread) Join() {
	thread.mux.Lock()
	thread.mux.Unlock()
}

func (r *Thread) run() {
	if r.numColumn == r.board.size {
		r.board.count++
		fmt.Print("Result №" + strconv.Itoa(r.board.count) + "\n")
		r.board.showBoard()
	}

	for i := 0; i < r.board.size; i++ {
		if r.board.tryQueen(r.numColumn, i) {
			r.board.board[r.numColumn][i] = 1
			thr := new(Thread)
			thr.board = r.board
			thr.numColumn = r.numColumn + 1
			thr.Start()
			thr.Join()
			r.board.board[r.numColumn][i] = 0
		}
	}
	r.mux.Unlock()
}

// Init
func Init(size int) *Board {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	return &Board{
		size:  size,
		count: 0,
		board: matrix,
	}
}
