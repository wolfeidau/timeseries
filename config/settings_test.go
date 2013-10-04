package config

import (
	"testing"
)

func TestLoadSettings(t *testing.T) {

	settings := LoadSettings("test.env")

	t.Fatalf("%v", settings)

}
