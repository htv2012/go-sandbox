package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage[T any] struct {
	FileName string
}

func GetDataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	absPath := filepath.Join(home, ".config", "todo.json")
	return absPath, nil
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
