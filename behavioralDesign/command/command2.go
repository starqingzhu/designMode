/*
 * @Author: bsun
 * @Date: 2024-07-10 15:12:32
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-10 16:17:19
 */
package command

import "fmt"

// Command 命令
type Command func() error

// StartCommandFunc 返回一个 Command 命令
// 是因为正常情况下不会是这么简单的函数
// 一般都会有一些参数
func StartCommandFunc() Command {
	return func() error {
		fmt.Println("game start")
		return nil
	}
}

// ArchiveCommandFunc ArchiveCommandFunc
func ArchiveCommandFunc() Command {
	return func() error {
		fmt.Println("game archive")
		return nil
	}
}
