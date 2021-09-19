package options

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

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

func (o *Options) Load(path string) (*Options, error) {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, o)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	return o, nil
}
