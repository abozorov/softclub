package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Validator interface {
    Validate(value string) error
}

type EmailValidator struct {
}

func hasDog(s []rune) bool {
	for _, v := range s {
		if v == '@' {
			return true
		}
	}
	return false
}

func (e EmailValidator) Validate(m string) error {
	if len(m) >=5 && hasDog([]rune(m)) {
		return nil
	} else {
		return fmt.Errorf("Плохой маил")
	}
}

type PasswordValidator struct {
}

func hasNumber(s []rune) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

func (p PasswordValidator) Validate(m string) error {
	if len(m)>= 6 && hasNumber([]rune(m)) {
		return nil
	} else {
		return fmt.Errorf("Plohoi password")
	}
}

type UsernameValidator struct {
}

func (u UsernameValidator) Validate(m string) error {
	
	if m = strings.TrimSpace(m); len(m) >= 3 {
		return nil
	} else {
		return fmt.Errorf("Plohoi username")
	}
}

func main() {
	var n Validator

	n = EmailValidator{}
	n.Validate("anu@anu")

	n = PasswordValidator{}
	n.Validate("jhsdhjjh61")

	n = UsernameValidator{}
	n.Validate("username")
}
