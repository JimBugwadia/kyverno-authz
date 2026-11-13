package httpserver

import (
	"reflect"

	"cel.dev/cel-go/cel"
	"cel.dev/cel-go/ext"
)

type lib struct{}

func Lib() cel.EnvOption {
	// create the cel lib env option
	return cel.Lib(&lib{})
}

func (*lib) LibraryName() string {
	return "kyverno.httpserver"
}

func (c *lib) CompileOptions() []cel.EnvOption {
	return []cel.EnvOption{
		// register types
		ext.NativeTypes(
			reflect.TypeFor[HttpResponse](),
			ext.ParseStructTags(true),
		),
	}
}

func (*lib) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{}
}
