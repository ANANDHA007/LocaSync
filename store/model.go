package store

import "time"

type ChangeLog struct {
	Key       string
	Value     string
	TimeStamp time.Time
	ClientID  string
	Deleted   bool
}
