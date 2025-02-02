package utils

import (
	"errors"
	"path"
	"runtime"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	once         sync.Once
	app_settings *AppSettings
)

type (
	AppSettings struct {
		Server      ServerSettings     `yaml:"server"`
		Services    APISericesSettings `yaml:"api_services"`
		TemplateDir string             `yaml:"templates_dir"`
		StaticDir   string             `yaml:"static_dir"`
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
		SSID         string `yaml:"twilio_ssid"`
		AuthToken    string `yaml:"twilio_auth_token"`
		PhoneNumeber string `yaml:"twilio_phone_number"`
	}
	AirtableSettings struct {
		ClientID     string `yaml:"airtable_client_id"`
		ClientSecret string `yaml:"airtable_client_secret"`
	}
)

func GetBasePath() (string, error) {
	_, basePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("Could not get file path")
	}
	return path.Dir(path.Dir(path.Dir(path.Dir(path.Dir(basePath))))), nil
}

func LoadConfig(filepath string, o interface{}) error {
	ymlBytes, err := loadConfig(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(ymlBytes, o)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig() AppSettings {
	pathStr, err := GetBasePath()
	if err != nil {
		panic(err)
	}
	ymlPath := path.Join(pathStr, "src/server/config/config.yml")
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
