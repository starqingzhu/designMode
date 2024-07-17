package test

import (
	"designmode/creativeDesign/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
go test .
*/
func TestHungry(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

/*
go test -bench .
*/
func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var temp = singleton.GetInstance()
			if temp != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
