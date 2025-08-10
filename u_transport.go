package quic

import (
	"context"
	"errors"
	"net"

	"github.com/XLESSGo/uquic/internal/protocol"
	"github.com/XLESSGo/uquic/internal/utils"
	"github.com/XLESSGo/uquic/logging"
	tls "github.com/ban6cat6/protean"
)

type UTransport struct {
	*Transport

	QUICSpec *QUICSpec // [UQUIC] using ptr to avoid copying
}

// Dial dials a new connection to a remote host (not using 0-RTT).
func (t *UTransport) Dial(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *Config) (Connection, error) {
	return t.dial(ctx, addr, "", tlsConf, conf, false)
}

// DialEarly dials a new connection, attempting to use 0-RTT if possible.
func (t *UTransport) DialEarly(ctx context.Context, addr net.Addr, tlsConf *tls.Config, conf *Config) (EarlyConnection, error) {
	return t.dial(ctx, addr, "", tlsConf, conf, true)
}

func (t *UTransport) dial(ctx context.Context, addr net.Addr, host string, tlsConf *tls.Config, conf *Config, use0RTT bool) (EarlyConnection, error) {
	if err := t.init(t.isSingleUse); err != nil {
		return nil, err
	}
	if err := validateConfig(conf); err != nil {
		return nil, err
	}
	conf = populateConfig(conf)

	var proteanConfig *tls.PConfig
	if t.QUICSpec != nil {
		proteanConfig = t.QUICSpec.PClientConfig
	}

	return t.doDial(ctx, addr, host, tlsConf, conf, proteanConfig, use0RTT)
}

func (t *UTransport) doDial(ctx context.Context, addr net.Addr, host string, tlsConf *tls.Config, conf *Config, proteanConfig *tls.PConfig, use0RTT bool) (EarlyConnection, error) {
	if host == "" {
		if tlsConf != nil {
			host = tlsConf.ServerName
		}
		if host == "" {
			return nil, errors.New("quic: 'tls.Config.ServerName' must be set")
		}
	}
	conf = set
	conf.logDefault()

	if t.Tracer != nil {
		conf.Tracer = t.Tracer
	}
	if t.ConnIDGenerator != nil {
		conf.ConnectionIDGenerator = t.ConnIDGenerator
	}
	if t.serverConn != nil {
		return nil, errors.New("uquic: a Transport cannot be used for dialing and listening at the same time")
	}

	conn, err := t.newClientConnection(ctx, addr, host, tlsConf, conf, proteanConfig, use0RTT)
	if err != nil {
		return nil, err
	}
	conn.handshakeEvent()
	return conn, nil
}
