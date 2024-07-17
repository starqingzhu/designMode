package test

import (
	"designmode/creativeDesign/factory"
	"fmt"
	"testing"
)

func TestBuildFactoryProduct(t *testing.T) {
	type args struct {
		t string
	}

	tests := []struct {
		name string
		args args
	}{
		{name: "p1", args: args{t: "p1"}},
		{name: "p2", args: args{t: "p2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := factory.BuildFactoryProduct(tt.args.t)
			got.CreateProduct().Print()

		})
	}
	fmt.Println("BuildFactoryProduct end ....")
}
