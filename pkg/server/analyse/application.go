package analyse

import (
	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/server/classifier"
)

func classifyProcesses(host *v1.Host, processes []*v1.Process, classifier classifier.Classifier) ([]*v1.ClassifiedProcess, error) {
	list := make([]*v1.ClassifiedProcess, 0)

	for _, process := range processes {
		application, err := classifier.Classify(host, process)
		if err != nil {
			return list, err
		}

		if application != nil {
			list = append(list, &v1.ClassifiedProcess{
				Process:     process,
				Application: application,
			})
		}
	}

	return list, nil
}
