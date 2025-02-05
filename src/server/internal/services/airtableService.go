package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	WEBHOOKCREATEURL = "https://api.airtable.com/v0/bases/%v/webhooks"
	TOKENREQUESTURL  = "https://airtable.com/oauth2/v1/token"
)

type mapping map[string]interface{}

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
	return &airtableService{baseID, parsonalToken, nil}
}

type airtableService struct {
	baseID        string
	parsonalToken string
	agent         *fiber.Agent
}

func (this airtableService) send() (int, []byte, []error) {
	if this.agent == nil {
		panic(errors.New("Run .request before calling .send"))
	}
	return this.agent.Bytes()
}

func (this *airtableService) reqBody(b interface{}) *airtableService {
	body, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	this.agent = this.agent.Body(body)
  return this
}

func (this *airtableService) request(method, url string, pathParams ...string) *airtableService {
	unpack := func() string {
		url := url
		for _, i := range pathParams {
			url = fmt.Sprintf(url, i)
		}
		return url
	}
	reqUrl := unpack()

	agent := fiber.AcquireAgent()
	req := agent.Request()
	req.Header.SetMethod(method)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", this.parsonalToken))

	req.SetRequestURI(reqUrl)
	this.agent = agent
	return this
}

func (this airtableService) GetAllTables() {}

func (this airtableService) CreateNewWebhookURL() {
	data := WebHookCreateSchema{
		spec: mapping{
			"options": mapping{
				"filters": mapping{
					"dataTypes":   []string{"tableData"},
					"changeTypes": []string{"add", "remove", "update"},
					"fromSources": []string{"formSubmission", "formPageSubmission"},
				},
			},
		},
	}
	req := this.request("POST", WEBHOOKCREATEURL).reqBody(data)
	req.send()
}

func (this airtableService) GetAllRecords() {}
