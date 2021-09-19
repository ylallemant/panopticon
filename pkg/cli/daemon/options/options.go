package options

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

var Current = NewOptions()

func NewOptions() *Options {
	options := new(Options)

	options.Period = time.Duration(30 * time.Second)

	return options
}

type Options struct {
	Period     time.Duration
	ConfigPath string
	Endpoint   string
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
