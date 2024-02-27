package types

import (
	"strings"
)

func Strings(subject string, search string, replace string) MyString {

	newstring := strings.Join(strings.Split(subject, search), replace)

	return MyString{
		subject,
		newstring,
		len(newstring),
	}
}

type MyString struct {
	input  string
	output string
	length int
}
