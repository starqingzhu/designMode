package test

import (
	"designmode/flyweight"
	"fmt"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board1 := flyweight.NewChessBoard()
	board1.Move(1, 1, 2)
	fmt.Println(board1.String())

	board2 := flyweight.NewChessBoard()
	board2.Move(2, 2, 3)
	fmt.Println(board2.String())

}
