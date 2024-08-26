package commands

import (
	"os"
	"task-tracker/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmd_Init(t *testing.T) {
	// arrage
	cmd := Cmd{}

	// act
	cmd.Init()

	// assert
	assert.NotNil(t, cmd.commands)
	assert.Empty(t, cmd.commands)
}

func TestCmd_Add(t *testing.T) {
	// arrange
	cmd := Cmd{}

	// act
	cmd.Init()
	cmd.Add("test", func(args []string, r repository.Repository) (string, error) {
		return "success", nil
	})

	// assert
	assert.Contains(t, cmd.commands, "test")
}

func TestCmd_Execute_NoArgs(t *testing.T) {
	// arrange
	cmd := Cmd{}
	os.Args = []string{`cmd`}

	// act
	cmd.Init()
	_, err := cmd.Execute()

	// assert
	assert.EqualError(t, err, "too few arguments")
}

func TestCmd_Execute_UnknownCommand(t *testing.T) {
	// arrange
	cmd := Cmd{}
	cmd.Init()
	os.Args = []string{"cmd", "test"}

	// act
	_, err := cmd.Execute()

	// assert
	assert.EqualError(t, err, "unknown command")
}

func TestCmd_Execute_Success(t *testing.T) {
	// arrange
	cmd := Cmd{}
	os.Args = []string{"cmd", "test"}

	// act
	cmd.Init()
	cmd.Add("test", func(args []string, r repository.Repository) (string, error) {
		return "success", nil
	})
	result, err := cmd.Execute()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "success", result)
}
