package cfg

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/evgenivanovi/gpl/cfg/koanfx"
	"github.com/evgenivanovi/gpl/std"
	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/gookit/goutil/strutil"
)

type Source interface {
	Get() (string, bool)
	Map(mapping func(string) (any, error)) stdx.Value
}

type ValueSource struct {
	value string
}

func (s ValueSource) Get() (string, bool) {
	return s.value, true
}

func (s ValueSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewValueSource(value string) *ValueSource {
	return &ValueSource{
		value: value,
	}
}

type ArgSource struct {
	key string
}

func (s ArgSource) Get() (string, bool) {

	if !flag.Parsed() {
		flag.Parse()
	}

	lookup := flag.Lookup(s.key)

	if lookup == nil {
		return "", false
	}

	value := lookup.Value.String()
	if value == "" {
		return "", false
	}

	return value, true

}

func (s ArgSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewArgSource(key string) *ArgSource {
	return NewArgSourceWithOps(key, "", "")
}

func NewArgSourceWithOps(key, onAbsence, usage string) *ArgSource {

	if strutil.IsEmpty(key) {
		panic("key is required")
	}

	flag.String(key, onAbsence, usage)

	return &ArgSource{
		key: key,
	}

}

type EnvSource struct {
	key string
}

func (s EnvSource) Get() (string, bool) {
	return os.LookupEnv(s.key)
}

func (s EnvSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewEnvSource(key string) *EnvSource {

	if strutil.IsEmpty(key) {
		panic("key is required")
	}

	return &EnvSource{
		key: key,
	}

}

type FileSourceOp func(*FileSource)

func WithKey(key string) FileSourceOp {
	return func(source *FileSource) {
		source.key = key
	}
}

func WithKeyFn(key func() string) FileSourceOp {
	return func(source *FileSource) {
		source.key = key()
	}
}

func WithPath(path string) FileSourceOp {
	return func(source *FileSource) {
		source.path = conv.Supplier(path)
	}
}

func WithPathFn(path func() string) FileSourceOp {
	return func(source *FileSource) {
		source.path = path
	}
}

func withOnce(fn func()) FileSourceOp {
	return func(source *FileSource) {
		source.once = sync.OnceFunc(fn)
	}
}

type FileSource struct {
	key  string
	path func() string
	once func()
}

func (s FileSource) Get() (string, bool) {

	s.once()

	value := cfg.Get(s.key)
	if value == nil {
		return std.Empty, false
	}

	return stdx.NewValue(value).String(), true

}

func (s FileSource) Map(mapping func(string) (any, error)) stdx.Value {

	value, ok := s.Get()
	if !ok {
		return stdx.NewValue(nil)
	}

	val, err := mapping(value)
	if err != nil {
		return stdx.NewNilValue(err)
	}

	return stdx.NewValue(val)

}

func NewJSONFileSource(key string, path string) *FileSource {

	once := func() {

		if strutil.IsBlank(path) {
			msg := fmt.Sprintf("filepath is empty, not able to find value for key='%s'", key)
			slogx.Log().Debug(msg)
			return
		}

		err := koanfx.ReadFromJSONFile(cfg, path)
		if err != nil {
			slogx.Log().Debug(err.Error())
		}

	}

	return newFileSourceWithOp(
		WithKey(key),
		WithPath(path),
		withOnce(once),
	)

}

func NewJSONFileSourceWithPath(key string, path func() string) *FileSource {

	once := func() {

		if strutil.IsBlank(path()) {
			msg := fmt.Sprintf("filepath is empty, not able to find value for key='%s'", key)
			slogx.Log().Debug(msg)
			return
		}

		err := koanfx.ReadFromJSONFile(cfg, path())
		if err != nil {
			slogx.Log().Debug(err.Error())
		}

	}

	return newFileSourceWithOp(
		WithKey(key),
		WithPathFn(path),
		withOnce(once),
	)

}

func NewYAMLFileSource(key string, path string) *FileSource {

	once := func() {

		if strutil.IsBlank(path) {
			msg := fmt.Sprintf("filepath is empty, not able to find value for key='%s'", key)
			slogx.Log().Debug(msg)
			return
		}

		err := koanfx.ReadFromYAMLFile(cfg, path)
		if err != nil {
			slogx.Log().Debug(err.Error())
		}

	}

	return newFileSourceWithOp(
		WithKey(key),
		WithPath(path),
		withOnce(once),
	)

}

func NewYAMLFileSourceWithPath(key string, path func() string) *FileSource {

	once := func() {

		if strutil.IsBlank(path()) {
			msg := fmt.Sprintf("filepath is empty, not able to find value for key='%s'", key)
			slogx.Log().Debug(msg)
			return
		}

		err := koanfx.ReadFromYAMLFile(cfg, path())
		if err != nil {
			slogx.Log().Debug(err.Error())
		}

	}

	return newFileSourceWithOp(
		WithKey(key),
		WithPathFn(path),
		withOnce(once),
	)

}

func newFileSourceWithOp(ops ...FileSourceOp) *FileSource {
	src := &FileSource{}

	for _, op := range ops {
		op(src)
	}

	if strutil.IsEmpty(src.key) {
		panic("key is required")
	}

	return src
}
