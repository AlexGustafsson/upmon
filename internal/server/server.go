package server

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/api"
	"github.com/AlexGustafsson/upmon/internal/clustering"
	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/AlexGustafsson/upmon/monitoring"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Server is an API server
type Server struct {
	config  *configuration.Configuration
	cluster *clustering.Cluster
}

// NewServer creates a new server
func NewServer(config *configuration.Configuration, cluster *clustering.Cluster) *Server {
	return &Server{
		config:  config,
		cluster: cluster,
	}
}

// Start starts the server on the given address and port
func (server *Server) Start(bind string) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/api/v1/origins", func(c *fiber.Ctx) error {
		origins := make([]api.Origin, 0)

		// TODO: Accept ghost reads? With this logic, an attacker may continually request this endpoint
		// and the server will be unable to make updates as the store is locked?
		server.cluster.Store.Lock()
		for _, origin := range server.cluster.Store.Origins {
			origins = append(origins, api.Origin{
				Id: origin.Id,
			})
		}
		server.cluster.Store.Unlock()

		response := api.NewOrigins(origins)

		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/origins/:originId", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")

			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Send(json)
		}

		response := api.Origin{
			Id: origin.Id,
		}

		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/services", func(c *fiber.Ctx) error {
		services := make([]api.Service, 0)

		for _, service := range server.cluster.Store.GetServices() {
			serviceResult := api.Service{
				Id: service.Id,
				Status: api.ServiceStatus{
					Status: service.Status().String(),
				},
			}

			services = append(services, serviceResult)
		}

		response := api.NewServices(services)

		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/origins/:originId/services", func(c *fiber.Ctx) error {
		services := make([]api.Service, 0)

		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")

			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Send(json)
		}

		origin.Lock()
		for _, service := range origin.Services {
			serviceResult := api.Service{
				Id: service.Id,
			}

			services = append(services, serviceResult)
		}
		origin.Unlock()

		response := api.NewServices(services)

		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/origins/:originId/services/:serviceId", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service, ok := origin.GetService(c.Params("serviceId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		response := api.Service{
			Id: service.Id,
		}
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(200).Send(json)
	})

	app.Get("/api/v1/origins/:originId/services/:serviceId/status", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service, ok := origin.GetService(c.Params("serviceId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		response := api.ServiceStatus{
			Status: service.Status().String(),
		}
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(200).Send(json)
	})

	app.Get("/api/v1/origins/:originId/services/:serviceId/monitors", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service, ok := origin.GetService(c.Params("serviceId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service.Lock()
		monitors := make([]api.Monitor, 0)
		for _, monitor := range service.Monitors {
			monitors = append(monitors, api.Monitor{
				Id: monitor.Id,
			})
		}
		service.Unlock()

		response := api.NewMonitors(monitors)
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(200).Send(json)
	})

	app.Get("/api/v1/origins/:originId/services/:serviceId/monitors/:monitorId", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service, ok := origin.GetService(c.Params("serviceId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		monitor, ok := service.GetMonitor(c.Params("monitorId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		response := api.Monitor{
			Id: monitor.Id,
		}
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(200).Send(json)
	})

	app.Get("/api/v1/origins/:originId/services/:serviceId/monitors/:monitorId/status", func(c *fiber.Ctx) error {
		origin, ok := server.cluster.Store.GetOrigin(c.Params("originId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		service, ok := origin.GetService(c.Params("serviceId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		monitor, ok := service.GetMonitor(c.Params("monitorId"))
		if !ok {
			response := api.NewErrorResponse("Not found")
			json, err := response.MarshalJSON()
			if err != nil {
				return err
			}

			return c.Status(404).Send(json)
		}

		status := monitor.Status()

		response := api.MonitorStatus{
			Up:                int32(status.Occurances(monitoring.StatusUp)),
			Down:              int32(status.Occurances(monitoring.StatusDown)),
			TransitioningUp:   int32(status.Occurances(monitoring.StatusTransitioningUp)),
			TransitioningDown: int32(status.Occurances(monitoring.StatusTransitioningDown)),
		}
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(200).Send(json)
	})

	app.Get("/api/v1/peers", func(c *fiber.Ctx) error {
		peers := make([]api.Peer, len(server.cluster.Memberlist.Members()))

		fmt.Printf("Got %d members vs %d members", server.cluster.Memberlist.NumMembers(), len(server.cluster.Memberlist.Members()))
		for i, member := range server.cluster.Memberlist.Members() {
			peer := api.Peer{
				Name:   member.Name,
				Bind:   member.FullAddress().Addr,
				Status: fmt.Sprintf("%d", member.State),
			}
			peers[i] = peer
		}

		response := api.NewPeers(peers)
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/peers/:peerId", func(c *fiber.Ctx) error {
		response := api.NewErrorResponse("Not found")
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(404).Send(json)
	})

	log.Infof("starting API server on %s", bind)
	return app.Listen(bind)
}
