package classifier

import (
	"regexp"
	"sync"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/server/cache"
)

func NewRegexpClassifier(cache cache.Cache) *RegexpClassifier {
	instance := new(RegexpClassifier)

	instance.matrix = make(map[string][]*regexp.Regexp)
	instance.cache = cache

	return instance
}

var _ Classifier = &RegexpClassifier{}

type RegexpClassifier struct {
	matrix  map[string][]*regexp.Regexp
	cache   cache.Cache
	stateID int64
	mux     sync.RWMutex
}

func (c *RegexpClassifier) syncMatrix() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.matrix = make(map[string][]*regexp.Regexp, 0)

	for _, application := range c.cache.GetApplications() {
		err := c.register(application)
		if err != nil {
			return err
		}
	}

	c.stateID = c.cache.StateID()

	return nil
}

func (c *RegexpClassifier) register(application *v1.Application) error {
	regexps := make([]*regexp.Regexp, 0)

	for _, current := range application.Regexps {
		definition, err := regexp.Compile(current)
		if err != nil {
			return err
		}

		regexps = append(regexps, definition)
	}

	c.matrix[application.Name] = regexps

	return nil
}

func (c *RegexpClassifier) Classify(host *v1.Host, process *v1.Process) (*v1.Application, error) {
	if c.cache.StateID() != c.stateID {
		err := c.syncMatrix()
		if err != nil {
			return nil, err
		}
	}

	c.mux.RLock()
	defer c.mux.RUnlock()

	for applicationName, regexps := range c.matrix {
		for _, current := range regexps {
			if !current.Match([]byte(process.GetCommand())) {
				continue
			}

			return c.cache.GetApplication(applicationName)
		}
	}

	return nil, nil
}
