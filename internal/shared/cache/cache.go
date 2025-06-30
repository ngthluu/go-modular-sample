package cache

type Cache interface {
	Set(key string, value string, ttl int) error
	Update(key string, value string) error
	Get(key string) (string, error)
	Remove(key string) error
	Close()
}
