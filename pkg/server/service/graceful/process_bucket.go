package graceful

import (
	"os"
	"sync"

	"log"
)

type Process interface {
	Name() string
	Start() error
	Stop()
}

func NewProcessBucket() *bucket {
	return &bucket{
		signalCh:  make(chan os.Signal),
		stopCh:    make(chan struct{}),
		processes: make([]Process, 0),
		mu:        sync.RWMutex{},
	}
}

type bucket struct {
	signalCh  chan os.Signal
	stopCh    chan struct{}
	processes []Process
	mu        sync.RWMutex
}

func (b *bucket) AddProcess(p Process) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.processes = append(b.processes, p)
}

func (b *bucket) Start() error {
	for n, _ := range b.processes {
		log.Printf("Starting process %s", b.processes[n].Name())
		/*
			if err := b.processes[n].Start(); err != nil {
				return err
			}
		*/
		go b.processes[n].Start()
	}
	return nil
}

func (b *bucket) Shutdown() {
	for n, _ := range b.processes {
		log.Printf("Shutting down process %s", b.processes[n].Name())
		b.processes[n].Stop()
	}
	b.stopCh <- struct{}{}
}

func (b *bucket) SigChan() chan os.Signal {
	return b.signalCh
}

func (b *bucket) Wait() {
	<-b.stopCh
}

func (b *bucket) Init() {
	go func() {
		var shutdownInProgress bool

		for {
			select {
			case sig := <-b.signalCh:
				if shutdownInProgress {
					log.Printf("Shutdown procedure interrupted by user (signal %s), exit now.", sig.String())
					os.Exit(1)
				}

				log.Printf("Received signal %s, initiate shutdown procedure", sig.String())
				shutdownInProgress = true
				go b.Shutdown()
			}
		}
	}()
}
