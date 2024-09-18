package object

type ObjectType = byte

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	_ byte = iota
	STRING_OBJ
)

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}
