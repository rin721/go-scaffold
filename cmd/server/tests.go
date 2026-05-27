package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/rei0721/go-scaffold/pkg/cli"
)

const (
	testsCommandName          = "tests"
	defaultTestPackagePattern = "./..."
)

type goTestRunner func(stdout, stderr io.Writer, args ...string) error

type TestsCommand struct {
	runner goTestRunner
}

func NewTestsCommand() *TestsCommand {
	return &TestsCommand{}
}

func (c *TestsCommand) Name() string {
	return testsCommandName
}

func (c *TestsCommand) Description() string {
	return "Run Go tests"
}

func (c *TestsCommand) Usage() string {
	return fmt.Sprintf("%s [--package=<pattern>]", testsCommandName)
}

func (c *TestsCommand) Flags() []cli.Flag {
	return []cli.Flag{
		{
			Name:        "package",
			ShortName:   "p",
			Type:        cli.FlagTypeString,
			Required:    false,
			Default:     defaultTestPackagePattern,
			Description: "Go package pattern passed to go test",
		},
	}
}

func (c *TestsCommand) Execute(ctx *cli.Context) error {
	pattern := defaultTestPackagePattern
	if ctx != nil {
		if value := ctx.GetString("package"); value != "" {
			pattern = value
		}
	}

	stdout, stderr := commandOutput(ctx)
	args := []string{"test", pattern}
	fmt.Fprintf(stdout, "Running go %s\n", strings.Join(args, " "))
	if err := c.run(stdout, stderr, args...); err != nil {
		return fmt.Errorf("go %s failed: %w", strings.Join(args, " "), err)
	}
	fmt.Fprintln(stdout, "Go tests passed")
	return nil
}

func (c *TestsCommand) run(stdout, stderr io.Writer, args ...string) error {
	if c.runner != nil {
		return c.runner(stdout, stderr, args...)
	}
	return runGoTest(stdout, stderr, args...)
}

func runGoTest(stdout, stderr io.Writer, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}

func commandOutput(ctx *cli.Context) (io.Writer, io.Writer) {
	if ctx == nil {
		return os.Stdout, os.Stderr
	}

	stdout := ctx.Stdout
	if stdout == nil {
		stdout = os.Stdout
	}

	stderr := ctx.Stderr
	if stderr == nil {
		stderr = os.Stderr
	}

	return stdout, stderr
}
