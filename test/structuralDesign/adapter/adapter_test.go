package test

import (
	"designmode/adapter"
	"testing"
)

func TestAliyunClientAdapterCreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a adapter.ICreateServer = &adapter.AliyunClientAdapter{
		Client: adapter.AliyunClient{},
	}

	a.CreateServer(1.0, 2.0)
}

func TestAwsClientAdapterCreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口
	var a adapter.ICreateServer = &adapter.AwsClientAdapter{
		Client: adapter.AWSClient{},
	}

	a.CreateServer(1.0, 2.0)
}
