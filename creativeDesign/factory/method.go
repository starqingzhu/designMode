package factory

type IFactory interface {
	CreateProduct() IProduct
}

type SFactory1 struct{}

func (f SFactory1) CreateProduct() IProduct {
	return SProduct1{SProduct{Name: "p1"}}
}

type SFactory2 struct{}

func (f SFactory2) CreateProduct() IProduct {
	return SProduct2{SProduct{Name: "p2"}}
}

func BuildFactoryProduct(name string) IFactory {
	switch name {
	case "p1":
		return SFactory1{}
	case "p2":
		return SFactory2{}
	}
	return nil
}
