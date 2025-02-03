package airtable

import (
	"github.com/emperorsixpacks/dailbot/pkg/utils"
)

const (
	AUTHORISEURL    = "https://airtable.com/oauth2/v1/authorize"
	TOKENREQUESTURL = "https://airtable.com/oauth2/v1/token"
)

var (
	scopes = []string{
		"data.records:read",
		"data.records:write",
		"data.recordComments:read",
		"data.recordComments:write",
		"schema.bases:read",
		"schema.bases:write",
		"user.email:read",
		"webhook:manage",
	}
)

func NewAirtableSerice(cfg utils.AirtableSettings) *airtableService {
	return nil
}

type airtableService struct {
	config utils.AirtableSettings
}

func (this airtableService) Authorise() {
	config := this.config
	genState, err := utils.GenerateRandomString(32)
	if err != nil {
		// TODO we would want to log the error here
		return
	}
	challengeVerifier, err := utils.GenerateRandomString(64)
	if err != nil {
		// TODO we would want to log the error here
		return
	}
	requestParameters := oAuthSchema{
		client_id:           config.ClientID,
		redirect_uri:        "google.com",
		response_type:       "code",
		scope:               scopes,
		state:               genState,
		codeChallenge:       challengeVerifier,
		codeChallengeMethod: "s256",
	}
}
