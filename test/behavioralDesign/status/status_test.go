package test

import (
	"designmode/behavioralDesign/status"
	"testing"
)

func TestStatus(t *testing.T) {
	m := &status.Machine{State: status.GetLeaderApproveState()}
	m.Approval()

	m.Reject()
	m.Approval()
	m.Approval()
}
