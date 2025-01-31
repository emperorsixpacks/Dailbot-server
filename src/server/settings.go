package main

type (
	AppSettings struct {
		server       ServerSettings
		api_services APISericesSettings
	}
	ServerSettings struct {
		name string
		port int16
		host string
	}
	APISericesSettings struct {
		twilio TwilioSettings
	}
	TwilioSettings struct {
		twilio_ssid          string
		twilio_auth_token    string
		twilio_phone_numeber string
	}

// AirtableSettings struct{}
)
