package synk

import (
	"sync"
)

type ChannelManager[T any] struct {
	Channel chan T
	wg      *sync.WaitGroup
}

func OpenBufferChannel[T any](size int) *ChannelManager[T] {
	ch := make(chan T, size)
	var wg sync.WaitGroup
	wg.Add(size)
	cm := &ChannelManager[T]{
		Channel: ch,
		wg:      &wg,
	}
	go cm.close()
	return cm
}

func (cm *ChannelManager[T]) Send(v T) {
	defer cm.wg.Done()
	cm.Channel <- v
}

func (cm *ChannelManager[T]) ReceiveWait() []T {
	var result []T
	for v := range cm.Channel {
		result = append(result, v)
	}
	return result
}

func (cm *ChannelManager[T]) close() {
	cm.wg.Wait()
	close(cm.Channel)
}
