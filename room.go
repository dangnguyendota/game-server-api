package gsi

import (
	"github.com/dangnguyendota/gs-interface/gs_proto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Room interface {
	ID() uuid.UUID
	Type() string
	Max() int
	Mode() ServerMode
	CreateTime() int64
	Logger() *zap.Logger
	MetaData() map[string]string
	CloseRoom()
	Players() []*User
	GetScheduler() Scheduler
	GetPusher() NotificationPusher
	GetDatabase() Database
	Send(sid uuid.UUID, packet *ip.Packet)
	SendAll(packet *ip.Packet)
}
