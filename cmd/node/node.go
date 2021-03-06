package node

import (
	"github.com/gookit/gcli/v3"
)

func Cmd() *gcli.Command {
	cmd := &gcli.Command{
		Name: "node",
		// allow color tag and {$cmd} will be replace to 'demo'
		Desc: "Interact with and get information about Nodes",
		Func: func(cmd *gcli.Command, args []string) error {
			cmd.ShowHelp()
			return nil
		},
		Subs: []*gcli.Command{statusCMD()},
	}

	return cmd
}

func statusCMD() *gcli.Command {
	cmd := &gcli.Command{
		Name: "status",
		Desc: "Query information about a node, the default is the current specified node",
		Func: func(cmd *gcli.Command, args []string) error {
			cmd.ShowHelp()
			return nil
		},
	}

	return cmd
}
