package airtable


const (
	WEBHOOKAUTHORISEURL = "https://api.airtable.com/v0/bases/{baseId}/webhooks"
	TOKENREQUESTURL     = "https://airtable.com/oauth2/v1/token"
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

func NewAirtableSerice(baseID, parsonalToken string) *airtableService {
	return &airtableService{baseID, parsonalToken}
}

type airtableService struct {
	baseID        string
	parsonalToken string
}

func (this airtableService) authoriseURL(url) {

}

func (this airtableService) GetAllTables() {}

func (this airtableService) CreateNewWebhookURL() {

}

func (this airtableService) GetAllRecords() {}
