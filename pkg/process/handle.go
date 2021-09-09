package process

import (
	"os"

	"github.com/ylallemant/panopticon/pkg/api"
)

func List() ([]*api.Process, error) {
	_, err := listProcesses()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func Kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}
