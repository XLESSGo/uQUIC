package quic

import (
	"crypto/rand"
	tls "github.com/ban6cat6/protean"
)

const (
	DefaultUDPDatagramMinSize = 1200
)

type QUICSpec struct {
	// InitialPacketSpec specifies the QUIC Initial Packet, which includes Initial
	// Packet Headers and Frames.
	InitialPacketSpec InitialPacketSpec

	// ClientHelloSpec specifies the TLS ClientHello to be sent in the first Initial
	// Packet. It is implemented by the protean library and a valid ClientHelloSpec
	// for QUIC MUST include (tls).QUICTransportParametersExtension.
	ClientHelloSpec *tls.ClientHelloSpec

	// PClientConfig is the protean-specific configuration for the client.
	// It includes keys and fingerprints for protean's fingerprint concealment.
	PClientConfig *tls.PConfig

	// UDPDatagramMinSize specifies the minimum size of the UDP Datagram (UDP payload).
	// If the UDP Datagram is smaller than this size, zeros will be padded to the end
	// of the UDP Datagram until this size is reached.
	UDPDatagramMinSize int
}

func (s *QUICSpec) UpdateConfig(config *Config) {
	s.InitialPacketSpec.UpdateConfig(config)
}
