package main

import (
	"strings"
	"testing"
)

func TestGetResourceListFromConfig(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "empty config",
			content:  "",
			expected: []string{},
		},
		{
			name:     "normal list",
			content:  "foo\nbar",
			expected: []string{"foo", "bar"},
		},
		{
			name:     "normal list with whitespace",
			content:  "foo  \n\n  bar\n",
			expected: []string{"foo", "bar"},
		},
	}

	for _, tst := range tests {
		actual, err := GetResourceListFromConfig(strings.NewReader(tst.content))
		if err != nil {
			t.Errorf("%s failed.Error:%s", tst.name, err)
		}
		if !equal(actual, tst.expected) {
			t.Errorf("%s failed.\nExpected:\n%s\nGot:\n%s", tst.name, tst.expected, actual)
		}
	}
}
func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
