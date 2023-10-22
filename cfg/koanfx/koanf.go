package koanfx

import (
	"fmt"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func ReadFromFile(cfg *koanf.Koanf, path string, parser koanf.Parser) error {
	if err := cfg.Load(file.Provider(path), parser); err != nil {
		// Config file was found but another error was produced
		return fmt.Errorf("error loading config from file '%s': %w", path, err)
	}
	return nil
}

func ReadFromJSONFile(cfg *koanf.Koanf, path string) error {
	return ReadFromFile(cfg, path, json.Parser())
}

func ReadFromYAMLFile(cfg *koanf.Koanf, path string) error {
	return ReadFromFile(cfg, path, yaml.Parser())
}
