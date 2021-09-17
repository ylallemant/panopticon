package process

import (
	"os"

	api "github.com/ylallemant/panopticon/pkg/api/v1"
)

func List() ([]*api.Process, error) {
	return listProcesses()
}

func Kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}
