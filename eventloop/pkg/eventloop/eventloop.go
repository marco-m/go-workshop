package eventloop

import (
	"context"
	"strconv"
)

type State int

const (
	StateSleeping State = iota
	StateEating
)

func (st State) String() string {
	switch st {
	case StateSleeping:
		return "sleeping"
	case StateEating:
		return "eating"
	default:
		return "invalid"
	}
}

type EventLoop struct {
	state    State
	actionCh chan func()
}

func New() *EventLoop {
	return &EventLoop{
		actionCh: make(chan func()),
	}
}

func (sm *EventLoop) State() State {
	return sm.state
}

func (sm *EventLoop) Run(ctx context.Context) error {
	for {
		select {
		case fn := <-sm.actionCh:
			fn()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (sm *EventLoop) Dispatch(ev int) (string, error) {
	type ret struct {
		string
		error
	}
	ch := make(chan ret)

	sm.actionCh <- func() {
		r := strconv.Itoa(ev)
		if sm.state == StateSleeping {
			sm.state = StateEating
		} else {
			sm.state = StateSleeping
		}
		ch <- ret{r, nil}
	}

	r := <-ch
	return r.string, r.error
}
