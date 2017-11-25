package tvcptl

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	headRe  = regexp.MustCompile(`^(#{1,6})\s*(.*)$`)
	listRe  = regexp.MustCompile(`^\s*[\-\+\*]\s+(.*)$`)
	emptyRe = regexp.MustCompile(`^\s*$`)
)

func ParseLines() []Element {
	return nil
}

type Parser struct {
	Pos   int
	Lines []string
}

func (p *Parser) Paragraph() (*Element, bool) {
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

func (p *Parser) Head() (*Element, bool) {
	if p.End() || !p.Match(headRe) {
		return nil, false
	}
	m := headRe.FindStringSubmatch(p.Peek())
	return &Element{
		Name:     fmt.Sprintf("h%d", len(m[1])),
		Children: []Ast{&Inline{Value: m[2]}},
	}, true
}

func (p *Parser) End() bool {
	return p.Pos >= len(p.Lines)
}

func (p *Parser) Peek() string {
	return p.Lines[p.Pos]
}

func (p *Parser) Inc() {
	p.Pos++
}

func (p *Parser) Match(r *regexp.Regexp) bool {
	return !p.End() && r.MatchString(p.Peek())
}
