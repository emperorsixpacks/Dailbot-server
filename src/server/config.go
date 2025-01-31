package main

import (
	"path"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func loadEnv() {

	pathStr, err := GetBasePath() // TODO what if we could just thow the error at the upper level
	envPath := path.Join(pathStr, ".env")

	err = validPath(envPath)
	if err != nil {
		// TODO raise an error here
		return
	}
	godotenv.Load(envPath)
}

func GetConfig() {
	pathStr, err := GetBasePath()
	if err != nil {
		// TODO raise the error here
		return
	}
	ymlPath := path.Join(pathStr, "src/server")
	err = validPath(ymlPath) // TODO would it not be nice if we could just pass this as a slice

	if err != nil {
		// TODO raise an error here
		return
	}
	once.Do(func() {
    loadConfig(ymlPath)
  })
}
