package object

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

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
