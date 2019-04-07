package array

import "testing"

type Test struct {
	in  [3]string
	out [3]string
}

var tests = []Test{
	{[3]string{"leslie", "luke", "yuzuki"}, [3]string{"leslie", "luke", "yuzuki"}},
	{[3]string{"name", "name2", "name3"}, [3]string{"name", "name2", "name3"}},
}

func TestName(t *testing.T) {
	for i, test := range tests {
		name := fill_array(test.in)
		if name != test.out {
			t.Errorf("#%d: fill_array(%s)=%s; want %s", i, test.in, name, test.out)
		}
	}
}
