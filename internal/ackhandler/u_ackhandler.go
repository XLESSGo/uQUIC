package ackhandler

import (
	"github.com/XLESSGo/uquic/internal/protocol"
	"github.com/XLESSGo/uquic/internal/utils"
	"github.com/XLESSGo/uquic/logging"
)

// [UQUIC]
func NewUAckHandler(
	initialPacketNumber protocol.PacketNumber,
	initialMaxDatagramSize protocol.ByteCount,
	rttStats *utils.RTTStats,
	clientAddressValidated bool,
	enableECN bool,
	pers protocol.Perspective,
	tracer *logging.ConnectionTracer,
	logger utils.Logger,
) (SentPacketHandler, ReceivedPacketHandler) {
	sph := newSentPacketHandler(initialPacketNumber, initialMaxDatagramSize, rttStats, clientAddressValidated, enableECN, pers, tracer, logger)
	return &uSentPacketHandler{
		sentPacketHandler: sph,
	}, newReceivedPacketHandler(sph, logger)
}
