package daemon

import (
	"log"
	"time"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
	"github.com/ylallemant/panopticon/pkg/cli/daemon/options"
	"github.com/ylallemant/panopticon/pkg/daemon/coerce"
	"github.com/ylallemant/panopticon/pkg/daemon/process"
	"github.com/ylallemant/panopticon/pkg/server/client"
	"github.com/ylallemant/panopticon/pkg/server/service/graceful"
)

func NewDaemon(options *options.Options) (*daemon, error) {
	process := new(daemon)

	process.stopCh = make(chan struct{})
	process.period = options.Period

	client, err := client.NewDaemonClient(options.Endpoint, &client.ClientOptions{})
	if err != nil {
		return nil, err
	}

	process.client = client

	return process, nil
}

var _ graceful.Process = &daemon{}

type daemon struct {
	client client.DaemonClient
	period time.Duration
	stopCh chan struct{}
}

func (d *daemon) Name() string {
	return "panopticon-daemon"
}

func (d *daemon) Start() error {
	go d.run()
	return nil
}

func (d *daemon) Stop() {
	d.client.Stop()
	d.stopCh <- struct{}{}
}

func (d *daemon) run() {
	duration := time.Second * 2

	for {
		select {
		case <-time.After(duration):
			if err := d.tick(); err != nil {
				log.Fatalf("Loop turn broke with error: %v", err)
			}
		case <-d.stopCh:
			return
		}

		duration = d.period
	}
}

func (d *daemon) tick() error {
	log.Print("new tick")
	list, err := process.List()
	if err != nil {
		return err
	}

	if len(list) == 0 {
		log.Printf("empty process list - nothing to report")
		return nil
	}

	log.Printf("%d processe found", len(list))
	request := &v1.HostProcessReportRequest{
		Report: &v1.HostProcessReport{
			Hostname:  d.client.Metadata().Hostname(),
			Arch:      d.client.Metadata().Arch(),
			Os:        d.client.Metadata().OS(),
			Timestamp: chronos.TimestampNano(),
			Interval:  options.Current.Period.Nanoseconds(),
			Processes: list,
		},
	}

	log.Printf("send report to server via %s", d.client.Type())
	response, err := d.client.Report(request)
	if err != nil {
		return err
	}

	err = coerce.Coerce(response)
	if err != nil {
		return err
	}

	return nil
}
