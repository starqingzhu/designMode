package test

import (
	"designmode/bridge"
	"testing"
)

func TestErrorNotifactionNotify(t *testing.T) {
	send := bridge.NewEmailMsgSender([]string{"test@test.com"})
	n := bridge.NewErrorNotification(send)
	n.Notify("test msg")
}
