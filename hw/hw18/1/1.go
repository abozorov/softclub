package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Notification struct {
	User    string
	Message string
}

type Sender interface {
	Send(n Notification) error
}

func MessageSaver(s string) error {
	file, err := os.OpenFile("hw18/1/data.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	wrtr := bufio.NewWriter(file)
	wrtr.WriteString(s)
	wrtr.Flush()
	return nil
}

type EmailSender struct{}

func (e EmailSender) Send(n Notification) error {
	if strings.TrimSpace(n.Message) == "" {
		return errors.New("Empty message")
	}
	err := MessageSaver(fmt.Sprintf("Email sent to %s | Message %s\n", n.User, n.User))
	if err != nil {
		return err
	}
	return nil
}

type SmsSender struct{}

func (e SmsSender) Send(n Notification) error {
	if utf8.RuneCountInString(strings.TrimSpace(n.Message)) > 50 {
		return errors.New("SMS is to long")
	}
	err := MessageSaver(fmt.Sprintf("SMS sent to %s\n", n.Message))
	if err != nil {
		return err
	}
	return nil
}

func SendAll(senders []Sender, n Notification) []error {
	errs := make([]error, 0)
	for _, v := range senders {
		err := v.Send(n)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func main() {
	notifications := []Notification{
		{User: "admin", Message: "System backup completed successfully."},
		{User: "alice_w", Message: "Your order #1042 has been shipped!"},
		{User: "j_smith", Message: "Security alert: New login from unrecognized device."},
		{User: "bob_builder", Message: "Reminder: Your subscription expires in 3 days."},
		{User: "dev_user", Message: "Build failed: syntax error in main.go:24."},
		{User: "manager_01", Message: ""},
		{User: "support_tech", Message: "Ticket #772 has been reassigned to you."},
		{User: "guest_user", Message: "Welcome! Please verify your email address."},
		{User: "security_officer", Message: "Critical: Multiple failed login attempts detected."},
		{User: "test_account", Message: "Test ping notification."},
	}
	senders := []Sender{
		EmailSender{},
		SmsSender{},
	}

	for _, v := range notifications {
		errs := SendAll(senders, v)
		for _, err := range errs {
			fmt.Println(v, err)
		}
	}
}
