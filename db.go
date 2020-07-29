package gsi

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gocql/gocql"
)

type Database interface {
	GetUser(username, password string) (User, error)
	GetUserFacebook(facebookToken string) (User, error)
	GetUserGoogle(googleToken string) (User, error)
	GetAdmin(username, password string) (User, error)
	MemcachedClient() *memcache.Client
	CassandraSession() *gocql.Session
	Close()
}
