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

func TestName(t *testing.T) {
	for i, test := range tests {
		name := Name(test.in)
		if name != test.out {
			t.Errorf("#%d: Name(%s)=%s; want %s", i, test.in, name, test.out)
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
