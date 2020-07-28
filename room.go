package gsi

import (
	"github.com/dangnguyendota/gs-interface/gs_proto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Room interface {
	ID() uuid.UUID
	Game() string
	MaxPlayers() int
	MinPlayers() int
	MaxViewers() int
	Mode() ServerMode
	Logger() *zap.Logger
	Metadata() map[string]string
	Players() []*User
	Viewers() []*User
	CreateTime() int64
	Destroy()
	GetScheduler() Scheduler
	GetPusher() NotificationPusher
	GetDatabase() Database
	Send(sid uuid.UUID, packet *ip.Packet)
	SendAll(packet *ip.Packet)
}
