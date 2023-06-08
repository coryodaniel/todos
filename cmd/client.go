package cmd

import (
	"flag"
	"fmt"
	"strings"

	"github.com/coryodaniel/todo/pkg/sdk"
)

func NewClientCommand() *ClientCommand {
	cc := &ClientCommand{
		fs: flag.NewFlagSet("client", flag.ContinueOnError),
	}

	cc.fs.StringVar(&cc.url, "url", "http://localhost:5555", "Todo API URL")

	return cc
}

type ClientCommand struct {
	fs  *flag.FlagSet
	url string
}

// Name of this subcommand
func (c *ClientCommand) Name() string {
	return c.fs.Name()
}

// Init command and parse flags
func (c *ClientCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *ClientCommand) Run() error {
	todoClient := sdk.NewClient(c.url)

	switch subcommand := c.fs.Arg(0); subcommand {
	case "list":
		todos, _ := todoClient.ListTodos()
		fmt.Println("Here are your todos:")

		for _, t := range *todos {
			fmt.Printf("- %s [%s]\n", t.Title, t.ID)
		}

	case "create":
		title := strings.Join(c.fs.Args()[1:], " ")
		todo, _ := todoClient.CreateTodo(title)

		fmt.Printf("Added todo: %s [%s]\n", todo.Title, todo.ID)
	default:
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}

	return nil
}
