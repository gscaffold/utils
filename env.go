package utils

import "os"

const (
	EnvironmentLocal = "local"
	EnvironmentTest  = "test"
	EnvironmentStage = "staging"
	EnvironmentProd  = "prod"
)

func IsProd() bool {
	return GetEnv() == EnvironmentProd
}

func IsStage() bool {
	return GetEnv() == EnvironmentStage
}

func IsTest() bool {
	return GetEnv() == EnvironmentTest
}

func IsLocal() bool {
	return GetEnv() == EnvironmentLocal
}

func GetEnv() string {
	return os.Getenv("env")
}
