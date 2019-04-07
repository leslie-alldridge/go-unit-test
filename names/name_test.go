package name

import "testing"

type Test struct {
	in  string
	out string
}

type Test2 struct {
	in  string
	out bool
}

type Test3 struct {
	in  string
	out int
}

var tests = []Test{
	{"leslie", "that is my name"},
	{"roy", "guess again"},
	{"bob", "guess again"},
}

var tests2 = []Test2{
	{"ç$€§/az@gmail.com", false},
	{"abcd@gmail_yahoo.com", false},
	{"abcd@gmail-yahoo.com", true},
	{"abcd@gmailyahoo", true},
	{"abcd@gmail.yahoo", true},
}

var tests3 = []Test3{
	{"leslie", 26},
	{"roy", 22},
	{"bob", 21},
}

func TestName(t *testing.T) {
	for i, test := range tests {
		name := Name(test.in)
		if name != test.out {
			t.Errorf("#%d: Name(%s)=%v; want %v", i, test.in, name, test.out)
		}
	}
}

func TestEmail(t *testing.T) {
	for i, test2 := range tests2 {
		email := Email(test2.in)
		if email != test2.out {
			t.Errorf("#%d: Email(%s)=%v; want %v", i, test2.in, email, test2.out)
		}
	}
}

func TestGetAge(t *testing.T) {
	for i, test3 := range tests3 {
		getage := GetAge(test3.in)
		if getage != test3.out {
			t.Errorf("#%d: GetAge(%s)=%d; want %d", i, test3.in, getage, test3.out)
		}
	}
}
