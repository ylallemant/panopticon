package classifier

import (
	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
)

type Classifier interface {
	Classify(*v1.Host, *v1.Process) (*v1.Application, error)
}
