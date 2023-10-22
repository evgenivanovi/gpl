package cfg

import (
	"testing"
)

const (
	Dir        = "../testdata/cfg"
	ConfigJSON = Dir + "/cfg.json"
	ConfigYAML = Dir + "/cfg.yaml"
)

func TestJSONFileSource(t *testing.T) {

	expected := "json"

	cfg := NewJSONFileSource("type", ConfigJSON)

	if actual, _ := cfg.Get(); actual != expected {
		t.Errorf("Get() = '%v', want '%v'", actual, expected)
	}

}

func TestJSONYAMLSource(t *testing.T) {

	expected := "yaml"

	cfg := NewYAMLFileSource("type", ConfigYAML)

	if actual, _ := cfg.Get(); actual != expected {
		t.Errorf("Get() = '%v', want '%v'", actual, expected)
	}

}
