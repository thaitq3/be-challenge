package utils

import "time"

// WaitChannels returns a channel which will be closed when receiving from all input channels.
func WaitChannels(chs ...<-chan struct{}) <-chan struct{} {
	ret := make(chan struct{})
	go func() {
		for _, ch := range chs {
			<-ch
		}
		close(ret)
	}()
	return ret
}

// WaitOrTimeout waits for all channels or timeout.
func WaitOrTimeout(timeout time.Duration, chs ...<-chan struct{}) {
	select {
	case <-WaitChannels(chs...):
	case <-time.After(timeout):
	}
}