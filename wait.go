package wait

import (
	"sync"
)

// Group waits for a collection of goroutines to finish.
// Group is an extension of sync.WaitGroup that exposes a goroutine runner method, Go.
// Like sync.WaitGroup, a Group must not be copied after first use.
type Group struct {
	sync.WaitGroup
}

// Go runs fn in a new goroutine.
// Go increments g before fn starts, and decrements g after fn finishes.
func (g *Group) Go(fn func()) {
	g.Add(1)
	go func() {
		defer g.Done()
		fn()
	}()
}

// GroupWithCancellation waits for a collection of goroutines to finish.
// GroupWithCancellation can signal its goroutines to cancel their current activity.
// GroupWithCancellation is an extension of sync.WaitGroup that exposes a goroutine runner method, Go.
// Like sync.WaitGroup, a GroupWithCancellation must not be copied after first use.
type GroupWithCancellation struct {
	sync.WaitGroup
	mux    sync.Mutex
	cancel chan struct{}
}

// Go runs fn in a new goroutine.
// fn receives a chan argument which, upon cancellation, will be closed.
// Go increments g before fn starts, and decrements g after fn finishes.
func (g *GroupWithCancellation) Go(fn func(cancel <-chan struct{})) {
	g.mux.Lock()
	if g.cancel == nil {
		g.cancel = make(chan struct{})
	}
	cancel := g.cancel
	g.mux.Unlock()
	g.Add(1)
	go func() {
		defer g.Done()
		fn(cancel)
	}()
}

// Cancel sends a cancellation signal to all goroutines managed by g.
// Cancel can safely be called concurrently, but subsequent calls will have no effect
// unless there is an intermediate call to Go.
func (g *GroupWithCancellation) Cancel() {
	g.mux.Lock()
	defer g.mux.Unlock()
	if g.cancel != nil {
		close(g.cancel)
		g.cancel = nil
	}
}
