package main

import (
	"context"
	"github.com/AlexGustafsson/upmon/core"
	"github.com/AlexGustafsson/upmon/rpc"
)

type upmonServer struct {
	rpc.UnimplementedUpmonServer
}

func (server *upmonServer) SendServicePing(ctx context.Context, ping *rpc.ServicePing) (*rpc.Empty, error) {
	core.LogDebug("Got request to send service ping")

	return &rpc.Empty{}, nil
}
