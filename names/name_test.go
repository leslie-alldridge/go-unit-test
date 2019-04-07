package name

import "testing"

type Test struct {
	in  string
	out string
}

var tests = []Test{
	{"leslie", "that is my name"},
	{"roy", "guess again"},
	{"bob", "guess again"},
}

func TestName(t *testing.T) {
	for i, test := range tests {
		name := Name(test.in)
		if name != test.out {
			t.Errorf("#%d: Name(%s)=%s; want %s", i, test.in, name, test.out)
		}
	}
}
