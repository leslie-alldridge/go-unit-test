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

// basic logic variant

// func GetAge(name string) int {
// 	if name == "leslie" {
// 		return 26
// 	} else if name == "bob" {
// 		return 21
// 	} else {
// 		return 22
// 	}
// }

//object variant

type human struct {
	name string
	age  int
}

// func GetAge(name string) int {
// 	people := human{name: "leslie", age: 26}
// 	return people.age
// }

// array of object variant
func GetAge(name string) int {
	var people []human
	people[0] = human{name: "leslie", age: 26}
	people[1] = human{name: "luke", age: 22}
	var age = 0
	for i := range people {
		if people[i].name == "luke" {
			age = people[i].age
		}
	}
	return age
}
