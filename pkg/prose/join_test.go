package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithComas(t *testing.T) {
	tests := []testData{
		{list: []string{}, want: ""},
		{list: []string{"orange"}, want: "orange"},
		{list: []string{"orange", "apple"}, want: "orange and apple"},
		{list: []string{"orange", "apple", "pear"}, want: "orange, apple, and pear"},
	}

	for _, test := range tests {
		got := JoinWithComas(test.list)
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithComas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
