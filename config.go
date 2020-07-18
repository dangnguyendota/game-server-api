package gs_interface

import "github.com/gocql/gocql"

type ConfigAll struct {
	Pusher   *PusherConfig
	Database *DatabaseConfig
}

type PusherConfig struct {
	RestApiKey string
	AppID      string
	ApiUrl     string
	Workers    int
	PoolQueue  int
}

type DatabaseConfig struct {
	Hosts       []string
	KeySpace    string
	Consistency gocql.Consistency
}
