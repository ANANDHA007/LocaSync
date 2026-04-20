package store

type ChangeLog struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
	ClientID  string `json:"client_id"`
	Deleted   bool   `json:"deleted"`
}

type SyncRequest struct {
	ClientID   string      `json:"client_id"`
	LastSyncTs int64       `json:"last_sync_ts"`
	Changes    []ChangeLog `json:"changes"`
}
