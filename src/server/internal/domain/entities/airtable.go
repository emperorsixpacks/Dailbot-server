package entities

type artable struct {
	user          user
	webhookID     string
	personalToken string // TODO look at storing this as a hash with a saltApp
}
