//go:build gomock || generate

package quic

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_send_conn_test.go github.com/refraction-networking/uquic SendConn"
type SendConn = sendConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_raw_conn_test.go github.com/refraction-networking/uquic RawConn"
type RawConn = rawConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_sender_test.go github.com/refraction-networking/uquic Sender"
type Sender = sender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_stream_internal_test.go github.com/refraction-networking/uquic StreamI"
type StreamI = streamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_receive_stream_internal_test.go github.com/refraction-networking/uquic ReceiveStreamI"
type ReceiveStreamI = receiveStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_send_stream_internal_test.go github.com/refraction-networking/uquic SendStreamI"
type SendStreamI = sendStreamI

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_stream_sender_test.go github.com/refraction-networking/uquic StreamSender"
type StreamSender = streamSender

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_stream_control_frame_getter_test.go github.com/refraction-networking/uquic StreamControlFrameGetter"
type StreamControlFrameGetter = streamControlFrameGetter

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_frame_source_test.go github.com/refraction-networking/uquic FrameSource"
type FrameSource = frameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_ack_frame_source_test.go github.com/refraction-networking/uquic AckFrameSource"
type AckFrameSource = ackFrameSource

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_stream_manager_test.go github.com/refraction-networking/uquic StreamManager"
type StreamManager = streamManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_sealing_manager_test.go github.com/refraction-networking/uquic SealingManager"
type SealingManager = sealingManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_unpacker_test.go github.com/refraction-networking/uquic Unpacker"
type Unpacker = unpacker

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_packer_test.go github.com/refraction-networking/uquic Packer"
type Packer = packer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_mtu_discoverer_test.go github.com/refraction-networking/uquic MTUDiscoverer"
type MTUDiscoverer = mtuDiscoverer

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_conn_runner_test.go github.com/refraction-networking/uquic ConnRunner"
type ConnRunner = connRunner

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_quic_conn_test.go github.com/refraction-networking/uquic QUICConn"
type QUICConn = quicConn

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_packet_handler_test.go github.com/refraction-networking/uquic PacketHandler"
type PacketHandler = packetHandler

//go:generate sh -c "go run go.uber.org/mock/mockgen -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_packet_handler_manager_test.go github.com/refraction-networking/uquic PacketHandlerManager"

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package quic -self_package github.com/refraction-networking/uquic -destination mock_packet_handler_manager_test.go github.com/refraction-networking/uquic PacketHandlerManager"
type PacketHandlerManager = packetHandlerManager

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package quic -self_package github.com/refraction-networking/uquic -self_package github.com/refraction-networking/uquic -destination mock_packetconn_test.go net PacketConn"
