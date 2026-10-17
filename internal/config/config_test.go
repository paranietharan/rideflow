package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := Load()

	if cfg == nil {
		t.Fatal("expected config to not be nil")
	}

	if cfg.ServerPort == "" {
		t.Fatal("expected ServerPort to be set")
	}

	if cfg.MongoURI == "" {
		t.Fatal("expected MongoURI to be set")
	}
}

func TestGetEnv(t *testing.T) {
	key := "TEST_ENV_VAR_RIDEFLOW"
	expected := "custom_val"
	_ = os.Setenv(key, expected)
	defer os.Unsetenv(key)

	if val := getEnv(key, "default"); val != expected {
		t.Fatalf("expected %s, got %s", expected, val)
	}

	if val := getEnv("NON_EXISTENT_KEY", "default"); val != "default" {
		t.Fatalf("expected default, got %s", val)
	}
}
