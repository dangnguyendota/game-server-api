package gsi

import "log"

type ServerMode uint8

const (
	AuthoritativeMultiPlayerServer = 0
	RelayedMultiPlayerServer       = 1
)

func (s *ServerMode) String(mode ServerMode) string {
	if mode == RelayedMultiPlayerServer {
		return "Relayed"
	}
	return "Authoritative"
}

func Parse(mode string) ServerMode {
	if mode == "authoritative" {
		return AuthoritativeMultiPlayerServer
	} else if mode == "relayed" {
		return RelayedMultiPlayerServer
	} else {
		log.Fatal("can not parse game mode of", mode)
	}
	return AuthoritativeMultiPlayerServer
}
