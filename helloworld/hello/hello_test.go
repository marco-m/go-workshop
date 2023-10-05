package hello_test

import (
	"testing"

	"github.com/go-quicktest/qt"

	"github.com/marco-m/go-project-templates/helloworld/hello"
)

func TestGreetSimple(t *testing.T) {
	got := hello.Greet("alice")

	qt.Assert(t, qt.Equals(got, "hello alice"))
}

// Although clearly overdoing for the implementation of Greet, this shows the
// typical approach to Go test: table-driven.
func TestGreetTable(t *testing.T) {
	type testCase struct {
		name string
		who  string
		want string
	}

	test := func(t *testing.T, tc testCase) {
		got := hello.Greet(tc.who)

		// Go tests tend to use:
		// "got" or "have" instead of "actual"
		// "want" instead of "expected".
		qt.Assert(t, qt.Equals(got, tc.want))
	}

	testCases := []testCase{
		{
			name: "works for fruits",
			who:  "banana",
			want: "hello banana",
		},
		{
			name: "works for aliens",
			who:  "Martian Bob",
			want: "hello Martian Bob",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { test(t, tc) })
	}
}
