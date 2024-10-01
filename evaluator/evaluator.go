package evaluator

import (
	"fmt"
	"slices"

	"github.com/nuflang/nuf/ast"
	"github.com/nuflang/nuf/object"
)

type Output struct {
	Node      map[string][]object.HTMLNode
	HTMLValue string
	NodeOrder []string
}

func NewOutput() *Output {
	return &Output{
		Node: map[string][]object.HTMLNode{},
		NodeOrder: []string{},
		HTMLValue: "",
	}
}

func (o *Output) GenerateHTML(nodes []object.HTMLNode, isParent bool) string {
	for _, node := range nodes {
		openTag := ""
		if isParent {
			openTag += "<" + node.Tag + ">"
		}

		closeTag := ""
		if isParent {
			closeTag += "</" + node.Tag + ">"
		}

		o.HTMLValue += openTag

		if node.Children != nil {
			o.GenerateHTML(node.Children, false)
		}

		text := node.Text
		if text != "" {
			openTag += "<" + node.Tag + ">"
			o.HTMLValue += openTag
			o.HTMLValue += text
			closeTag += "</" + node.Tag + ">"
		}

		o.HTMLValue += closeTag
	}

	return o.HTMLValue
}

func (o *Output) FlattenNodes(allNodes map[string][]object.HTMLNode) []object.HTMLNode {
	result := make([]object.HTMLNode, len(o.NodeOrder))

	for _, nodes := range allNodes {
		for _, node := range nodes {
			index := slices.Index(o.NodeOrder, node.CustomName)
			if index != -1 {
				result[index] = node
			}
		}
	}

	return result
}

func (o *Output) Eval(node ast.Node, env *object.Environment, skip bool) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return o.evalStatements(node.Statements, env)
	case *ast.ExpressionStatement:
		return o.Eval(node.Expression, env, false)
	case *ast.StringLiteral:
		result := &object.String{Value: node.Value}

		return result
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.CallExpression:
		function := o.Eval(node.Function, env, false)
		if isError(function) {
			return function
		}

		args := o.evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		result := applyFunction(function, args)

		if !skip {
			customName := result.(*object.HTMLNode).CustomName
			htmlNode := *result.(*object.HTMLNode)

			o.Node[customName] = append(o.Node[customName], htmlNode)
			o.NodeOrder = append(o.NodeOrder, customName)
		}

		return result
	case *ast.CustomNameExpression:
		htmlNodes := o.Node[node.Value]
		if len(htmlNodes) == 0 {
			return newError("Element with custom name %s not found", node.Value)
		}

		return &htmlNodes[0]
	case *ast.InfixExpression:
		left := o.Eval(node.Left, env, true)
		right := o.Eval(node.Right, env, true)

		result := o.evalInfixExpression(node.Operator, left, right)
		if result == nil {
			// FIXME: Better error message
			return newError("Couldn't evaluate infix expression")
		}

		customName := result.(*object.HTMLNode).CustomName
		htmlNode := *result.(*object.HTMLNode)

		if len(o.Node[customName]) == 0 {
			o.Node[customName] = make([]object.HTMLNode, 1)
		}

		o.Node[customName][0].Children = append(o.Node[customName][0].Children, htmlNode)

		return result
	case *ast.HashLiteral:
		return o.evalHashLiteral(node, env)
	}

	return nil
}

func (o *Output) evalStatements(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = o.Eval(statement, env, false)
	}

	return result
}

func (o *Output) evalExpressions(expressions []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range expressions {
		evaluated := o.Eval(e, env, false)
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

func (o *Output) evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case operator == "inside" && left.Type() == object.HTML_NODE_OBJ && right.Type() == object.HTML_NODE_OBJ:
		return o.evalInsideInfixExpression(left, right)
	default:
		return nil
	}
}

func (o *Output) evalInsideInfixExpression(left, right object.Object) object.Object {
	return &object.HTMLNode{
		Tag: right.(*object.HTMLNode).Tag,
		Children: []object.HTMLNode{
			{
				Tag:        left.(*object.HTMLNode).Tag,
				Text:       left.(*object.HTMLNode).Text,
				CustomName: left.(*object.HTMLNode).CustomName,
			},
		},
		CustomName: right.(*object.HTMLNode).CustomName,
	}
}

func (o *Output) evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := o.Eval(keyNode, env, false)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("Unusable as hash key: %s", key.Type())
		}

		value := o.Eval(valueNode, env, false)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{
		Pairs: pairs,
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
