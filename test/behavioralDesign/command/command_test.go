/*
 * @Author: bsun
 * @Date: 2024-07-10 14:37:58
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-10 14:39:45
 */

package test

import (
	"designmode/behavioralDesign/command"
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	// 用于测试，模拟来自客户端的事件
	eventChan := make(chan string)
	go func() {
		events := []string{"start", "archive", "start", "archive", "start", "start"}
		for _, e := range events {
			eventChan <- e
		}
	}()
	defer close(eventChan)

	// 使用命令队列缓存命令
	commands := make(chan command.ICommand, 1000)
	defer close(commands)

	go func() {
		for {
			// 从请求或者其他地方获取相关事件参数
			event, ok := <-eventChan
			if !ok {
				return
			}

			var icommand command.ICommand
			switch event {
			case "start":
				icommand = command.NewStartCommand()
			case "archive":
				icommand = command.NewArchiveCommand()
			}

			// 将命令入队
			commands <- icommand
		}
	}()

	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
