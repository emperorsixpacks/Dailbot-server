package airtable

// NOTE future
type oAuthSchema struct {
	ClientID            string   `json:"client_id"`
	Redirect_uri        string   `json:"redirect_uri"`
	Response_type       string   `json:"response_type"`
	Scope               []string `json:"scope"`
	State               string   `json:"state"`
	CodeChallenge       string   `json:"code_challenge"`
	CodeChallengeMethod string   `json:"code_challenge_method"`
}
//

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
