/*
 * @Author: bsun
 * @Date: 2024-07-16 09:34:15
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-17 09:47:05
 */

package mediator

import (
	"fmt"
	"reflect"
)

/*
	Package mediator 中介模式
	采用原课程的示例，并且做了一些裁剪
	假设我们现在有一个较为复杂的对话框，里面包括，登录组件，注册组件，以及选择框
	当选择框选择“登录”时，展示登录相关组件
	当选择框选择“注册”时，展示注册相关组件
*/

// input假设这表示一个输入框
type Input string

// String String
func (i Input) String() string {
	return string(i)
}

// Selection 假设这表示一个选择框
type Selection string

// Selected 当前选中的对象
func (s Selection) Selected() string {
	return string(s)
}

// Button 假设表示一个按钮
type Button struct {
	onClick func()
}

// SetOnClick 添加点击事件回调
func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

// IMediator 中介者接口
type IMediator interface {
	HandleEvent(component interface{})
}

// Dialog 对话框
type Dialog struct {
	// 按钮种类
	LoginButton *Button
	RegButton   *Button

	// 选择信息
	Selection *Selection
	// 属性信息
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

// HandleEvent 根据选择不同的按钮输出对应一组属性的信息
func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select ", d.Selection.Selected())
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select ", d.Selection.Selected())
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
		// others, 如果点击了登录按钮，注册按钮
	}
}
