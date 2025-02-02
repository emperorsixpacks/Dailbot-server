package utils 

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func validMapping(in interface{}) (map[string]interface{}, error) {
	comma, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.New("Invalid config")
	}
	return comma, nil
}

func validPath(configPath string) error {
	_, err := os.Stat(configPath)
	if !os.IsNotExist(err) {
		return err
	}
	return nil
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

func loadConfig(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config, err := ymltoMap(file)
	if err != nil {
		return nil, err
	}
	newConfig, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}
	return newConfig, nil
}

func ymltoMap(file []byte) (interface{}, error) {
	var config interface{}
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	err = resolveConfig(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// TODO why don't we just pass the variable around without the pointer
func resolveConfig(config *interface{}) error {
	MapConfig, err := validMapping((*config))
	if err != nil {
		return err
	}
	for k, v := range MapConfig {
		if MapConfig[k], err = resolveConfigVars(v); err != nil {
			return err
		}
	}

	return nil
}

// TODO accept any val of any type, currently only works with str
func resolveConfigVars(config interface{}) (interface{}, error) {
	MapConfig, err := validMapping(config)
	if err != nil {
		return nil, err
	}
	for k, v := range MapConfig {
		if str, ok := v.(string); ok {
			MapConfig[k] = resolvePlaceHolder(str)
			continue
		}
		if MapConfig[k], err = resolveConfigVars(v); err != nil {
			return nil, err
		}
	}
	return config, nil // MapConfig is a reference to config
}

func resolvePlaceHolder(value string) string {
	if strings.Contains(value, "${") {
		last_index := len(value) - 1
		first_index := 2
		env_value := value[first_index:last_index]
		return os.Getenv(env_value)
	}
	return value
}
