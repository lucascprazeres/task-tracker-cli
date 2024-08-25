package cmd

import (
	"errors"
	"os"
	"task-tracker/internal/commands"
	"task-tracker/internal/config"
	"task-tracker/internal/utils"
)

type Command = func(p []string) (string, error)

type Cmd struct {
	commands map[string]Command
}

func New() Cmd {
	cmd := Cmd{}

	cmd.Init()
	cmd.Add("add", commands.AddTask)
	cmd.Add("list", commands.List)
	cmd.Add("delete", commands.Delete)
	cmd.Add("update", commands.Update)
	cmd.Add("mark-in-progress", commands.MarkInProgress)
	cmd.Add("mark-done", commands.MarkDone)

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

	params := args[1:]
	return command(params)
}

func (c *Cmd) Init() {
	c.commands = make(map[string]func(p []string) (string, error))
}

func (c *Cmd) Add(name string, command Command) {
	c.commands[name] = command
}
