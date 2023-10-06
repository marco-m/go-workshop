// The file name ends with "_test.go", so this is a test file.
// The package name ends with "_test", so this file has access only to the
// public API of package "banana".
//
// The majority of the tests should happen here, at the public boundary.

package banana_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-quicktest/qt"

	"github.com/marco-m/go-project-templates/fruits/pkg/banana"
)

// The simplest possible test.
// We use the AAA (Arrange, Act, Assert) convention of separating each step with
// an empty line, to make it apparent what is happening (normally one would not
// comment each step as is done in this example).
func TestABananaCanBePeeled(t *testing.T) {
	// Arrange
	ladyFinger := banana.Banana{}

	// Act
	err := ladyFinger.Peel()

	// Assert
	qt.Assert(t, qt.IsNil(err))
}

// Go does not support the xUnit convention of setup and teardown as methods of
// a test class or test case. Instead, it is customary to use a table driven test.
// The setup and teardown happen either once at beginning and end of the test, or
// per each testcase. It is up to the programmer to do so.
//
// This particular layout is slightly different from the standard Go examples.
// I find this approach better for readability and discoverability.
// More information about the rationale at the very good
// https://github.com/gotestyourself/gotest.tools/wiki/Go-Testing-Patterns
//
// This particular example is quite silly given the banana API...
func TestBananas(t *testing.T) {
	type testCase struct {
		name       string
		cuts       int
		wantPieces []string
	}

	test := func(t *testing.T, tc testCase) {
		sut := banana.Banana{}
		err := sut.Peel()
		qt.Assert(t, qt.IsNil(err))

		got := sut.Cut(tc.cuts)

		// Go tests tend to use:
		// "got" or "have" instead of "actual"
		// "want" instead of "expected".
		qt.Assert(t, qt.DeepEquals(got, tc.wantPieces))
	}

	testCases := []testCase{
		{
			name:       "zero cuts -> one slice",
			cuts:       0,
			wantPieces: []string{"p"},
		},
		{
			name:       "one cut -> two slices",
			cuts:       1,
			wantPieces: []string{"p", "p"},
		},
		{
			name:       "7 cuts -> 8 slices",
			cuts:       7,
			wantPieces: []string{"p", "p", "p", "p", "p", "p", "p", "p"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { test(t, tc) })
	}
}

// This shows dependency injection to insert a simple test double.
func TestBananaWithSpy(t *testing.T) {
	spy := SleepSpy{}
	pan := banana.NewPan(banana.PanOpts{SleepFn: spy.Sleep})
	sut := banana.Banana{}

	pan.Cook(&sut, 42*time.Second)
	pan.Cook(&sut, 5*time.Second)

	want := []time.Duration{42 * time.Second, 5 * time.Second}
	qt.Assert(t, qt.DeepEquals(spy.sleeps, want))
}

// Example integration test, that is, a test that takes time or requires
// authentication. We show how to use the Skip method.
func TestBananaIntegration(t *testing.T) {
	tokenName := "BANANA_TEST_TOKEN"
	token := os.Getenv(tokenName)
	if token == "" {
		t.Skip("Skipping integration test: missing env var " + tokenName)
	}

	// TODO: write actual test here...
}

type SleepSpy struct {
	sleeps []time.Duration
}

func (spy *SleepSpy) Sleep(d time.Duration) {
	spy.sleeps = append(spy.sleeps, d)
}
