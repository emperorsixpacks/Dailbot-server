package repositories

type Repository[T any] interface {
	Create(entity *T) T 
}
