package tvcptl_test

import (
	. "github.com/keitax/tv-cptl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Paragraph()", func() {
		var p *Parser

		doParseParagraph := func(lines []string) (*Element, bool) {
			p = &Parser{Lines: lines}
			return p.Paragraph()
		}

		It("parses a line", func() {
			got, ok := doParseParagraph([]string{
				"hello",
				"# bye",
				"# bye",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "p",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("doesn't parse lists", func() {
			got, ok := doParseParagraph([]string{
				"hello",
				"- bye",
				"- bye",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "p",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("doesn't parse heads", func() {
			got, ok := doParseParagraph([]string{
				"hello",
				"# bye",
				"# bye",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "p",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("doesn't parse empty lines", func() {
			got, ok := doParseParagraph([]string{
				"hello",
				" ",
				" ",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "p",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("parses no lines", func() {
			got, ok := doParseParagraph([]string{
				" ",
			})
			Expect(ok).To(BeFalse())
			Expect(got).To(BeNil())
			Expect(p.Pos).To(Equal(0))
		})
	})

	Describe("Head()", func() {
		var p *Parser

		doParseHead := func(lines []string) (*Element, bool) {
			p = &Parser{Lines: lines}
			return p.Head()
		}

		It("parses h1 line", func() {
			got, ok := doParseHead([]string{
				"# hello",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "h1",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("parses h6 line", func() {
			got, ok := doParseHead([]string{
				"###### hello",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "h6",
				Children: []Ast{&Inline{Value: "hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})

		It("doesn't parse h7 line", func() {
			got, ok := doParseHead([]string{
				"####### hello",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name:     "h6",
				Children: []Ast{&Inline{Value: "# hello"}},
			}))
			Expect(p.Pos).To(Equal(1))
		})
	})
})
