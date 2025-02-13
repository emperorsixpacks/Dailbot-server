package airtable

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/emperorsixpacks/dailbot/src/pkg/logger"
	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

const (
	AUTHORIZEURL     = "https://airtable.com/oauth2/v1/authorize"
	TOKENREQUESTURL  = "https://airtable.com/oauth2/v1/token" // NOTE not currently using
	WEBHOOKCREATEURL = "https://api.airtable.com/v0/bases/%s/webhooks"
)

type mapping map[string]interface{}

// NOTE not currently using
var (
	scopes = []string{
		"data.records:read",
		"data.records:write",
		"schema.bases:read",
		"user.email:read",
		"webhook:manage",
	}
)

func NewAirtableSerice(settings utils.AppSettings) *AirtableService {
	logger.DefaultLogger.Info("Created new airtable service")
	return &AirtableService{nil, settings}
}

type AirtableService struct {
	agent       *fiber.Agent
	appSettings utils.AppSettings
}

func (this *AirtableService) AuthURL() string {
	state, err := utils.GenerateSecret()
	if err != nil {
		panic(err)
	}
	endpoint := oauth2.Endpoint{
		AuthURL:  AUTHORIZEURL,
		TokenURL: TOKENREQUESTURL,
	}
	authConfig := &oauth2.Config{
		ClientID:    this.appSettings.Services.Airtable.ClientID,
    ClientSecret: this.appSettings.Services.Airtable.ClientSecret,
		RedirectURL: this.appSettings.Server.AuthCallbackUrl,
		Scopes:      scopes,
		Endpoint:    endpoint,
	}
	codeVerifier, err := utils.GenerateRandomString(45)
	if err != nil {
		panic(err)
	}
	return authConfig.AuthCodeURL(state,
		oauth2.SetAuthURLParam("code_challenge", oauth2.S256ChallengeFromVerifier(codeVerifier)),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"))
}

func (this *AirtableService) send() (int, []byte) {
	if this.agent == nil {
		panic(errors.New("Run .request before calling .send"))
	}

	if err := this.agent.Parse(); err != nil {
		panic(err)
	}
	code, body, _ := this.agent.Bytes()
	return code, body
}

func (this *AirtableService) reqBody(b interface{}) *AirtableService {
	body, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	this.agent = this.agent.Body(body)
	return this
}

func (this *AirtableService) request(method, url string, pathParams ...string) *AirtableService {
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
	//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", this.parsonalToken))
	req.Header.Set("Content-Type", "application/json")

	req.SetRequestURI(reqUrl)
	this.agent = agent
	return this
}

func (this AirtableService) GetAllTables() {}

func (this AirtableService) CreateNewWebhookURL(baseID string) (webHookCreateSchemaReponse, error) {
	data := webHookCreateSchema{
		NotifyUrl: this.appSettings.Server.PublicUrl,
		Spec: mapping{
			"options": mapping{
				"filters": mapping{
					"dataTypes":   []string{"tableData"},
					"changeTypes": []string{"add", "remove", "update"},
					"fromSources": []string{"formSubmission", "formPageSubmission"},
				},
			},
		},
	}
	// TODO clean the error message to log only the message from airtbale
	req := this.request(fiber.MethodPost, WEBHOOKCREATEURL, baseID).reqBody(data)
	code, body := req.send()
	if code != 200 {
		// TODO log error message here
		logger.DefaultLogger.Error(string(body))
		return webHookCreateSchemaReponse{}, fmt.Errorf("Could not create webhook")
	}
	var responseSechema webHookCreateSchemaReponse
	err := json.Unmarshal(body, &responseSechema)
	if err != nil {
		// TODO log error here
		logger.DefaultLogger.Error(string(body))
		return webHookCreateSchemaReponse{}, fmt.Errorf("Could not create webhook")
	}
	logger.DefaultLogger.Println("Created new webhook")
	return responseSechema, nil
}

func (this AirtableService) GetAllRecords() {}
