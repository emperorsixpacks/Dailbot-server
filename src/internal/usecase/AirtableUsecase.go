package usecase

import (
	"github.com/emperorsixpacks/dailbot/src/internal/domain/entities"
	"github.com/emperorsixpacks/dailbot/src/internal/infrastructure/repositories"
	"github.com/emperorsixpacks/dailbot/src/internal/infrastructure/schemas"
	"github.com/emperorsixpacks/dailbot/src/internal/services/airtable"
	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
)

// TODO the repository layer seems to be very repetitve we could fix that
func NewAirtableRepository(settings utils.ServerSettings, repo repositories.Repository[entities.AirtableWebhook]) *airtableWebhookUsecase {
	return &airtableWebhookUsecase{repo, settings}
}

type airtableWebhookUsecase struct {
	airtableRepository repositories.Repository[entities.AirtableWebhook]
	appSettings        utils.ServerSettings
}

func (a *airtableWebhookUsecase) CreateNewWebHook(wbkreq schemas.WebHookCreateSchema, secret string) {
	airtableService := airtable.NewAirtableSerice(a.appSettings, wbkreq.PersonalToken)
	webhook, err := airtableService.CreateNewWebhookURL(wbkreq.BaseID)
	if err != nil {
		logger.DefaultLogger.Error(err)
		return
	}
	newWebhook := entities.NewWebHook(wbkreq.BaseID, webhook.MacSecretBase64, wbkreq.TableIDs)
	newWebhook.SetPToken(wbkreq.PersonalToken, secret)

}
