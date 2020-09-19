package config

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetInt(t *testing.T) {
	key := "a"
	value := 1

	viper := viper.New()
	viper.Set(key, value)

	cfg := &concreteConfig{
		cfg: viper,
	}

	if !cfg.IsSet(key) {
		t.Errorf("Expected to find a value at %s", key)
	}

	number := cfg.GetInt(key)
	if number != value {
		t.Errorf("Config returned invalid value. Got %d, expected %d", number, value)
	}
}

func TestGetInt_WithoutValue(t *testing.T) {
	key := "a"

	viper := viper.New()
	cfg := &concreteConfig{
		cfg: viper,
	}

	if cfg.IsSet(key) {
		t.Errorf("Expected to find no value at %s", key)
	}

	number := cfg.GetInt(key)
	if number != 0 {
		t.Errorf("Config returned invalid value. Got %d, expected 0", number)
	}
}

func TestGetString(t *testing.T) {
	key := "a"
	value := "b"

	viper := viper.New()
	viper.Set(key, value)

	cfg := &concreteConfig{
		cfg: viper,
	}

	if !cfg.IsSet(key) {
		t.Errorf("Expected to find a value at %s", key)
	}

	text := cfg.GetString(key)
	if text != value {
		t.Errorf("Config returned invalid value. Got %s, expected %s", text, value)
	}
}

func TestGetString_WithoutValue(t *testing.T) {
	key := "a"

	viper := viper.New()
	cfg := &concreteConfig{
		cfg: viper,
	}

	if cfg.IsSet(key) {
		t.Errorf("Expected to find no value at %s", key)
	}

	text := cfg.GetString(key)
	if text != "" {
		t.Errorf("Config returned invalid value. Got %s, expected an empty string", text)
	}
}
