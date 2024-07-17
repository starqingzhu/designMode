package test

import (
	"designmode/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmsSend(t *testing.T) {
	tel := template.NewTelecomSms()
	err := tel.Send("test", 1239999)
	assert.NoError(t, err)
}
