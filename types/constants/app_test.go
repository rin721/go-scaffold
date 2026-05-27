package constants

import "testing"

func TestApplicationConstants(t *testing.T) {
	if AppPrefix != "Rin" {
		t.Fatalf("AppPrefix = %q, want %q", AppPrefix, "Rin")
	}

	if AppServerCommandName != "server" {
		t.Fatalf("AppServerCommandName = %q, want %q", AppServerCommandName, "server")
	}

	if EnvConfigPathName != "RIN_CONFIG_PATH" {
		t.Fatalf("EnvConfigPathName = %q, want %q", EnvConfigPathName, "RIN_CONFIG_PATH")
	}
}
