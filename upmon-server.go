package main

import (
  "github.com/AlexGustafsson/upmon/rpc"
  "github.com/AlexGustafsson/upmon/core"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/credentials"
	"context"
	"fmt"
  "crypto/sha1"
	"encoding/base64"
)

type upmonServer struct {
  rpc.UnimplementedUpmonServer
}

func (server *upmonServer) SendServicePing(ctx context.Context, ping *rpc.ServicePing) (*rpc.Empty, error) {
  peer, ok := peer.FromContext(ctx)
  if ok {
    tlsInfo := peer.AuthInfo.(credentials.TLSInfo)

    if len(tlsInfo.State.PeerCertificates) == 0 {
      core.LogWarning("Peer sent no certificates, closing connection")
      return nil, fmt.Errorf("No peer certificates")
    } else if len(tlsInfo.State.PeerCertificates) != 1 {
      core.LogWarning("Peer has multiple certificates, closing connection")
      return nil, fmt.Errorf("Peer has multiple certificates")
    }

    certificate := tlsInfo.State.PeerCertificates[0]

  	shasum := sha1.Sum(certificate.Raw)
  	fingerprint := base64.StdEncoding.EncodeToString(shasum[:])
  	core.LogDebug("Peer has the fingerprint fingerprint: %v", fingerprint)
  } else {
    core.LogError("Failed to get info about peer")
    return nil, fmt.Errorf("Unable to get peer info")
  }

  core.LogDebug("Got request to send service ping")

  return &rpc.Empty{}, nil
}
