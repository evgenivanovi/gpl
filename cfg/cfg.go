package cfg

import (
	"errors"

	"github.com/evgenivanovi/gpl/std"
	"github.com/knadh/koanf/v2"
)

var cfg = koanf.New(std.Dot)

var ErrPropertyNotFound = errors.New("property not found in sources")
