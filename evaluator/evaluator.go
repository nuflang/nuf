package evaluator

import (
	"fmt"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/object"
)

type Output struct {
	Value                  string
	CanAppendStringLiteral bool
}

func NewOutput() *Output {
	return &Output{
		Value:                  "",
		CanAppendStringLiteral: true,
	}
}

func (o *Output) appendToOutput(input string) {
	o.Value += input
}

func (o *Output) Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return o.evalStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return o.Eval(node.Expression, env)
	case *ast.StringLiteral:
		result := &object.String{Value: node.Value}

		if o.CanAppendStringLiteral {
			o.appendToOutput("<p>" + result.Inspect() + "</p>")
		}

		return result
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.CallExpression:
		o.CanAppendStringLiteral = false

		function := o.Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := o.evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		result := applyFunction(function, args)
		o.appendToOutput(result.Inspect())

		o.CanAppendStringLiteral = true

		return result
	}

	return nil
}

func (o *Output) evalStatements(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = o.Eval(statement, env)
	}

	return result
}

func (o *Output) evalExpressions(expressions []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range expressions {
		evaluated := o.Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("Identifier not found: " + node.Value)
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("Not a function: %d", fn.Type())
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, a...),
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}
