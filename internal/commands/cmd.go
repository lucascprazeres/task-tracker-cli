package commands

import (
	"errors"
	"os"
	"task-tracker/internal/config"
	"task-tracker/internal/repository"
	"task-tracker/internal/utils"
)

type Command = func(args []string, r repository.Repository) (string, error)

type Cmd struct {
	commands   map[string]Command
	repository repository.Repository
}

func New() Cmd {
	cmd := Cmd{
		repository: repository.NewJSONRepository(),
	}

	cmd.Init()
	cmd.Add("add", AddTask)
	cmd.Add("list", List)
	cmd.Add("delete", Delete)
	cmd.Add("update", Update)
	cmd.Add("mark-in-progress", MarkInProgress)
	cmd.Add("mark-done", MarkDone)

	return cmd
}

func (c *Cmd) Execute() (string, error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return "", errors.New("too few arguments")
	}

	name := args[0]
	command, ok := c.commands[name]
	if !ok {
		return "", errors.New("unknown command")
	}

	if err := utils.CreateFileIfNotExists(config.Filename); err != nil {
		return "", err
	}

	return command(args[1:], c.repository)
}

func (c *Cmd) Init() {
	c.commands = make(map[string]Command)
}

func (c *Cmd) Add(name string, command Command) {
	c.commands[name] = command
}
