package test

import (
	"designmode/proxy"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProxy(t *testing.T) {
	proxy := proxy.NewUserProxy(&proxy.User{})
	err := proxy.Login("test", "password")

	require.Nil(t, err)
}
