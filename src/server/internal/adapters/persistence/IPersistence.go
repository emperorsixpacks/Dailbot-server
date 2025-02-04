package persistence

type Persistence interface {
	Connect() interface{}
}
