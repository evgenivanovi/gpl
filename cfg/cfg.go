package cfg

import (
	"errors"

	"github.com/evgenivanovi/gpl/std"
	"github.com/knadh/koanf/v2"
)

var cfg = koanf.New(std.Dot)

const PropertyNotFoundFormat = "property '%s' not found in sources"

var PropertyNotFoundError = errors.New("property not found in sources")
