package object

import (
	"bytes"
	"fmt"
	"strings"
)

type ObjectType = byte

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	_ byte = iota
	STRING_OBJ
	BUILTIN_OBJ
	CUSTOM_NAME_OBJ
	HTML_NODE_OBJ
	HASH_OBJ
	ERROR_OBJ
)

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type CustomName struct {
	Value string
}

func (cn *CustomName) Type() ObjectType { return CUSTOM_NAME_OBJ }
func (cn *CustomName) Inspect() string  { return cn.Value }

type HTMLNode struct {
	Tag        string
	Text       string
	Children   []HTMLNode
	CustomName string
}

func (hn *HTMLNode) Type() ObjectType { return HTML_NODE_OBJ }
func (hn *HTMLNode) Inspect() string  { return "" }

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  ObjectType
	Value string
}

func (s *String) HashKey() HashKey {
	return HashKey{Type: s.Type(), Value: s.Value}
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}

	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
