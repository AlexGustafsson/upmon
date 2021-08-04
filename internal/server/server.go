package server

import (
	"fmt"

	"github.com/AlexGustafsson/upmon/api"
	"github.com/AlexGustafsson/upmon/internal/clustering"
	"github.com/AlexGustafsson/upmon/internal/configuration"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	config  *configuration.Configuration
	cluster *clustering.Cluster
}

func NewServer(config *configuration.Configuration, cluster *clustering.Cluster) *Server {
	return &Server{
		config:  config,
		cluster: cluster,
	}
}

func (server *Server) Start(bind string) error {
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

	log.Infof("starting server on %s", bind)
	return app.Listen(bind)
}
