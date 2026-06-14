package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
}

func (e EmailNotifier) Send(m string) error {
	fmt.Println(m)
	return nil
}

type SMSNotifier struct {
}

func (e SMSNotifier) Send(m string) error {
	fmt.Println(m)
	return nil
}

type TelegramNotifier struct {
}

func (e TelegramNotifier) Send(m string) error {
	fmt.Println(m)
	return nil
}

func main() {
	var n Notifier

	n = EmailNotifier{}
	n.Send("Hello mail")

	n = SMSNotifier{}
	n.Send("Hello sms")

	n = TelegramNotifier{}
	n.Send("Hello telegram")
}
