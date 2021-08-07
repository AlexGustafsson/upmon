package storage

import (
	"sync"

	"github.com/AlexGustafsson/upmon/monitoring"
)

type Store struct {
	sync.Mutex
	Origins map[string]*Origin
}

type Origin struct {
	sync.Mutex
	Id       string
	Services map[string]*Service
}

type Service struct {
	sync.Mutex
	Id       string
	Monitors map[string]*Monitor
}

type Monitor struct {
	sync.Mutex
	Id       string
	Statuses map[string]monitoring.Status
}

type ServiceStatus struct {
}

type MonitorStatus struct {
	Statuses map[monitoring.Status]int
}

func NewStore() *Store {
	return &Store{
		Origins: make(map[string]*Origin),
	}
}

func (store *Store) GetOrigin(id string) (*Origin, bool) {
	store.Lock()
	defer store.Unlock()

	origin, ok := store.Origins[id]
	return origin, ok
}

func (store *Store) AssertOrigin(id string) *Origin {
	store.Lock()
	defer store.Unlock()

	origin, ok := store.Origins[id]
	if !ok {
		origin = &Origin{
			Id:       id,
			Services: make(map[string]*Service),
		}
		store.Origins[id] = origin
	}

	return origin
}

func (store *Store) GetServices() []*Service {
	store.Lock()
	defer store.Unlock()

	services := make([]*Service, 0)
	for _, origin := range store.Origins {
		origin.Lock()
		for _, service := range origin.Services {
			services = append(services, service)
		}
		origin.Unlock()
	}

	return services
}

func (origin *Origin) GetService(id string) (*Service, bool) {
	origin.Lock()
	defer origin.Unlock()

	service, ok := origin.Services[id]
	return service, ok
}

func (origin *Origin) AssertService(id string) *Service {
	origin.Lock()
	defer origin.Unlock()

	service, ok := origin.Services[id]
	if !ok {
		service = &Service{
			Id:       id,
			Monitors: make(map[string]*Monitor),
		}
		origin.Services[id] = service
	}

	return service
}

func (service *Service) GetMonitor(id string) (*Monitor, bool) {
	service.Lock()
	defer service.Unlock()

	monitor, ok := service.Monitors[id]
	return monitor, ok
}

func (service *Service) AssertMonitor(id string) *Monitor {
	service.Lock()
	defer service.Unlock()

	monitor, ok := service.Monitors[id]
	if !ok {
		monitor = &Monitor{
			Id:       id,
			Statuses: make(map[string]monitoring.Status),
		}
		service.Monitors[id] = monitor
	}

	return monitor
}

func (service *Service) Status() monitoring.Status {
	if len(service.Monitors) == 0 {
		return monitoring.StatusUnknown
	}

	// TODO: Merge concensus?
	return monitoring.StatusUnknown
}

func (monitor *Monitor) Status() *MonitorStatus {
	result := &MonitorStatus{
		Statuses: make(map[monitoring.Status]int),
	}

	for _, status := range monitor.Statuses {
		count, ok := result.Statuses[status]
		if !ok {
			count = 0
			result.Statuses[status] = 0
		}
		result.Statuses[status] = count + 1
	}

	return result
}

func (monitor *Monitor) SetStatusForNode(id string, status monitoring.Status) {
	monitor.Statuses[id] = status
}

func (monitorStatus *MonitorStatus) Occurances(status monitoring.Status) int {
	occurances, ok := monitorStatus.Statuses[status]
	if !ok {
		return 0
	}

	return occurances
}
