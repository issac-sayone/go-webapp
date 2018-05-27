package config

import (
	"os"
)

// Environment configuration file
// Multiple environment configurations can be configured for switching

const MaxAge int = 365 * 24 * 60 * 60

//Env Enviornment config
type Env struct {
	DEBUG             bool
	HOST              string
	SERVER_PORT       string
	ACCESS_LOG        bool
	ACCESS_LOG_PATH   string
	ERROR_LOG         bool
	ERROR_LOG_PATH    string
	VERSION           string
}

var enviornment = Env{
	DEBUG: Getenv("DEBUG"),
	SERVER_PORT:       os.Getenv("SERVER_PORT"),
	HOST:              os.Getenv("HOST"),
	ACCESS_LOG:      Getenv("ACCESS_LOG"),
	ACCESS_LOG_PATH: os.Getenv("ACCESS_LOG_PATH"),
	ERROR_LOG:      Getenv("ERROR_LOG"),
	ERROR_LOG_PATH: os.Getenv("ERROR_LOG_PATH"),
	VERSION:         os.Getenv("VERSION"),
}

//GetEnv get the current enviornment configuration
func GetEnv() *Env {
	return &enviornment
}

//GetEnv ... to get the value from enviornment
func Getenv(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}
