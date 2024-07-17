package test

import (
	"designmode/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	data, sensitive := getData()
	strategyType := "file"

	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, err := strategy.NewStorageStrategy(strategyType)
	assert.NoError(t, err)
	assert.NoError(t, storage.Save("./test.txt", data))
}

func getData() ([]byte, bool) {
	return []byte("test data"), true
}
