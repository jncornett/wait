package wait

import (
	"testing"
)

func TestGroup_Go(t *testing.T) {
	work := make(chan bool)
	var wg Group
	wg.Go(func() {
		work <- true
	})
	<-work
	wg.Wait()
}

func TestGroupWithCancellation_Go(t *testing.T) {
	work := make(chan bool)
	var wg GroupWithCancellation
	wg.Go(func(<-chan struct{}) {
		work <- true
	})
	<-work
	wg.Wait()
}

func TestGroupWithCancellation_Cancel(t *testing.T) {
	var wg GroupWithCancellation
	wg.Go(func(cancel <-chan struct{}) {
		<-cancel
	})
	wg.Cancel()
	wg.Wait()
	wg.Go(func(cancel <-chan struct{}) {
		<-cancel
	})
	wg.Cancel()
	wg.Wait()
}
