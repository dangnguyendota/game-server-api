package api

import "time"

type Scheduler interface {
	Schedule(action string, data map[string]interface{}, d time.Duration) error
	Cancel(action string) error
	CheckExistedAction(action string) bool
	CancelIfExist(action string)
	CancelAll()
	Stop()
}