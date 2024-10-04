package evaluator

import "github.com/nuflang/nuf/object"

var builtins = map[string]*object.Builtin{
	"section_title": {
		Fn: func(args ...object.Object) object.Object {
			switch arg := args[0].(type) {
			case *object.String:
				// FIXME: Support multiple heading level
				return &object.HTMLNode{Tag: "h1", Text: arg.Value, CustomName: "h1"}
			default:
				return newError("Argument to `section_title` not supported. Got `%d`.", args[0].Type())
			}
		},
	},
	"section": {
		Fn: func(args ...object.Object) object.Object {
			properties := object.Hash{}.Pairs
			customName := args[0].Inspect()

			if len(args) > 1 {
				properties = args[1].(*object.Hash).Pairs

				for _, hashPair := range properties {
					if hashPair.Key.Inspect() == "name" {
						customName = hashPair.Value.Inspect()
					}
				}
			}

			switch arg := args[0].(type) {
			case *object.String:
				switch arg.Value {
				case "main":
					return &object.HTMLNode{Tag: "main", CustomName: customName}
				case "site_navigation":
					return &object.HTMLNode{Tag: "nav", CustomName: customName}
				case "region":
					return &object.HTMLNode{Tag: "section", CustomName: customName}
				default:
					return newError("Unrecognized section `%s`.", arg.Value)
				}
			default:
				return newError("Argument to `section` not supported. Got `%d`.", args[0].Type())
			}
		},
	},
}
