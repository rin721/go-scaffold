package cli

// 本测试文件固定轻量 CLI 解析器的命令、flag 和错误输出契约，防止注释补全和后续重构改变外部可观察行为。

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

type testCommand struct {
	name        string
	description string
	flags       []Flag
	ctx         *Context
	err         error
}

// Name 实现 CLI 测试命令的同名方法，帮助验证命令注册、flag 解析和执行错误路径。
func (c *testCommand) Name() string {
	return c.name
}

// Description 实现 CLI 测试命令的同名方法，帮助验证命令注册、flag 解析和执行错误路径。
func (c *testCommand) Description() string {
	return c.description
}

// Usage 实现 CLI 测试命令的同名方法，帮助验证命令注册、flag 解析和执行错误路径。
func (c *testCommand) Usage() string {
	return c.name + " [flags]"
}

// Flags 实现 CLI 测试命令的同名方法，帮助验证命令注册、flag 解析和执行错误路径。
func (c *testCommand) Flags() []Flag {
	return c.flags
}

// Execute 实现 CLI 测试命令的同名方法，帮助验证命令注册、flag 解析和执行错误路径。
func (c *testCommand) Execute(ctx *Context) error {
	c.ctx = ctx
	if c.err != nil {
		return c.err
	}
	return nil
}

// TestRunWithIOParsesFlagsEnvAndArgs 固定轻量 CLI 解析器的命令、flag 和错误输出契约，确保后续注释补全或结构调整不改变该场景。
func TestRunWithIOParsesFlagsEnvAndArgs(t *testing.T) {
	t.Setenv("CLI_TEST_OUTPUT", "from-env")

	cmd := &testCommand{
		name:        "run",
		description: "run command",
		flags: []Flag{
			{Name: "name", Type: FlagTypeString, Required: true},
			{Name: "count", Type: FlagTypeInt, Default: 1},
			{Name: "verbose", Type: FlagTypeBool},
			{Name: "tags", Type: FlagTypeStringSlice},
			{Name: "output", Type: FlagTypeString, EnvVar: "CLI_TEST_OUTPUT"},
		},
	}
	app := NewApp("tool")
	if err := app.AddCommand(cmd); err != nil {
		t.Fatalf("AddCommand() error = %v", err)
	}

	var stdout, stderr bytes.Buffer
	err := app.RunWithIO([]string{
		"run",
		"--name", "alice",
		"--count", "3",
		"--verbose",
		"--tags", "alpha,beta",
		"positional",
	}, strings.NewReader("input"), &stdout, &stderr)
	if err != nil {
		t.Fatalf("RunWithIO() error = %v", err)
	}

	if cmd.ctx == nil {
		t.Fatal("Execute was not called")
	}
	if got := cmd.ctx.GetString("name"); got != "alice" {
		t.Fatalf("GetString(name) = %q, want %q", got, "alice")
	}
	if got := cmd.ctx.GetInt("count"); got != 3 {
		t.Fatalf("GetInt(count) = %d, want 3", got)
	}
	if got := cmd.ctx.GetBool("verbose"); got != true {
		t.Fatalf("GetBool(verbose) = %v, want true", got)
	}
	if got, want := cmd.ctx.GetStringSlice("tags"), []string{"alpha", "beta"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("GetStringSlice(tags) = %#v, want %#v", got, want)
	}
	if got := cmd.ctx.GetString("output"); got != "from-env" {
		t.Fatalf("GetString(output) = %q, want %q", got, "from-env")
	}
	if got, want := cmd.ctx.Args, []string{"positional"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Args = %#v, want %#v", got, want)
	}
}

