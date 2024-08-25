package utils

import (
	"encoding/json"
	"os"
	"task-tracker/internal/config"
)

func WriteJSON(data any) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(config.Filename, b, 0666); err != nil {
		return err
	}

	return nil
}
