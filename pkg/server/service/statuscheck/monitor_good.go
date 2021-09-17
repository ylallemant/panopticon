package statuscheck

func NewAlwaysGoodMonitor() *alwaysGoodMonitor {
	return new(alwaysGoodMonitor)
}

type alwaysGoodMonitor struct{}

func (g *alwaysGoodMonitor) Status() Status {
	return Good
}
