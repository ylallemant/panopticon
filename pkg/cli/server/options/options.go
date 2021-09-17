package options

var Current = NewOptions()

func NewOptions() *Options {
	options := new(Options)

	options.Ports = Ports{
		GRPC:        int32(1984),
		HTTP:        int32(1989),
		Maintenance: int32(2012),
	}

	return options
}

type Ports struct {
	GRPC        int32
	HTTP        int32
	Maintenance int32
}

type Options struct {
	Ports Ports
}
