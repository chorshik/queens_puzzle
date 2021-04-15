package main

import (
	"runtime"

	"github.com/ebladrocher/queens_puzzle/chess"
)

func main() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
	board := chess.Init(12)
	board.SetQueen(0)
}
