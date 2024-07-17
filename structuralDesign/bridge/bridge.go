package bridge

import "fmt"

type IMsgSender interface {
	Send(msg string) error
}

// emailMsgSender 发送邮件
type EmailMsgSender struct {
	emails []string
}

// new
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	//发送消息逻辑
	fmt.Println("msg: ", msg)
	return nil
}

type INotifycation interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
