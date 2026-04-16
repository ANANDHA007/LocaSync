package store

type Store interface {
	Set(key, value, clientId string)
	Get(key string) string
	Delete(key, clientId string)
}
