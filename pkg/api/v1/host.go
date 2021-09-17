package v1

func (h *Host) UserIdByIdentifier(identifier string) (int32, bool) {
	if user, exists := h.UserMapping[identifier]; exists {
		return user.Id, true
	}

	return -1, false
}

func (h *Host) UsernameByIdentifier(identifier string) (string, bool) {
	if user, exists := h.UserMapping[identifier]; exists {
		return user.Name, true
	}

	return "", false
}

func (h *Host) UserIdExists(id int32) bool {
	return userExistsByID(id, h.Users)
}

func (h *Host) AdminIdExists(id int32) bool {
	return userExistsByID(id, h.Admins)
}

func userExistsByID(id int32, list []*HostUser) bool {
	for _, user := range list {
		if user.Id == id {
			return true
		}
	}

	return false
}

func (h *Host) UserByName(username string) (*HostUser, bool) {
	for _, user := range h.Users {
		if user.Name == username {
			return user, true
		}
	}

	return nil, false
}

func (h *Host) UserByID(uid int32) (*HostUser, bool) {
	for _, user := range h.Users {
		if user.Id == uid {
			return user, true
		}
	}

	return nil, false
}
