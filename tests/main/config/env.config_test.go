package test_config

import (
	"testing"

	"github.com/claudiomozer/starwarsapi/src/main/config"
)

func TestShouldProvidesEnvironmentVarsFromEnvFile(t *testing.T) {
	env, err := config.LoadVars("../../../.env.example")

	if env.DB_HOST != "host" {
		t.Error("Should load exactly the same DB_HOST from env sample file")
	}

	if env.DB_NAME != "name" {
		t.Error("Should load exactly the same DB_NAME from env sample file")
	}

	if env.DB_PORT != "port" {
		t.Error("Should load exactly the same DB_PORT from env sample file")
	}

	if env.DB_USER != "user" {
		t.Error("Should load exactly the same DB_USER from env sample file")
	}

	if env.DB_PASS != "pass" {
		t.Error("Should load exactly the same DB_PASS from env sample file")
	}

	if err != nil {
		t.Error("Should not return erros")
	}
}
