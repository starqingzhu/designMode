package factory

import (
	"fmt"
)

type IProduct interface {
	Desc() string
	Print()
}

type SProduct struct {
	Name string
}

func (p SProduct) Desc() string {
	return p.Name
}

func (p SProduct) Print() {
	fmt.Printf("%s ...\n", p.Name)
}

type SProduct1 struct {
	SProduct
}

type SProduct2 struct {
	SProduct
}

func BuildProduct(name string) IProduct {
	switch name {
	case "p1":
		return SProduct1{SProduct{Name: "p1"}}
	case "p2":
		return SProduct2{SProduct{Name: "p2"}}
	}
	return nil
}
