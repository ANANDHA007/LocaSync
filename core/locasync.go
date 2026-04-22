package core

import (
	"github.com/ANANDHA007/LocaSync/config"
	"github.com/ANANDHA007/LocaSync/store"
	"github.com/ANANDHA007/LocaSync/utils"
)

type LocaSync struct {
	Store    store.Store
	ClientID string
}

func New(config config.Config) (LocaSync, error) {
	store, err := store.NewStore(config.Store)
	if err != nil {
		return LocaSync{}, err
	}
	clientId := config.ClientID
	if config.ClientID == "" {
		clientId = utils.GenerateClientID()
	}
	return LocaSync{
		Store:    store,
		ClientID: clientId,
	}, nil
}

func (l *LocaSync) Set(key, value string) {
	l.Store.Set(key, value, l.ClientID)
}
func (l *LocaSync) Get(key string) string {
	value, _ := l.Store.Get(key)
	return value
}
func (l *LocaSync) Delete(key string) {
	l.Store.Delete(key, l.ClientID)
}
