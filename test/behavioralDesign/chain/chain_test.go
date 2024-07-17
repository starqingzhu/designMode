package test

import (
	"designmode/behavioralDesign/chain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	chainObj := &chain.SensitiveWordFilterChain{}

	chainObj.AddFilter(&chain.AdSensitiveWordFilter{})
	assert.Equal(t, false, chainObj.Filter("test"))

	chainObj.AddFilter(&chain.PoliticalWordFilter{})
	assert.Equal(t, true, chainObj.Filter("test"))
}
