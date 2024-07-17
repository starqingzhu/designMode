package test

import (
	"designmode/creativeDesign/factory"
	"fmt"
	"testing"
)

func TestAbstract(t *testing.T) {
	s := factory.GirlFactoryStore{
		Factory: factory.ChineseGrilFactory{},
	}

	g := s.Factory.CreatGirl(factory.FAT)
	g.Weight()
	fmt.Println("abstractfactory end ....")
}
