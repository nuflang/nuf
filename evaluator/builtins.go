package evaluator

import "github.com/nuflang/nuf/object"

var builtins = map[string]*object.Builtin{
	"section_title": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Expected `1` argument. Got `%d`.", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				// FIXME: Support multiple heading level
				return &object.String{Value: "<h1>" + arg.Value + "</h1>"}
			default:
				return newError("Argument to `section_title` not supported. Got `%d`.", args[0].Type())
			}
		},
	},
}
