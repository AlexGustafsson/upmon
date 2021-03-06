package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/AlexGustafsson/upmon/core"
	"math/big"
	"net"
	"os"
	"time"
)

// SEE: https://golang.org/src/crypto/tls/generate_cert.go
func createSelfSignedCertificate(hostnames ...string) (*ecdsa.PrivateKey, *x509.Certificate, []byte, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		core.LogError("Unable to create serial number")
		return nil, nil, nil, err
	}

	now := time.Now()
	certificate := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Upmon"},
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Append IPs and hostnames to the certificate
	for _, hostname := range hostnames {
		ip := net.ParseIP(hostname)
		if ip != nil {
			certificate.IPAddresses = append(certificate.IPAddresses, ip)
		} else {
			certificate.DNSNames = append(certificate.DNSNames, hostname)
		}
	}

	privateKey, err := generateKeys()
	if err != nil {
		return nil, nil, nil, err
	}

	certificateBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificate, privateKey.Public().(ecdsa.PublicKey), privateKey)
	if err != nil {
		core.LogError("Unable to create certificate authority")
		return nil, nil, nil, err
	}

	return privateKey, certificate, certificateBytes, nil
}

func writeCertificate(certificateBytes []byte, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, &pem.Block{Type: "CERTIFICATE", Bytes: certificateBytes})
	if err != nil {
		core.LogError("Unable to encode certificate")
		return err
	}

	return nil
}

func writeKey(privateKey *ecdsa.PrivateKey, path string) error {
	file, err := os.Create(path)
	if err != nil {
		core.LogError("Unable to create file '%s'", path)
		return err
	}
	defer file.Close()

	keyBytes, err := x509.MarshalPKCS8PrivateKey(*privateKey)
	if err != nil {
		core.LogError("Unable to marshal key")
		return err
	}

	err = pem.Encode(file, &pem.Block{Type: "PRIVATE KEY", Bytes: keyBytes})
	if err != nil {
		core.LogError("Unable to encode key")
		return err
	}

	return nil
}

func generateKeys() (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		core.LogError("Unable to generate keys")
		return nil, err
	}

	return privateKey, nil
}
