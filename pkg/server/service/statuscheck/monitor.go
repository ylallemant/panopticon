package statuscheck

const (
	Good Status = "good"
	Bad         = "bad"
)

type Monitor interface {
	Status() Status
}

func NewLivenessMonitor() *alwaysGoodMonitor {
	return NewAlwaysGoodMonitor()
}

func NewReadinessMonitor() *lockingStatusMonitor {
	return NewLockingStatusMonitor("ReadinessMonitor")
}
