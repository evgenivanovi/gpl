package cfg

import (
	"errors"

	"github.com/evgenivanovi/gpl/std"
	"github.com/knadh/koanf/v2"
)

/* __________________________________________________ */

var cfg = koanf.New(std.Dot)

var PropertyNotFoundError = errors.New("property not found in sources")

/* __________________________________________________ */
