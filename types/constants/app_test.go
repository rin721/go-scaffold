package constants

import "testing"

func TestApplicationConstants(t *testing.T) {
	if AppPrefix != "Rin" {
		t.Fatalf("AppPrefix = %q, want %q", AppPrefix, "Rin")
	}

	if AppServerCommandName != "server" {
		t.Fatalf("AppServerCommandName = %q, want %q", AppServerCommandName, "server")
	}

	if AppInitDBCommandName != "initdb" {
		t.Fatalf("AppInitDBCommandName = %q, want %q", AppInitDBCommandName, "initdb")
	}

	if EnvConfigPathName != "RIN_CONFIG_PATH" {
		t.Fatalf("EnvConfigPathName = %q, want %q", EnvConfigPathName, "RIN_CONFIG_PATH")
	}
}
