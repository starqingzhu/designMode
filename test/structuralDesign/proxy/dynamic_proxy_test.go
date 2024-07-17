package test

import (
	"designmode/proxy"
	"fmt"
	"testing"
)

func TestDynamicProxy(t *testing.T) {
	got, _ := proxy.Generate("../../proxy/base.go")
	fmt.Println(got)
}
