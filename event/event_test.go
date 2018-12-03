package event

import (
	"fmt"
	"github.com/ethereum/go-ethereum/event"
	"sync"
	"testing"
)

func TestEvent(t *testing.T) {
	type someEvent struct {I int}
	var feed event.Feed
	var wg sync.WaitGroup

	ch := make(chan someEvent)
	sub := feed.Subscribe(ch)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range ch {
			fmt.Println("Receive: %#v", event.I)
		}
		sub.Unsubscribe()
		fmt.Println("done")
	}()
	feed.Send(someEvent{5})
	feed.Send(someEvent{10})
	feed.Send(someEvent{7})
	close(ch)
	wg.Wait()
}
