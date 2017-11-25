package tvcptl

const (
	ElementT = iota
	InlineT
)

type Ast interface {
	Type() int
}

type Element struct {
	Name     string
	Attrs    map[string]string
	Children []Ast
}

func (e *Element) Type() int {
	return ElementT
}

type Inline struct {
	Value string
}

func (e *Inline) Type() int {
	return InlineT
}
