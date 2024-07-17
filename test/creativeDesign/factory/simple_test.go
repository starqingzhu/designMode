package test

import (
	"designmode/creativeDesign/factory"
	"fmt"
	"testing"
)

func TestBuildProduct(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
		want factory.IProduct
	}{
		{name: "p1", args: args{t: "p1"}, want: factory.SProduct1{}},
		{name: "p2", args: args{t: "p2"}, want: factory.SProduct2{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := factory.BuildProduct(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("BuildProduct() = %v, want %v", got, tt.want)
			// } else {
			// 	got.Print()
			// }
			got := factory.BuildProduct(tt.args.t)
			got.Print()

		})
	}
	fmt.Println("BuildProduct end ....")
}
