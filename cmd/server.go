package cmd

import (
	"flag"

	"github.com/coryodaniel/todo/internal/server"
)

func NewServerCommand() *ServerCommand {
	sc := &ServerCommand{
		fs: flag.NewFlagSet("server", flag.ContinueOnError),
	}

	sc.fs.StringVar(&sc.addr, "addr", ":5555", "Host address & port to listen on. Format hostname:port")

	return sc
}

type ServerCommand struct {
	fs   *flag.FlagSet
	addr string
}

// Name of this subcommand
func (s *ServerCommand) Name() string {
	return s.fs.Name()
}

// Init command and parse flags
func (s *ServerCommand) Init(args []string) error {
	return s.fs.Parse(args)
}

func (s *ServerCommand) Run() error {
	return server.NewServer(s.addr)
}
