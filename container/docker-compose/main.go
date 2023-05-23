package main

import (
	"fmt"
	dockercli "github.com/docker/cli/cli"
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"os"

	commands "github.com/docker/compose/v2/cmd/compose"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
)

type logOut struct {
	cmd []string
}

func (l *logOut) Write(p []byte) (n int, err error) {
	fmt.Printf("execute %v command result: %s\n", l.cmd, string(p))
	return len(p), nil
}

type logErr struct {
	cmd []string
}

func (l *logErr) Write(p []byte) (n int, err error) {
	fmt.Printf("execute %v command result: %s\n", l.cmd, string(p))
	return len(p), nil
}

func main() {
	os.Args = []string{"docker", "compose", "up"}
	makeCmd := func(dockerCli command.Cli) *cobra.Command {
		serviceProxy := api.NewServiceProxy().WithService(compose.NewComposeService(dockerCli))
		cmd := commands.RootCommand(dockerCli, serviceProxy)
		originalPreRun := cmd.PersistentPreRunE
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			if err := plugin.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			if originalPreRun != nil {
				return originalPreRun(cmd, args)
			}
			return nil
		}
		cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
			return dockercli.StatusError{
				StatusCode: compose.CommandSyntaxFailure.ExitCode,
				Status:     err.Error(),
			}
		})

		// 打印输出参数
		cmd.SetOut(&logOut{os.Args})
		cmd.SetErr(&logErr{os.Args})
		return cmd
	}
	meta := manager.Metadata{
		SchemaVersion: "0.1.0",
		Vendor:        "Docker Inc.",
		Version:       "dev",
	}

	dockerCli, err := command.NewDockerCli()
	if err != nil {
		panic(err)
	}

	makePlugin := makeCmd(dockerCli)

	if err := plugin.RunPlugin(dockerCli, makePlugin, meta); err != nil {
		panic(err)
	}
}
