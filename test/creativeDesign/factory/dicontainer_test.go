package test

import (
	"designmode/creativeDesign/factory"
	"fmt"
	"testing"
)

// A 依赖关系 A -> B -> C
type A struct {
	B *B
}

// NewA NewA
func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

// B B
type B struct {
	C *C
}

// NewB NewB
func NewB(c *C) *B {
	return &B{C: c}
}

// C C
type C struct {
	Num int
}

// NewC NewC
func NewC() *C {
	return &C{
		Num: 1,
	}
}

func TestDicontainer(t *testing.T) {
	container := factory.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d\n", a, a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Dicontainer end ....")
}
