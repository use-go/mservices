// Package cli is a command line interface
package cli

import (
	"fmt"
	"os"
	osexec "os/exec"
	"strings"

	"github.com/chzyer/readline"
	"github.com/micro/micro/v3/client/cli/util"
	"github.com/micro/micro/v3/cmd"
	"github.com/urfave/cli/v2"

	_ "github.com/micro/micro/v3/client/cli/auth"
	_ "github.com/micro/micro/v3/client/cli/config"
	_ "github.com/micro/micro/v3/client/cli/gen"
	_ "github.com/micro/micro/v3/client/cli/init"
	_ "github.com/micro/micro/v3/client/cli/network"
	_ "github.com/micro/micro/v3/client/cli/new"
	_ "github.com/micro/micro/v3/client/cli/run"
	_ "github.com/micro/micro/v3/client/cli/store"
	_ "github.com/micro/micro/v3/client/cli/user"
)

var (
	prompt = "micro> "

	// TODO: only run fixed set of commands for security purposes
	commands = map[string]*command{}
)

type command struct {
	name  string
	usage string
	exec  util.Exec
}

func Run(c *cli.Context) error {
	// take the first arg as the binary
	binary := os.Args[0]

	r, err := readline.New(prompt)
	if err != nil {
		return err
	}
	defer r.Close()

	for {
		args, err := r.Readline()
		if err != nil {
			fmt.Fprint(os.Stdout, err)
			return err
		}

		args = strings.TrimSpace(args)

		// skip no args
		if len(args) == 0 {
			continue
		}

		parts := strings.Split(args, " ")
		if len(parts) == 0 {
			continue
		}

		cmd := osexec.Command(binary, parts...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(string(err.(*osexec.ExitError).Stderr))
		}
	}

	return nil
}

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "cli",
			Usage:  "Run the interactive CLI",
			Action: Run,
		},
		&cli.Command{
			Name:   "call",
			Usage:  `Call a service e.g micro call greeter Say.Hello '{"name": "John"}'`,
			Action: util.Print(CallService),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "address",
					Usage:   "Set the address of the service instance to call",
					EnvVars: []string{"MICRO_ADDRESS"},
				},
				&cli.StringFlag{
					Name:    "output, o",
					Usage:   "Set the output format; json (default), raw",
					EnvVars: []string{"MICRO_OUTPUT"},
				},
				&cli.StringSliceFlag{
					Name:    "metadata",
					Usage:   "A list of key-value pairs to be forwarded as metadata",
					EnvVars: []string{"MICRO_METADATA"},
				},
				&cli.StringFlag{
					Name:  "request_timeout",
					Usage: "timeout duration",
				},
			},
		},
		&cli.Command{
			Name:  "get",
			Usage: `Get resources from micro`,
			Subcommands: []*cli.Command{
				{
					Name:   "service",
					Usage:  "Get a specific service from the registry",
					Action: util.Print(GetService),
				},
			},
		},
		&cli.Command{
			Name:   "health",
			Usage:  `Get the service health`,
			Action: util.Print(QueryHealth),
		},
		&cli.Command{
			Name:   "stream",
			Usage:  `Create a service stream e.g. micro stream foo Bar.Baz '{"key": "value"}'`,
			Action: util.Print(streamService),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output, o",
					Usage:   "Set the output format; json (default), raw",
					EnvVars: []string{"MICRO_OUTPUT"},
				},
				&cli.StringSliceFlag{
					Name:    "metadata",
					Usage:   "A list of key-value pairs to be forwarded as metadata",
					EnvVars: []string{"MICRO_METADATA"},
				},
			},
		},
		&cli.Command{
			Name:   "stats",
			Usage:  "Query the stats of specified service(s), e.g micro stats srv1 srv2 srv3",
			Action: util.Print(queryStats),
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "all",
					Usage: "to list all builtin services use --all builtin, for user's services use --all custom",
				},
			},
		},
		&cli.Command{
			Name:   "env",
			Usage:  "Get/set micro cli environment",
			Action: util.Print(listEnvs),
			Subcommands: []*cli.Command{
				{
					Name:   "get",
					Usage:  "Get the currently selected environment",
					Action: util.Print(getEnv),
				},
				{
					Name:   "set",
					Usage:  "Set the environment to use for subsequent commands e.g. micro env set dev",
					Action: util.Print(setEnv),
				},
				{
					Name:   "add",
					Usage:  "Add a new environment e.g. micro env add foo 127.0.0.1:8081",
					Action: util.Print(addEnv),
				},
				{
					Name:   "del",
					Usage:  "Delete an environment from your list e.g. micro env del foo",
					Action: util.Print(delEnv),
				},
			},
		},
		&cli.Command{
			Name:   "services",
			Usage:  "List services in the registry",
			Action: util.Print(ListServices),
		},
	)
}