// TestRunWithIOMissingRequiredFlagReturnsUsageError 固定轻量 CLI 解析器的命令、flag 和错误输出契约，确保后续注释补全或结构调整不改变该场景。
func TestRunWithIOMissingRequiredFlagReturnsUsageError(t *testing.T) {
	cmd := &testCommand{
		name:  "run",
		flags: []Flag{{Name: "name", Type: FlagTypeString, Required: true}},
	}
	app := NewApp("tool")
	if err := app.AddCommand(cmd); err != nil {
		t.Fatalf("AddCommand() error = %v", err)
	}

	err := app.RunWithIO([]string{"run"}, nil, &bytes.Buffer{}, &bytes.Buffer{})
	var usageErr *UsageError
	if !errors.As(err, &usageErr) {
		t.Fatalf("RunWithIO() error = %T, want *UsageError", err)
	}
	if got := GetExitCode(err); got != ExitUsage {
		t.Fatalf("GetExitCode() = %d, want %d", got, ExitUsage)
	}
}

// TestRunWithIOWrapsCommandExecutionError 固定轻量 CLI 解析器的命令、flag 和错误输出契约，确保后续注释补全或结构调整不改变该场景。
func TestRunWithIOWrapsCommandExecutionError(t *testing.T) {
	cause := errors.New("boom")
	cmd := &testCommand{
		name:  "run",
		flags: []Flag{{Name: "name", Type: FlagTypeString, Required: true}},
		err:   cause,
	}
	app := NewApp("tool")
	if err := app.AddCommand(cmd); err != nil {
		t.Fatalf("AddCommand() error = %v", err)
	}

	err := app.RunWithIO([]string{"run", "--name", "alice"}, nil, &bytes.Buffer{}, &bytes.Buffer{})
	var commandErr *CommandError
	if !errors.As(err, &commandErr) {
		t.Fatalf("RunWithIO() error = %T, want *CommandError", err)
	}
	if !errors.Is(err, cause) {
		t.Fatalf("RunWithIO() error does not wrap cause")
	}
	if got := GetExitCode(err); got != ExitError {
		t.Fatalf("GetExitCode() = %d, want %d", got, ExitError)
	}
}

// TestAddCommandRejectsDuplicateNames 固定轻量 CLI 解析器的命令、flag 和错误输出契约，确保后续注释补全或结构调整不改变该场景。
func TestAddCommandRejectsDuplicateNames(t *testing.T) {
	app := NewApp("tool")
	if err := app.AddCommand(&testCommand{name: "run"}); err != nil {
		t.Fatalf("first AddCommand() error = %v", err)
	}

	err := app.AddCommand(&testCommand{name: "run"})
	if err == nil {
		t.Fatal("second AddCommand() error = nil, want duplicate error")
	}
	if !strings.Contains(err.Error(), ErrMsgDuplicateCommand) {
		t.Fatalf("second AddCommand() error = %q, want duplicate message", err.Error())
	}
}

// TestRunWithIOPrintsHelpAndVersion 固定轻量 CLI 解析器的命令、flag 和错误输出契约，确保后续注释补全或结构调整不改变该场景。
func TestRunWithIOPrintsHelpAndVersion(t *testing.T) {
	app := NewApp("tool")
	app.SetVersion("1.2.3")
	app.SetDescription("test cli")
	if err := app.AddCommand(&testCommand{name: "run", description: "run command"}); err != nil {
		t.Fatalf("AddCommand() error = %v", err)
	}

	var help bytes.Buffer
	if err := app.RunWithIO([]string{"--help"}, nil, &help, &bytes.Buffer{}); err != nil {
		t.Fatalf("RunWithIO(--help) error = %v", err)
	}
	helpText := help.String()
	for _, want := range []string{"tool v1.2.3", "test cli", "run command", "Usage:"} {
		if !strings.Contains(helpText, want) {
			t.Fatalf("help output %q does not contain %q", helpText, want)
		}
	}

	var version bytes.Buffer
	if err := app.RunWithIO([]string{"--version"}, nil, &version, &bytes.Buffer{}); err != nil {
		t.Fatalf("RunWithIO(--version) error = %v", err)
	}
	if got, want := strings.TrimSpace(version.String()), "tool version 1.2.3"; got != want {
		t.Fatalf("version output = %q, want %q", got, want)
	}
}
