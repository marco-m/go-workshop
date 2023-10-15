package eventloop_test

import (
	"context"
	"testing"

	"github.com/go-quicktest/qt"

	"github.com/marco-m/go-workshop/eventloop/pkg/eventloop"
)

func TestAcceptance(t *testing.T) {
	// loop
	// pass function
	// takes other func as closure (the http client)
	// in the test we pass a testhttp server to the client as url
	// then we need to abstract reading from the testhttp server
	// channels to enqueue and to get back, like in the talk by peter bourgon
	// loop has a select on the channels
	// need to think more about the size of the channel...
	// but here from the test what do i see? how do i control the loop?
	// well with a classic foo.Loop() or foo.Run(), blocking ?
	// or (the same) Start and Wait ?
	// this is a sort of IoC Inversion of Control
	// func with a closure can pass anything for example the testserver url

	// so here this i want to write a super basic acceptance test with a silly function

	//ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello, client")
	//}))
	//defer ts.Close()
	//
	//type MyEvent struct {
	//}
	//inputCh := make(chan *MyEvent)
	//handler := func(ev *MyEvent) {
	//	http.Post(ts.URL, "application/json", &buf)
	//}
	//loop := eventloop.New(inputCh, handler)

	loop := eventloop.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, 1)
	go func() {
		errCh <- loop.Run(ctx)
	}()

	got, err := loop.Dispatch(1)

	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, qt.Equals(got, "1"))
	qt.Assert(t, qt.Equals(loop.State(), eventloop.StateSleeping))

	cancel()
	err = <-errCh
	qt.Assert(t, qt.Equals(err, context.Canceled))
}
