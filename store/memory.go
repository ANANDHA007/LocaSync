package store

import (
	"sync"
	"time"
)

type InMemoryStore struct {
	mu        sync.RWMutex
	Data      map[string]string
	ChangeLog []ChangeLog
}

func (inmemoryStore *InMemoryStore) Set(key string, value string, clientId string) {
	inmemoryStore.mu.Lock()
	defer inmemoryStore.mu.Unlock()
	inmemoryStore.Data[key] = value
	inmemoryStore.ChangeLog = append(inmemoryStore.ChangeLog, ChangeLog{
		Key:       key,
		Value:     value,
		ClientID:  clientId,
		Deleted:   false,
		TimeStamp: time.Now(),
	})
}

func (inmemoryStore *InMemoryStore) Get(key string) string {
	inmemoryStore.mu.RLock()
	defer inmemoryStore.mu.Unlock()
	return inmemoryStore.Data[key]
}

func (inmemoryStore *InMemoryStore) Delete(key, clientId string) {
	inmemoryStore.mu.RLock()
	defer inmemoryStore.mu.Unlock()
	delete(inmemoryStore.Data, key)
	inmemoryStore.ChangeLog = append(inmemoryStore.ChangeLog, ChangeLog{
		Key:       key,
		ClientID:  clientId,
		Deleted:   true,
		TimeStamp: time.Now(),
	})
}
