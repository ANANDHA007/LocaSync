package store

type Store interface {
	Set(key, value, clientId string)
	Get(key string) (string, bool)
	Delete(key, clientId string)
}
