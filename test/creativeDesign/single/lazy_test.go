package test

import (
	"designmode/creativeDesign/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
go test .
*/
func TestHungryLazy(t *testing.T) {
	assert.Equal(t, singleton.GetLazySingleton(), singleton.GetLazySingleton())
}

/*
go test -bench .
*/
func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var temp = singleton.GetLazySingleton()
			if temp != singleton.GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
