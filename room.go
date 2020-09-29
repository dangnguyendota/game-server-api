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

type Metadata interface {
	String() string
	Byte() (byte, error)
	Bytes() []byte
	Int64() (int64, error)
	Int32() (int32, error)
	Uint64() (uint64, error)
	Uint32() (uint32, error)
	Float64() (float64, error)
	Float32() (float32, error)
	Json(interface{}) error
}

type Room interface {
	// returns unique room id
	ID() uuid.UUID
	// returns number of game loops per second
	TickRate() int64
	// returns unique game id
	Game() string
	// name of matchmaker node
	Node() string
	// returns maximum of players each room
	MaxPlayers() int
	// returns the minimum of players each room
	MinPlayers() int
	// returns the maximum of viewer each room
	MaxViewers() int
	// returns server mode (Authoritative MultiPlayer Server or Relayed MultiPlayer Server)
	Mode() Mode
	// logger to write log
	Logger() *zap.Logger
	// returns fixed, added metadata
	Metadata(name string) Metadata
	// returns current game state
	State() interface{}
	// returns current players in this room
	Players() []*User
	// returns current viewers in this room
	Viewers() []*User
	// returns time this room is created
	CreateTime() int64
	// kick user
	Kick(userId uuid.UUID) error
	// destroy this room, kick all players and viewers
	Destroy()
	// pause game, wont handle any incoming packet
	Pause()
	// unpause game, allow client to send packet into room
	Unpause()
	// returns game is paused
	Paused() bool
	// returns scheduler to schedule actions
	Scheduler() Scheduler
	// returns pusher, using one signal to push notification to clients
	Pusher() NotificationPusher
	// returns overwritten database
	DB() Database
	// returns dispatcher to send messages to clients
	Dispatcher() Dispatcher
}
