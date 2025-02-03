package entities

import (
	"github.com/emperorsixpacks/dailbot/pkg/utils"
)

type AirtableWebhook struct {
	webhookID     string
	baseID        string
	personalToken string // TODO look at storing this as a hash with a saltApp
	cursor        int
	tableIds      []string
	webhookSecret string
}

func (a *AirtableWebhook) SetPToken(token, secret string) error {
	hashToken, err := utils.EncryptString(token, []byte(secret))
	if err != nil {
		return err
	}
	a.personalToken = hashToken
	return nil
}

// TODO encrypt personal token in db
