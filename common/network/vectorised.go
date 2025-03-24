package network

import (
	"github.com/konglong147/sing/common/buf"
	M "github.com/konglong147/sing/common/metadata"
)

type VectorisedWriter interface {
	WriteVectorised(buffers []*buf.Buffer) error
}

type VectorisedPacketWriter interface {
	WriteVectorisedPacket(buffers []*buf.Buffer, destination M.Socksaddr) error
}
