package services

type Service interface {
	Get(interface{}) interface{}
}
