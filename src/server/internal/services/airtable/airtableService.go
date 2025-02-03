package airtable

import (
	"encoding/json"
	"fmt"

	"github.com/emperorsixpacks/dailbot/pkg/utils"
	"github.com/gofiber/fiber/v2"
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
	return &airtableService{
		config: cfg,
	}
}

type airtableService struct {
	config utils.AirtableSettings
}

func (this airtableService) Authorise() {
}

func (this airtableService) GetAllTables(){}

func (this airtableService) CreateNewWebhookURL(){}

func (this airtableService) GetAllRecords(){}
