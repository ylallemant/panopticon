package coerce

import (
	"log"
	"os"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
)

func Coerce(action *v1.HostActionResponse) error {

	for _, process := range action.Processes {
		err := Kill(int(process.GetPID()))
		if err != nil {
			return err
		}

		log.Printf("killed process #%d because of: %s", process.GetPID(), process.GetReason())
	}

	for _, user := range action.Users {
		err := Kill(int(user.GetUserID()))
		if err != nil {
			return err
		}

		log.Printf("logged out user #%d because of: %s", user.GetUserID(), user.GetReason())
	}

	return nil
}

func Kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func Logout(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}
