package process

import (
	api "github.com/ylallemant/panopticon/pkg/api/v1"
)

func List() ([]*api.Process, error) {
	return listProcesses()
}
