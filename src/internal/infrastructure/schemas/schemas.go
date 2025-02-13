package schemas

type WebHookCreateSchema struct {
	PersonalToken string   `json:"personal-token"`
	BaseID        string   `json:"base-id"`
	TableIDs      []string `json:"table-ids"`
}

type WebHookCreateResposeSchema struct {
	Status  int    `json:"status-code"`
	Message string `json:"message"`
}
