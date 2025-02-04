package entities

import (
	"github.com/emperorsixpacks/dailbot/internal/infrastructure/models"
	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/google/uuid"
)

func NewWebHook(webhId, ptn, baseID, webS string, tableIDs []string) *AirtableWebhook {
	return &AirtableWebhook{
		webhookID:     webhId,
		baseID:        baseID,
		webhookSecret: webS,
		personalToken: ptn,
		tableIds:      tableIDs,
	}
}

type AirtableWebhook struct {
	webhookID     string
	baseID        string
	personalToken string // TODO look at storing this as a hash with a saltApp
	cursor        int
	tableIds      []string
	webhookSecret string
}

func (a *AirtableWebhook) NewPToken(token, secret string) error {
	encryptedToken, err := utils.EncryptString(token, []byte(secret))
	if err != nil {
		return err
	}
	a.personalToken = encryptedToken
	return nil
}

func (a *AirtableWebhook) IncrementCursor() {
	a.cursor += 1
}

type AirtableWebHookRepository interface {
	CreateWebhook(*models.WebhookModel)
	DeletetWebhhook(uuid.UUID) error
	GetWebhook(uuid.UUID) AirtableWebhook
}
