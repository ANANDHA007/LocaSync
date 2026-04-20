package store

import (
	"sync"
	"time"
)

type Entry struct {
	Value     string
	Timestamp int64
	ClientID  string
	Deleted   bool
}

type InMemoryStore struct {
	mu        sync.RWMutex
	Data      map[string]Entry
	ChangeLog []ChangeLog
}

func (s *InMemoryStore) Set(key, value, clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := time.Now().UnixNano()

	entry := Entry{
		Value:     value,
		Timestamp: ts,
		ClientID:  clientID,
		Deleted:   false,
	}

	s.Data[key] = entry

	s.ChangeLog = append(s.ChangeLog, ChangeLog{
		Key:       key,
		Value:     value,
		Timestamp: ts,
		ClientID:  clientID,
		Deleted:   false,
	})
}

func (s *InMemoryStore) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entry, ok := s.Data[key]
	if !ok || entry.Deleted {
		return "", false
	}

	return entry.Value, true
}

func (s *InMemoryStore) Delete(key, clientID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := time.Now().UnixNano()

	delete(s.Data, key)

	s.ChangeLog = append(s.ChangeLog, ChangeLog{
		Key:       key,
		Timestamp: ts,
		ClientID:  clientID,
		Deleted:   true,
	})
}

func (s *InMemoryStore) ApplyChanges(changes []ChangeLog) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, incoming := range changes {
		existing, exists := s.Data[incoming.Key]

		if !exists || shouldApply(existing, incoming) {
			s.Data[incoming.Key] = Entry{
				Value:     incoming.Value,
				Timestamp: incoming.Timestamp,
				ClientID:  incoming.ClientID,
				Deleted:   incoming.Deleted,
			}

			s.ChangeLog = append(s.ChangeLog, incoming)
		}
	}
}

func shouldApply(existing Entry, incoming ChangeLog) bool {
	if incoming.Timestamp > existing.Timestamp {
		return true
	}

	if incoming.Timestamp == existing.Timestamp {
		return incoming.ClientID > existing.ClientID
	}

	return false
}

func (s *InMemoryStore) GetChangesSince(ts int64) []ChangeLog {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []ChangeLog
	for _, c := range s.ChangeLog {
		if c.Timestamp > ts {
			result = append(result, c)
		}
	}
	return result
}

func (s *InMemoryStore) GetAllChanges() []ChangeLog {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return append([]ChangeLog{}, s.ChangeLog...)
}
