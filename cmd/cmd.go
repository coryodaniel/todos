package cmd

import (
	"errors"
	"fmt"
	"os"
)

var banner string = `
@@@@@@@   @@@@@@   @@@@@@@    @@@@@@
@@@@@@@  @@@@@@@@  @@@@@@@@  @@@@@@@@
  @@!    @@!  @@@  @@!  @@@  @@!  @@@
  !@!    !@!  @!@  !@!  @!@  !@!  @!@
  @!!    @!@  !@!  @!@  !@!  @!@  !@!
  !!!    !@!  !!!  !@!  !!!  !@!  !!!
  !!:    !!:  !!!  !!:  !!!  !!:  !!!
  :!:    :!:  !:!  :!:  !:!  :!:  !:!
   ::    ::::: ::   :::: ::  ::::: ::
   :      : :  :   :: :  :    : :  :
`

// Runner executes subcommands
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Execute(args []string) error {
	cmds := []Runner{
		NewServerCommand(),
		NewClientCommand(),
	}

	if len(args) < 1 {
		errMsg := banner + "\nEnter a subcommand:\n"

		for _, cmd := range cmds {
			errMsg = errMsg + "- " + cmd.Name() + "\n"
		}

		return errors.New(errMsg)
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
