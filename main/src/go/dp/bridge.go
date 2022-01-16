package main

// 不同的通知类型, 有不同的报警级别  (即两个维度)

// send接口 不同的通知类型
type IMsgSender interface {
	Send(msg string) error
}

// 发送邮件; 还可以扩展 电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

// notify接口 不同的报警级别
type INotification interface {
	Notify(msg string) error
}

// 错误通知; 还可以扩展 warning 等级别
type ErrorNotification struct {
	sender IMsgSender // send接口
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send("error: " + msg)
}
