package api

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Dispatcher interface {
	// send to separate user
	Send(sid uuid.UUID, data []byte)
	// send all viewers and players
	SendAll(data []byte)
	// send all players
	SendAllPlayers(data []byte)
	// send all viewers
	SendAllViewers(data []byte)
	// send error to client
	SendError(sid uuid.UUID, code uint32, message string)
}

type Room interface {
	// unique room id
	ID() uuid.UUID
	// number of game loop per second
	TickRate() int64
	// unique game id
	Game() string
	// maximum of players per room
	MaxPlayers() int
	// minimum of players per room
	MinPlayers() int
	// maximum of viewer per room
	MaxViewers() int
	// server mode (Authoritative MultiPlayer Server or Relayed MultiPlayer Server)
	Mode() ServerMode
	// logger to write log
	Logger() *zap.Logger
	// fixed metadata when init room
	Metadata() map[string]string
	// current game state (store game state in memory for every game loop)
	State() interface{}
	// current players in this room
	Players() []*User
	// current viewers in this room
	Viewers() []*User
	// time this room is created
	CreateTime() int64
	// destroy this room, kick all players and viewers
	Destroy()
	// scheduler to schedule actions
	Scheduler() Scheduler
	// use one signal to push notification to clients
	Pusher() NotificationPusher
	// get database
	DB() Database
	// get dispatcher to send message to clients
	Dispatcher() Dispatcher
}
