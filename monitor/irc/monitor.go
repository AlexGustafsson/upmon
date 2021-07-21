package irc

type Monitor struct {
}

func NewMonitor() (*Monitor, error) {
	return &Monitor{}, nil
}

func (monitor *Monitor) Name() string {
	return "irc"
}

func (monitor *Monitor) Description() string {
	return "IRC monitor"
}

func (monitor *Monitor) Version() string {
	return "0.1.0"
}
