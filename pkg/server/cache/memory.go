package cache

import (
	"sync"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
)

func NewMemoryCache() *memory {
	cache := new(memory)

	cache.user = make(map[string]*v1.User)
	cache.host = make(map[string]*v1.Host)
	cache.application = make(map[string]*v1.Application)
	cache.userLimits = make(map[string]*v1.UserLimits)
	cache.userReports = make(map[string]*v1.UserDailyReport)

	cache.PutUser(yann.Identifier, yann)
	cache.PutApplication(firefox.Name, firefox)
	cache.PutUserLimits(yannLimits.Identifier, yannLimits)

	cache.stateID = int64(0)

	return cache
}

var (
	_    Cache = &memory{}
	yann       = &v1.User{
		Identifier: "yann.lallemant@gmail.com",
		Username:   "Yann",
		Devices: map[string]string{
			"Yanns-MacBook-Pro.local": "y.lallemant",
		},
	}
	firefox = &v1.Application{
		Name: "Firefox",
		Regexps: []string{
			"firefox",
		},
		DefaultDurationLimits: []*v1.DailyDurationLimit{},
		DefaultTimeLimits:     []*v1.DailyTimeLimit{},
	}
	yannLimits = &v1.UserLimits{
		Identifier: "yann.lallemant@gmail.com",
	}
)

type memory struct {
	user        map[string]*v1.User
	host        map[string]*v1.Host
	application map[string]*v1.Application
	userLimits  map[string]*v1.UserLimits
	userReports map[string]*v1.UserDailyReport
	stateID     int64
	mux         sync.RWMutex
}

func (m *memory) StateID() int64 {
	m.mux.RLock()
	defer m.mux.RUnlock()

	return m.stateID
}

func (m *memory) GetUser(name string) (*v1.User, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if user, ok := m.user[name]; ok {
		return user, nil
	}

	return nil, nil
}

func (m *memory) PutUser(name string, user *v1.User) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.user[name] = user
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) DeleteUser(name string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.user, name)
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) GetUsers() map[string]*v1.User {
	m.mux.RLock()
	defer m.mux.RUnlock()

	return m.user
}

func (m *memory) GetHost(name string) (*v1.Host, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if host, ok := m.host[name]; ok {
		return host, nil
	}

	return nil, nil
}

func (m *memory) PutHost(name string, host *v1.Host) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.host[name] = host
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) DeleteHost(name string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.host, name)
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) GetApplication(name string) (*v1.Application, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if application, ok := m.application[name]; ok {
		return application, nil
	}

	return nil, nil
}

func (m *memory) PutApplication(name string, application *v1.Application) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.application[name] = application
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) DeleteApplication(name string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.application, name)
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) GetApplications() map[string]*v1.Application {
	m.mux.RLock()
	defer m.mux.RUnlock()

	return m.application
}

func (m *memory) GetUserLimits(name string) (*v1.UserLimits, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if limits, ok := m.userLimits[name]; ok {
		return limits, nil
	}

	return nil, nil
}

func (m *memory) IsUserLimited(name string) bool {
	m.mux.RLock()
	defer m.mux.RUnlock()

	_, exists := m.userLimits[name]

	return exists
}

func (m *memory) PutUserLimits(name string, limits *v1.UserLimits) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.userLimits[name] = limits
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) DeleteUserLimits(name string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.userLimits, name)
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) GetUserDailyReport(name string) (*v1.UserDailyReport, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	if report, ok := m.userReports[name]; ok {
		return report, nil
	}

	return nil, nil
}

func (m *memory) PutUserDailyReport(name string, report *v1.UserDailyReport) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	m.userReports[name] = report
	m.stateID = chronos.TimestampNano()

	return nil
}

func (m *memory) DeleteUserDailyReport(name string) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	delete(m.userReports, name)
	m.stateID = chronos.TimestampNano()

	return nil
}
