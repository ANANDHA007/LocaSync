package core

import (
	"github.com/ANANDHA007/LocaSync/config"
	"github.com/ANANDHA007/LocaSync/store"
)

type LocaSync struct {
	Store store.Store
}

func New(config config.Config) (LocaSync, error) {
	store, err := store.NewStore(config.Store)
	if err != nil {
		return LocaSync{}, err
	}
	return LocaSync{
		Store: store,
	}, nil
}

func (l *LocaSync) Set(key, value, clientId string) {
	l.Store.Set(key, value, clientId)
}
func (l *LocaSync) Get(key string) string {
	value, _ := l.Store.Get(key)
	return value
}
func (l *LocaSync) Delete(key, clientId string) {
	l.Store.Delete(key, clientId)
}
