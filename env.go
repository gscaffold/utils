package utils

import "os"

const (
	EnvironmentLocal = "local"
	EnvironmentTest  = "test"
	EnvironmentStage = "staging"
	EnvironmentProd  = "prod"
)

var (
	app = os.Getenv("APP_NAME")
	env = os.Getenv("APP_ENV")
)

func IsProd() bool {
	return env == EnvironmentProd
}

func IsStage() bool {
	return env == EnvironmentStage
}

func IsTest() bool {
	return env == EnvironmentTest
}

func IsLocal() bool {
	return env == EnvironmentLocal
}

func GetEnv() string {
	return env
}

func GetApp() string {
	return app
}
