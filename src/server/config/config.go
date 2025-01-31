package config

import (
	"path"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once         sync.Once
	app_settings *AppSettings
)

type (
	AppSettings struct {
		Server   ServerSettings     `yaml:"server"`
		Services APISericesSettings `yaml:"api_services"`
	}
	ServerSettings struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	APISericesSettings struct {
		Twilio TwilioSettings `yaml:"twilio"`
	}
	TwilioSettings struct {
		TwilioSSID         string `yaml:"twilio_ssid"`
		TwilioAuthToken    string `yaml:"twilio_auth_token"`
		TwilioPhoneNumeber string `yaml:"twilio_phone_number"`
	}

// AirtableSettings struct{}
)

func loadEnv() {
	pathStr, err := GetBasePath() // TODO what if we could just thow the error at the upper level
	if err != nil {
		// TODO raise an error here
		return
	}
	envPath := path.Join(pathStr, ".env")

	err = validPath(envPath)
	if err != nil {
		// TODO raise an error here
		return
	}
	err = godotenv.Load(envPath)
	if err != nil {
		// TODO raise the error here
		return
	}
}

func GetConfig() AppSettings {
	pathStr, err := GetBasePath()
	if err != nil {
		panic(err)
	}
	ymlPath := path.Join(pathStr, "src/server/config.yml")
	err = validPath(ymlPath) // TODO would it not be nice if we could just pass this as a slice

	if err != nil {
		// TODO raise an error here
		panic(err)

	}
	once.Do(func() {
		loadEnv()
		if app_settings == nil {
			err := LoadConfig(ymlPath, &app_settings)
			if err != nil {
				panic(err)
			}
		}
	})
	return *app_settings
}
