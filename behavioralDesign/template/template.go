package template

import "fmt"

/*
举个 🌰，假设我现在要做一个短信推送的系统，那么需要

检查短信字数是否超过限制
检查手机号是否正确
发送短信
返回状态
我们可以发现，在发送短信的时候由于不同的供应商调用的接口不同，所以会有一些实现上的差异，但是他的算法（业务逻辑）是固定的

*/

type ISMS interface {
	send(content string, phone int) error
}

// SMS 短信发送基类
type sms struct {
	ISMS
}

// Valid 校验短信字数
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// 发送
func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	// 调用子类的方法发送短信
	return s.send(content, phone)
}

// TelecomSms 走电信通道
type TelecomSms struct {
	*sms
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Printf("send %s to %d\n", content, phone)
	return nil
}

// NewTelecomSms
func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}
