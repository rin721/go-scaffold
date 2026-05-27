package main

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/rei0721/go-scaffold/pkg/cli"
)

func TestTestsCommandMetadata(t *testing.T) {
	cmd := NewTestsCommand()

	if got := cmd.Name(); got != testsCommandName {
		t.Fatalf("Name() = %q, want %q", got, testsCommandName)
	}
	if !strings.Contains(cmd.Description(), "Go tests") {
		t.Fatalf("Description() = %q, want Go tests wording", cmd.Description())
	}
	if strings.Contains(strings.ToLower(cmd.Description()), "yaml") {
		t.Fatalf("Description() = %q, should not describe yaml2go demo behavior", cmd.Description())
	}
	if !strings.Contains(cmd.Usage(), "--package") {
		t.Fatalf("Usage() = %q, want package flag", cmd.Usage())
	}

	flags := cmd.Flags()
	if len(flags) != 1 {
		t.Fatalf("len(Flags()) = %d, want 1", len(flags))
	}
	if flags[0].Name != "package" || flags[0].ShortName != "p" {
		t.Fatalf("Flags()[0] = %+v, want package flag with short p", flags[0])
	}
	if flags[0].Default != defaultTestPackagePattern {
		t.Fatalf("package flag default = %v, want %q", flags[0].Default, defaultTestPackagePattern)
	}
}

func TestTestsCommandExecuteDefaultsToAllPackages(t *testing.T) {
	var gotArgs []string
	cmd := &TestsCommand{
		runner: func(_ io.Writer, _ io.Writer, args ...string) error {
			gotArgs = append([]string(nil), args...)
			return nil
		},
	}

	var stdout bytes.Buffer
	err := cmd.Execute(&cli.Context{
		Flags:  map[string]interface{}{},
		Stdout: &stdout,
		Stderr: io.Discard,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	wantArgs := []string{"test", defaultTestPackagePattern}
	if !reflect.DeepEqual(gotArgs, wantArgs) {
		t.Fatalf("runner args = %v, want %v", gotArgs, wantArgs)
	}
	if !strings.Contains(stdout.String(), "Running go test ./...") {
		t.Fatalf("stdout = %q, want command summary", stdout.String())
	}
	if !strings.Contains(stdout.String(), "Go tests passed") {
		t.Fatalf("stdout = %q, want pass summary", stdout.String())
	}
}

func TestTestsCommandExecuteUsesPackageFlag(t *testing.T) {
	var gotArgs []string
	cmd := &TestsCommand{
		runner: func(_ io.Writer, _ io.Writer, args ...string) error {
			gotArgs = append([]string(nil), args...)
			return nil
		},
	}

	err := cmd.Execute(&cli.Context{
		Flags: map[string]interface{}{
			"package": "./cmd/server",
		},
		Stdout: io.Discard,
		Stderr: io.Discard,
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	wantArgs := []string{"test", "./cmd/server"}
	if !reflect.DeepEqual(gotArgs, wantArgs) {
		t.Fatalf("runner args = %v, want %v", gotArgs, wantArgs)
	}
}

func TestTestsCommandExecuteReturnsRunnerError(t *testing.T) {
	wantErr := errors.New("runner failed")
	cmd := &TestsCommand{
		runner: func(_ io.Writer, _ io.Writer, _ ...string) error {
			return wantErr
		},
	}

	err := cmd.Execute(&cli.Context{
		Flags:  map[string]interface{}{},
		Stdout: io.Discard,
		Stderr: io.Discard,
	})
	if err == nil {
		t.Fatal("Execute() error = nil, want runner error")
	}
	if !errors.Is(err, wantErr) {
		t.Fatalf("Execute() error = %v, want wrapping %v", err, wantErr)
	}
	if !strings.Contains(err.Error(), "go test ./... failed") {
		t.Fatalf("Execute() error = %q, want command context", err.Error())
	}
}
