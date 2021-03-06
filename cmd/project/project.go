package project

import (
	"github.com/gookit/gcli/v3"
)

func Cmd() *gcli.Command {
	cmd := &gcli.Command{
		Name: "project",
		// allow color tag and {$cmd} will be replace to 'demo'
		Desc: "Querying subcommands",
		Func: func(cmd *gcli.Command, args []string) error {
			cmd.ShowHelp()
			return nil
		},
		Subs: []*gcli.Command{initCMD()},
	}

	return cmd
}

func initCMD() *gcli.Command {
	cmd := &gcli.Command{
		Name: "init",
		Desc: "Create a new project in the current directory",
		Func: func(cmd *gcli.Command, args []string) error {
			cmd.ShowHelp()
			return nil
		},
	}

	return cmd
}
