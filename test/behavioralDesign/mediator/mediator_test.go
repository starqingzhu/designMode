/*
 * @Author: bsun
 * @Date: 2024-07-17 09:42:47
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-17 09:45:56
 */

package test

import (
	"designmode/behavioralDesign/mediator"
	"testing"
)

func TestMediator(t *testing.T) {
	usernameInput := mediator.Input("username input")
	passwordInput := mediator.Input("password input")
	repeatPwdInput := mediator.Input("repeat password input")

	selection := mediator.Selection("登录")
	d := &mediator.Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPwdInput,
	}
	d.HandleEvent(&selection)

	regSelection := mediator.Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
