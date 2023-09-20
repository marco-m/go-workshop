// Package banana offers some bananas.
// According to wikipedia, "(banana) cultivar names are highly confused", which
// I find particularly apt for software :-)
package banana

import (
	"errors"
	"math/rand"
	"time"
)

var ErrAlreadyPeeled = errors.New("already peeled")

type Banana struct {
	cookedFor time.Duration
	peeled    bool
}

// Peel peels a banana. It is an error peeling more than once.
//
// Why it is an error peeling more than once ??? Just as an example; it would
// probably be better instead to "define the error out of existence" and simply
// do nothing in case the banana is already peeled!
// For the concept of "defining errors out of existence", see
//   - the excellent book by John Ousterhout, "A Philosophy of Software Design"
//   - a presentation of the book: https://www.youtube.com/watch?v=bmSAYlu0NcY
//   - an article inspired by the book, for Go:
//     https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors
func (ba *Banana) Peel() error {
	if ba.peeled {
		return ErrAlreadyPeeled
	}
	ba.peeled = true
	return nil
}

// Cut cuts a banana into pieces. On purpose the API is bad: it requires as
// parameter the number of cuts, instead of the number of pieces we want. This
// will lead to the classic off-by-one or fencepost error, which we show in the
// tests.
//
// To show again the importance of error handling, the API needs more thinking
// and as a minimum a better rationale and explanation: it does not return an
// error if the banana has already been cut or if the required number of cuts is
// negative. What should it do instead? For contrast, compare with the reasoning
// in Banana.Peel.
func (ba *Banana) Cut(cuts int) []string {
	if cuts < 0 {
		cuts = 0
	}
	filler := "s" // s = skin
	if ba.peeled {
		filler = "p" // p = peeled
	}

	// Silly bug introduced to show the Go fuzzing support; see the tests.
	if cuts == 42 {
		cuts++
	}

	pieces := make([]string, cuts+1)
	for i := range pieces {
		pieces[i] = filler
	}
	return pieces
}

func (ba *Banana) Color() string {
	if !ba.peeled {
		if rand.Intn(10) < 5 {
			return "sometimes yellow"
		}
		return "sometimes brown"
	}
	return "more or less white"
}

type SleepFunc func(d time.Duration)

type PanOpts struct {
	SleepFn SleepFunc // Optional; overridable for testing.
}

// Pan is a pan to cook bananas. Use NewPan to instantiate; do not use directly.
type Pan struct {
	PanOpts
}

// NewPan returns a pan ready to cook bananas.
//
// Using a function prefixed New to construct an instance of a struct is only a
// convention, there is no way to enforce it. When the struct _must_ be
// initialized before being used, this is obtained purely by convention; see also
// the comment for [Pan].
func NewPan(opts PanOpts) *Pan {
	if opts.SleepFn == nil { // Example of how to handle default values.
		opts.SleepFn = time.Sleep
	}
	pan := &Pan{opts}

	return pan
}

func (pa Pan) Cook(banana *Banana, howlong time.Duration) {
	pa.SleepFn(howlong)
	banana.cookedFor += howlong
}
