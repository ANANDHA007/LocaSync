package config

import "time"

type SyncMode int

const (
	SyncManual SyncMode = iota
	SyncAuto
	SyncHybrid
)

type Config struct {
	StoragePath  string
	Store        string
	SyncMode     SyncMode
	SyncInterval time.Duration
}
