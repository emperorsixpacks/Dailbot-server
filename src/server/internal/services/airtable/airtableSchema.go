package airtable

type oAuthSchema struct {
	client_id           string   `json:"client_id"`
	redirect_uri        string   `json:"redirect_uri"`
	response_type       string   `json:"response_type"`
	scope               []string `json:"scope"`
	state               string   `json:"state"`
	codeChallenge       string   `json:"code_challenge"`
	codeChallengeMethod string   `json:"code_challenge_method"`
}

type oAuthSuccessSchema struct {
	code                string
	state               string
	codeChallenge       string
	codeChallengeMethod string
}

type oAuthErroorSchema struct {
	errorStr         string
	errorDescription string
	state            string
}
