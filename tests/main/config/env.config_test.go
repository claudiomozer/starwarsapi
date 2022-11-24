package test_config

import (
	"os"
	"testing"

	"github.com/claudiomozer/starwarsapi/src/main/config"
)

func mockEnvs() {
	os.Setenv("DB_NAME", "DB_TEST_NAME")
	os.Setenv("DB_HOST", "DB_TEST_HOST")
	os.Setenv("DB_PORT", "DB_TEST_PORT")
	os.Setenv("DB_USER", "DB_TEST_USER")
	os.Setenv("DB_PASS", "DB_TEST_PASS")
}

func TestShouldProvidesEnvironmentVarsFromEnvFile(t *testing.T) {
	env := config.LoadVars("../../../.env.example")

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
}

func TestShouldLoadEnviromentVarsFromOsIfNoPathIsGiven(t *testing.T) {
	mockEnvs()
	env := config.LoadVars("")

	if env.DB_HOST != "DB_TEST_HOST" {
		t.Error("Should load DB_TEST_HOST from os")
	}

	if env.DB_NAME != "DB_TEST_NAME" {
		t.Error("Should load DB_TEST_NAME from os")
	}

	if env.DB_PORT != "DB_TEST_PORT" {
		t.Error("Should load DB_TEST_PORT from os")
	}

	if env.DB_USER != "DB_TEST_USER" {
		t.Error("Should load DB_TEST_USER from os")
	}

	if env.DB_PASS != "DB_TEST_PASS" {
		t.Error("Should load DB_TEST_PASS from os")
	}
}
