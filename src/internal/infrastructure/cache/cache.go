package cache

type Cache interface {
	Get(string)
	Set(string, interface{})
}
