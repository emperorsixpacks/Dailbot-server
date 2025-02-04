package schemas

type WebHookCreateSchema struct {
	BaseID   string   `json:"base-id"`
	TableIDs []string `json:table-ids"`
}
