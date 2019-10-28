package transport

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"google.golang.org/grpc/credentials"
	"github.com/AlexGustafsson/upmon/core"
	"net"
)

// SEE: https://github.com/grpc/grpc-go/blob/027cd627f8565bf731a6e1bdf54ce79ba09bda5c/credentials/credentials.go

type tlsTransport struct {
	credentials.TransportCredentials
	config *tls.Config
}

type tlsTransportError struct {
	text string
}

var allowedFingerprints []string

func (transportError *tlsTransportError) Error() string {
	return transportError.text
}

// See below:
// If the returned error is a wrapper error, implementations should make sure that
// the error implements Temporary() to have the correct retry behaviors.
func (transportError *tlsTransportError) Temporary() bool {
	return true
}

func newTransportError(text string) error {
	return &tlsTransportError{text}
}

// SetAllowedFingerprints sets the allowed fingerprints
func SetAllowedFingerprints(fingerprints []string) {
	allowedFingerprints = fingerprints
}

// ClientHandshake does the authentication handshake specified by the corresponding
// authentication protocol on rawConn for clients. It returns the authenticated
// connection and the corresponding auth information about the connection.
// Implementations must use the provided context to implement timely cancellation.
// gRPC will try to reconnect if the error returned is a temporary error
// (io.EOF, context.DeadlineExceeded or err.Temporary() == true).
// If the returned error is a wrapper error, implementations should make sure that
// the error implements Temporary() to have the correct retry behaviors.
//
// If the returned net.Conn is closed, it MUST close the net.Conn provided.
func (transport *tlsTransport) ClientHandshake(ctx context.Context, authority string, rawConnection net.Conn) (_ net.Conn, _ credentials.AuthInfo, err error) {
	config := cloneTLSConfig(transport.config)

	connection := tls.Client(rawConnection, config)
	errChannel := make(chan error, 1)
	go func() {
		errChannel <- connection.Handshake()
	}()
	select {
	case err := <-errChannel:
		if err != nil {
			return nil, nil, err
		}
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	}

	fingerprint, err := getFingerprint(connection)
	if err != nil {
		return nil, nil, err
	}

	for _, allowedFingerprint := range allowedFingerprints {
		if fingerprint == allowedFingerprint {
			core.LogDebug("Handshake succeeded with peer with fingerprint '%v'", fingerprint)
			return connection, credentials.TLSInfo{connection.ConnectionState()}, nil
		}
	}

	core.LogWarning("Handshake with peer was unsucessful - fingeprint '%v' is not allowed", fingerprint)
	return nil, nil, newTransportError(fmt.Sprintf("Handshake with peer '%v' failed, fingeprint '%v' is not allowed", connection.RemoteAddr(), fingerprint))
}

// ServerHandshake does the authentication handshake for servers. It returns
// the authenticated connection and the corresponding auth information about
// the connection.
//
// If the returned net.Conn is closed, it MUST close the net.Conn provided.
func (transport *tlsTransport) ServerHandshake(rawConnection net.Conn) (net.Conn, credentials.AuthInfo, error) {
	connection := tls.Server(rawConnection, transport.config)

	err := connection.Handshake()
	if err != nil {
		return nil, nil, newTransportError(fmt.Sprintf("Hanshake with peer '%v' failed", connection.RemoteAddr()))
	}

	fingerprint, err := getFingerprint(connection)
	if err != nil {
		return nil, nil, err
	}

	for _, allowedFingerprint := range allowedFingerprints {
		if fingerprint == allowedFingerprint {
			core.LogDebug("Handshake succeeded with peer with fingerprint '%v'", fingerprint)
			return connection, credentials.TLSInfo{connection.ConnectionState()}, nil
		}
	}

	core.LogWarning("Handshake with peer was unsucessful - fingeprint '%v' is not allowed", fingerprint)
	return nil, nil, newTransportError(fmt.Sprintf("Handshake with peer '%v' failed, fingeprint '%v' is not allowed", connection.RemoteAddr(), fingerprint))
}

func getFingerprint(connection *tls.Conn) (string, error) {
	state := connection.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		connection.Close()
		return "", newTransportError(fmt.Sprintf("Peer from %s sent no certificates, closing connection", connection.RemoteAddr()))
	} else if len(state.PeerCertificates) != 1 {
		connection.Close()
		return "", newTransportError(fmt.Sprintf("Peer from %s has multiple certificates, closing connection", connection.RemoteAddr()))
	}

	certificate := state.PeerCertificates[0]

	shasum := sha1.Sum(certificate.Raw)
	fingerprint := base64.StdEncoding.EncodeToString(shasum[:])

	return fingerprint, nil
}

// Info provides the ProtocolInfo of this TransportCredentials.
func (transport *tlsTransport) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{
		SecurityProtocol: "tls",
		SecurityVersion:  "1.3",
		ServerName:       transport.config.ServerName,
	}
}

// Clone makes a copy of this TransportCredentials.
func (transport *tlsTransport) Clone() credentials.TransportCredentials {
	return WithConfig(transport.config)
}

func cloneTLSConfig(config *tls.Config) *tls.Config {
	return config.Clone()
}

// WithConfig creates a Transport with a given TLS config
func WithConfig(config *tls.Config) credentials.TransportCredentials {
	transportCredentials := &tlsTransport{
		config: cloneTLSConfig(config),
	}
	return transportCredentials
}
