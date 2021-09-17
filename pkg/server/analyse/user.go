package analyse

import (
	"log"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/server/cache"
)

func getTrackedUsers(host *v1.Host, cache cache.Cache) ([]*v1.User, error) {
	list := make([]*v1.User, 0)

	for userIdentifier := range host.GetUserMapping() {
		user, err := cache.GetUser(userIdentifier)
		if err != nil {
			return list, err
		}

		if user == nil {
			log.Printf("warning: user %s (%s) not found", userIdentifier, user.GetUsernameByHost(host.Name))
			continue
		}

		list = append(list, user)
	}

	return list, nil
}

func identifyNewTrackedUsers(host *v1.Host, cache cache.Cache) {
	for _, user := range cache.GetUsers() {
		username, exists := user.Devices[host.Name]
		if !exists {
			// user has no account on device
			continue
		}

		_, exists = host.UserIdByIdentifier(user.Identifier)
		if exists {
			// user already registered for device
			continue
		}

		deviceUser, found := host.UserByName(username)
		if found {
			log.Printf("identified new tracked user %s (%s, %d) for device %s",
				user.Identifier, deviceUser.Name, deviceUser.Id, host.Name)
			host.UserMapping[user.Identifier] = deviceUser
		}
	}
}
