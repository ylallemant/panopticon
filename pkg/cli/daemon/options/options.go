package options

import "time"

var Current = NewOptions()

func NewOptions() *Options {
	options := new(Options)

	options.Period = time.Duration(30 * time.Second)

	return options
}

type Options struct {
	Period   time.Duration
	Endpoint string
}
