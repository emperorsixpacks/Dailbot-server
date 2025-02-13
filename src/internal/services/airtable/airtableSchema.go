package airtable

type oAuthErroorSchema struct {
	errorStr         string
	errorDescription string
	state            string
}

type webHookCreateSchema struct {
	NotifyUrl string                 `json:"notificationUrl"`
	Spec      map[string]interface{} `json:"specification"`
}

type webHookCreateSchemaReponse struct {
	ExpirationTime  string `json:"expirationTime"`
	Id              string `json:"id"`
	MacSecretBase64 string `json:"macSecreteBase64"`
}
