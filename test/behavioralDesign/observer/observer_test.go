package test

import (
	"designmode/observer"
	"testing"
)

func TestSubjectNotify(t *testing.T) {
	sub := &observer.Subject{}
	sub.Register(&observer.Observer1{})
	sub.Register(observer.Observer2{})

	sub.Notify("hi")
}
