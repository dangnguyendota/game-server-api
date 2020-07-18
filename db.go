package gsi

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gocql/gocql"
)

type Database interface {
	Config() *DatabaseConfig
	GetUser(username, password string) (User, error)
	GetUserFacebook(facebookToken string) (User, error)
	GetUserGoogle(googleToken string) (User, error)
	GetAdmin(username, password string) (User, error)
	GetMemcachedClient() *memcache.Client
	GetCassandraSession() *gocql.Session
	Close()
}