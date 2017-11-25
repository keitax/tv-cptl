package tvcptl_test

import (
	. "github.com/keitax/tv-cptl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Describe("Paragraph()", func() {
		doParseParagraph := func(lines []string) (*Element, bool) {
			p := &Parser{Lines: lines}
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
		})

		It("parses no lines", func() {
			got, ok := doParseParagraph([]string{
				" ",
			})
			Expect(ok).To(BeFalse())
			Expect(got).To(BeNil())
		})
	})
})
