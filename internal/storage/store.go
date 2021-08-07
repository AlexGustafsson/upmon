package storage

import (
	"sync"

	"github.com/AlexGustafsson/upmon/monitoring"
)

// Store is a central state storage for services and their statuses
type Store struct {
	sync.Mutex
	// Origins are all of the origins available in the store, by their id
	Origins map[string]*Origin
}

// Origin is an origin node from which services originate
type Origin struct {
	sync.Mutex
	Id string
	// Services are the services requested for monitoring by the origin, by their id
	Services map[string]*Service
}

// Service is a monitored service
type Service struct {
	sync.Mutex
	Id string
	// Monitors are the monitors configured for the service, by their id
	Monitors map[string]*Monitor
}

// Monitor is a monitor for a service
type Monitor struct {
	sync.Mutex
	Id string
	// Statuses is the current status, by the nodes' ids
	Statuses map[string]monitoring.Status
}

// ServiceStatus is the status of a service, assessed by using all configured monitors
type ServiceStatus struct {
}

// MonitorStatus is the status of a service, assessed by using all results from all cluster nodes
type MonitorStatus struct {
	// Statuses are the occurances for each observed status for the monitor
	Statuses map[monitoring.Status]int
}

// NewStore creates a new store
func NewStore() *Store {
	return &Store{
		Origins: make(map[string]*Origin),
	}
}

// GetOrigin retrieves an origin by its id
func (store *Store) GetOrigin(id string) (*Origin, bool) {
	store.Lock()
	defer store.Unlock()

	origin, ok := store.Origins[id]
	return origin, ok
}

// AssertOrigin retrieves an origin and creates it if it does not already exist
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

// GetServices retrieves all configured services
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

// GetService retrieves a service by its id
func (origin *Origin) GetService(id string) (*Service, bool) {
	origin.Lock()
	defer origin.Unlock()

	service, ok := origin.Services[id]
	return service, ok
}

// AssertService retrieves a service by its id and creates it if it does not already exist
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

// GetMonitor retrieves a monitor by its id
func (service *Service) GetMonitor(id string) (*Monitor, bool) {
	service.Lock()
	defer service.Unlock()

	monitor, ok := service.Monitors[id]
	return monitor, ok
}

// AssertMonitor retrieves a monitor by its id and creates it if it does not already exist
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

// Status retrieves the current status of the service, taking all monitors into account
func (service *Service) Status() monitoring.Status {
	if len(service.Monitors) == 0 {
		return monitoring.StatusUnknown
	}

	// TODO: Merge concensus?
	return monitoring.StatusUnknown
}

// Status retrieves the current status of the monitored service
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

// SetStatusForNode sets the status for a node by its id
func (monitor *Monitor) SetStatusForNode(id string, status monitoring.Status) {
	monitor.Statuses[id] = status
}

// Occurances retrieves the number of observed statuses of the given type
func (monitorStatus *MonitorStatus) Occurances(status monitoring.Status) int {
	occurances, ok := monitorStatus.Statuses[status]
	if !ok {
		return 0
	}

	return occurances
}
