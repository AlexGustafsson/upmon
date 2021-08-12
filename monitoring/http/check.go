package http

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/AlexGustafsson/upmon/monitoring/core"
	log "github.com/sirupsen/logrus"
)

type Check struct {
	config *MonitorConfiguration
}

// NewCheck creates a new check for a service
func NewCheck(options map[string]interface{}) (*Check, error) {
	config, err := ParseConfiguration(options)
	if err != nil {
		return nil, err
	}

	return &Check{
		config: config,
	}, nil
}

func (check *Check) Name() string {
	return "http"
}

func (check *Check) Description() string {
	return "HTTP client"
}

func (check *Check) Version() string {
	return "0.1.0"
}

func (check *Check) Perform() (core.Status, error) {
	switch check.config.Method {
	case "GET":
		return check.PerformGet()
	default:
		return core.StatusUnknown, fmt.Errorf("unsupported HTTP method")
	}
}

func (check *Check) PerformGet() (core.Status, error) {
	redirects := 0

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if check.config.FollowRedirects {
				if redirects > check.config.MaximumRedirects {
					return fmt.Errorf("maximum number of redirects exceeded")
				}
				redirects++
				return nil
			} else {
				return http.ErrUseLastResponse
			}
		},
	}

	response, err := client.Get(check.config.URL)
	if err != nil {
		return core.StatusUnknown, err
	}

	if check.config.Expect.Status != 0 {
		if response.StatusCode != check.config.Expect.Status {
			return core.StatusDown, nil
		}
	}

	return core.StatusUp, nil
}

func (check *Check) Config() core.MonitorConfiguration {
	return check.config
}

func (check *Check) Watch(update chan<- *core.ServiceStatus, stop <-chan bool, wg *sync.WaitGroup) error {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				wg.Done()
				log.WithFields(log.Fields{"check": "http"}).Debugf("exiting")
				return
			default:
				status, err := check.Perform()
				update <- &core.ServiceStatus{
					Err:    err,
					Status: status,
					Check:  check,
				}
			}
			time.Sleep(check.config.Interval)
		}
	}()

	return nil
}
