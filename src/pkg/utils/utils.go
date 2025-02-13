package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math/big"
	"path"
	"runtime"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

var (
	once         sync.Once
	app_settings *AppSettings
)

type (
	AppSettings struct {
		Server   ServerSettings      `yaml:"server"`
		Services APISericesSettings  `yaml:"api_services"`
		DB       PersistenceSettings `yaml:"persistence"`
	}
	ServerSettings struct {
		Name            string `yaml:"name"`
		Port            string `yaml:"port"`
		Host            string `yaml:"host"`
		AuthCallbackUrl string `yaml:"auth_callback_url"`
		PublicUrl       string `yaml:"public_url"`
		TemplateDir     string `yaml:"templates_dir"`
		StaticDir       string `yaml:"static_dir"`
	}
	APISericesSettings struct {
		Twilio   TwilioSettings   `yaml:"twilio"`
		Airtable AirtableSettings `yaml:"airtable"`
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
	DBSettings struct {
		Host          string `yaml:"host"`
		Port          string `yaml:"port"`
		ConnectionUrl string `yaml:"connection_url"`
		UserName      string `yaml:"db_username"`
		Password      string `yaml:"password"`
		DB            string `yaml:"db"` // NOTE for Redis connection
		DataBase      string `yaml:"database"`
	}

	PersistenceSettings struct {
		PostgresQl DBSettings `yaml:"potgres"`
	}
)

// Returns the root dir
func GetBasePath() (string, error) {
	_, basePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("Could not get file path")
	}
	return path.Dir(path.Dir(path.Dir(path.Dir(basePath)))), nil
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
	ymlPath := path.Join(pathStr, "config/config.yml")
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

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func GenerateSecret() (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(salt), nil
}

func HashSting(s string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashStr), err
}

func EncryptString(s string, secret []byte) (string, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, []byte(s), nil)

	return base64.RawStdEncoding.EncodeToString(append(nonce, cipherText...)), nil
}
