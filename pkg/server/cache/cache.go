package cache

import v1 "github.com/ylallemant/panopticon/pkg/api/v1"

type CacheUser interface {
	GetUser(string) (*v1.User, error)
	PutUser(string, *v1.User) error
	DeleteUser(string) error
	GetUsers() map[string]*v1.User
}

type CacheHost interface {
	GetHost(string) (*v1.Host, error)
	PutHost(string, *v1.Host) error
	DeleteHost(string) error
}

type CacheApplication interface {
	GetApplication(string) (*v1.Application, error)
	PutApplication(string, *v1.Application) error
	DeleteApplication(string) error
	GetApplications() map[string]*v1.Application
}

type CacheUserLimits interface {
	IsUserLimited(string) bool
	GetUserLimits(string) (*v1.UserLimits, error)
	PutUserLimits(string, *v1.UserLimits) error
	DeleteUserLimits(string) error
}

type CacheUserDailyReport interface {
	GetUserDailyReport(string) (*v1.UserDailyReport, error)
	PutUserDailyReport(string, *v1.UserDailyReport) error
	DeleteUserDailyReport(string) error
}

type Cache interface {
	CacheUser
	CacheHost
	CacheApplication
	CacheUserLimits
	CacheUserDailyReport
	StateID() int64
}
