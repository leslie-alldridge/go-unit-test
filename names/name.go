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
