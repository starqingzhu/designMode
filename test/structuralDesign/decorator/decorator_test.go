package test

import (
	"designmode/decorator"
	"fmt"
	"testing"
)

func TestColorSquareDraw(t *testing.T) {
	sq := decorator.Square{}

	csq := decorator.NewColorSquare(sq, "red")
	got := csq.Draw()
	fmt.Printf("%s\n", got)
}
