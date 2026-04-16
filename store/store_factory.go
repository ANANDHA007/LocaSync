package store

import "errors"

func NewStore(store string) (Store, error) {
	switch store {
	case "InMemoryStore":
		return &InMemoryStore{
			Data:      make(map[string]string),
			ChangeLog: make([]ChangeLog, 0),
		}, nil
	default:
		return nil, errors.New("The provided storage type is not supported ")
	}
}
