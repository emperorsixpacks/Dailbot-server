package entities

import (
	"github.com/emperorsixpacks/dailbot/src/internal/infrastructure/models"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
)

func NewWebHook(baseID, webS string, tableIDs []string) *AirtableWebhook {
	return &AirtableWebhook{
		baseID:        baseID,
		webhookSecret: webS,
		tableIds:      tableIDs,
	}
}

type AirtableWebhook struct {
	webhookID     string
	baseID        string
	personalToken string
	cursor        int
	tableIds      []string
	webhookSecret string
}

func (a *AirtableWebhook) SetPToken(token, secret string) error {
	encryptedToken, err := utils.EncryptString(token, []byte(secret))
	if err != nil {
		return err
	}
	a.personalToken = encryptedToken
	return nil
}

func (a *AirtableWebhook) SetWebhookID(newID ...string) *AirtableWebhook {
	if len(newID) == 0 {
		id, err := utils.GenerateRandomString(16)
		if err != nil {
			panic(err)
		}
		newID[0] = id
	}
	a.webhookID = newID[0]
	return a
}

func (a *AirtableWebhook) IncrementCursor() {
	a.cursor += 1
}

func (a AirtableWebhook) Model() *models.WebhookModel {
	return &models.WebhookModel{
		WebHookID:     a.webhookID,
		BaseID:        a.baseID,
		PersonalToken: a.personalToken,
		Cursor:        a.cursor,
		TablesIDs:     a.tableIds,
		WebHookSecret: a.webhookSecret,
	}
}
