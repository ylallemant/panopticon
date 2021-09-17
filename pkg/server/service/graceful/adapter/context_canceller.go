package adapter

import (
	"context"
	"fmt"
	"time"
)

type ContextCanceller struct {
	name       string
	cancelFunc context.CancelFunc
	timeout    time.Duration
}

func NewContextCanceller(name string, cancelFunc context.CancelFunc, timeout time.Duration) *ContextCanceller {
	return &ContextCanceller{
		name:       name,
		cancelFunc: cancelFunc,
		timeout:    timeout,
	}
}

func (w *ContextCanceller) Name() string {
	return fmt.Sprintf("%s (via ContextCanceller)", w.name)
}

func (w *ContextCanceller) Start() error {
	return nil
}

func (w *ContextCanceller) Stop() {
	w.cancelFunc()
	time.Sleep(w.timeout)
}
