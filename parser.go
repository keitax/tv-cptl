package tvcptl

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	headRe  = regexp.MustCompile(`^(#{1,6})\s*(.*)$`)
	listRe  = regexp.MustCompile(`^(\s{0,3})[\-\+\*]\s+(.*)$`)
	emptyRe = regexp.MustCompile(`^\s*$`)
)

func ParseLines() []Element {
	return nil
}

type BlockParser struct {
	Pos   int
	Lines []string
}

func (p *BlockParser) Paragraph() (*Element, bool) {
	ls := []string{}
	for !(p.End() || p.Match(headRe) || p.Match(listRe) || p.Match(emptyRe)) {
		ls = append(ls, p.Peek())
		p.Inc()
	}
	if len(ls) <= 0 {
		return nil, false
	}
	return &Element{
		Name:     "p",
		Children: []Ast{&Inline{Value: strings.Join(ls, "\n")}},
	}, true
}

func (p *BlockParser) Head() (*Element, bool) {
	if p.End() || !p.Match(headRe) {
		return nil, false
	}
	m := headRe.FindStringSubmatch(p.Peek())
	p.Inc()
	return &Element{
		Name:     fmt.Sprintf("h%d", len(m[1])),
		Children: []Ast{&Inline{Value: m[2]}},
	}, true
}

func (p *BlockParser) UList(indent int) (*Element, bool) {
	c := []Ast{}
	i := indent
	for {
		var it *Element
		var ok bool
		it, i, ok = p.UItem(i)
		if !ok {
			break
		}
		c = append(c, it)
	}
	if len(c) <= 0 {
		return nil, false
	}
	return &Element{
		Name:     "ul",
		Children: c,
	}, true
}

func (p *BlockParser) UItem(indent int) (*Element, int, bool) {
	if p.End() || !p.Match(listRe) {
		return nil, 0, false
	}
	c := []Ast{}
	m := listRe.FindStringSubmatch(p.Peek())
	i := len(m[1])
	if i < indent {
		return nil, 0, false
	}
	c = append(c, &Inline{Value: m[2]})
	p.Inc()
	for {
		l, ok := p.UList(i + 1)
		if !ok {
			break
		}
		c = append(c, l)
	}
	return &Element{
		Name:     "li",
		Children: c,
	}, i, true
}

func (p *BlockParser) End() bool {
	return p.Pos >= len(p.Lines)
}

func (p *BlockParser) Peek() string {
	return p.Lines[p.Pos]
}

func (p *BlockParser) Inc() {
	p.Pos++
}

func (p *BlockParser) Match(r *regexp.Regexp) bool {
	return !p.End() && r.MatchString(p.Peek())
}
