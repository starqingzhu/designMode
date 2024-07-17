package factory

import "fmt"

const (
	FAT  = "fat"
	THIN = "thin"
)

type IGirl interface {
	Weight()
}

type (
	FatGirl  struct{}
	ThinGirl struct{}

	AmericanFatGirl  struct{}
	AmericanThinGirl struct{}
)

func (g FatGirl) Weight() {
	fmt.Println("chinese girl weight: 80kg")
}

func (g ThinGirl) Weight() {
	fmt.Println("chinese girl weight: 50kg")
}

func (g AmericanFatGirl) Weight() {
	fmt.Println("chinese girl weight: 80kg")
}

func (g AmericanThinGirl) Weight() {
	fmt.Println("chinese girl weight: 50kg")
}

type IAFactory interface {
	CreatGirl(like string) IGirl
}

type (
	ChineseGrilFactory  struct{}
	AmericanGirlFactory struct{}
)

func (f ChineseGrilFactory) CreatGirl(like string) IGirl {
	if like == FAT {
		return &FatGirl{}
	} else if like == THIN {
		return &ThinGirl{}
	}
	return nil
}

func (f AmericanGirlFactory) CreatGirl(like string) IGirl {
	if like == FAT {
		return &AmericanFatGirl{}
	} else if like == THIN {
		return &AmericanThinGirl{}
	}
	return nil
}

type GirlFactoryStore struct {
	Factory IAFactory
}

func (s GirlFactoryStore) CreateGirl(like string) IGirl {
	return s.Factory.CreatGirl(like)
}
