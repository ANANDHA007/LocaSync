package store

import "time"

type InMemoryStore struct {
	Data      map[string]string
	ChangeLog []ChangeLog
}

func (inmemoryStore *InMemoryStore) Set(key string, value string, clientId string) {
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
	return inmemoryStore.Data[key]
}

func (inmemoryStore *InMemoryStore) Delete(key, clientId string) {
	delete(inmemoryStore.Data, key)
	inmemoryStore.ChangeLog = append(inmemoryStore.ChangeLog, ChangeLog{
		Key:       key,
		ClientID:  clientId,
		Deleted:   true,
		TimeStamp: time.Now(),
	})
}
