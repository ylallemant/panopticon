package v1

func (u *User) GetUsernameByHost(hostname string) string {
	if username, exists := u.Devices[hostname]; exists {
		return username
	}

	return ""
}
