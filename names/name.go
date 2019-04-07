package name

import (
	"regexp"
)

func Name(a string) string {
	if a == "leslie" {
		return "that is my name"
	} else {
		return "guess again"
	}
}

func Email(e string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(e)
}

func GetAge(name string) int {
	if name == "leslie" {
		return 26
	} else if name == "bob" {
		return 21
	} else {
		return 22
	}
}
