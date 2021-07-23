package server

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/api"
	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/memberlist"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config *configuration.Configuration
	list   *memberlist.Memberlist
}

func NewServer(config *configuration.Configuration, list *memberlist.Memberlist) *Server {
	return &Server{
		config: config,
		list:   list,
	}
}

func (server *Server) Start(address string, port uint16) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/api/v1/services", func(c *fiber.Ctx) error {
		services := make([]api.Service, 0)
		response := api.NewServices(services)
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/services/:serviceId", func(c *fiber.Ctx) error {
		response := api.NewErrorResponse("Not found")
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Status(404).Send(json)
	})

	app.Get("/api/v1/services/:serviceId/status", func(c *fiber.Ctx) error {
		response := api.NewServiceStatus("unknown")
		json, err := response.MarshalJSON()
		if err != nil {
			return err
		}

		return c.Send(json)
	})

	app.Get("/api/v1/peers", func(c *fiber.Ctx) error {
		peers := make([]api.Peer, len(server.list.Members()))

		for _, member := range server.list.Members() {
			peer := api.Peer{
				Name:    member.Name,
				Address: member.Address(),
				Port:    float32(member.Port),
				Status:  fmt.Sprintf("%d", member.State),
			}
			peers = append(peers, peer)
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

	log.Infof("starting server on %s:%d", address, port)
	return app.Listen(fmt.Sprintf("%s:%d", address, port))
}
