package gsi

import (
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
	Send(sid uuid.UUID, data []byte)                      // send to separate user
	SendAll(data []byte)                                  // send all viewers and players
	SendAllPlayers(data []byte)                           // send all players
	SendAllViewers(data []byte)                           // send all viewers
	SendError(sid uuid.UUID, code uint32, message string) // send error to client
}
