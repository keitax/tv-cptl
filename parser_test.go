package tvcptl_test

import (
	. "github.com/keitax/tv-cptl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseBlocks()", func() {
	It("parses block markups", func() {
		got := ParseBlocks(`# h1
## h2

- a
- b

p`)
		Expect(got).To(Equal([]Ast{
			&Element{
				Name:     "h1",
				Children: []Ast{&Inline{Value: "h1"}},
			},
			&Element{
				Name:     "h2",
				Children: []Ast{&Inline{Value: "h2"}},
			},
			&Element{
				Name: "ul",
				Children: []Ast{
					&Element{
						Name:     "li",
						Children: []Ast{&Inline{Value: "a"}},
					},
					&Element{
						Name:     "li",
						Children: []Ast{&Inline{Value: "b"}},
					},
				},
			},
			&Element{
				Name: "p",
				Children: []Ast{
					&Inline{Value: "p"},
				},
			},
		}))
	})
})

var _ = Describe("BlockParser", func() {
	Describe("Paragraph()", func() {
		var p *BlockParser

		doParseParagraph := func(lines []string) (Ast, bool) {
			p = &BlockParser{Lines: lines}
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
		var p *BlockParser

		doParseHead := func(lines []string) (Ast, bool) {
			p = &BlockParser{Lines: lines}
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

	Describe("UList()", func() {
		var p *BlockParser

		doParseUList := func(lines []string) (Ast, bool) {
			p = &BlockParser{Lines: lines}
			return p.UList(-1)
		}

		It("parses flat items", func() {
			got, ok := doParseUList([]string{
				"+ hello0",
				"- hello1",
				"* hello2",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name: "ul",
				Children: []Ast{
					&Element{
						Name: "li",
						Children: []Ast{
							&Inline{Value: "hello0"},
						},
					},
					&Element{
						Name: "li",
						Children: []Ast{
							&Inline{Value: "hello1"},
						},
					},
					&Element{
						Name: "li",
						Children: []Ast{
							&Inline{Value: "hello2"},
						},
					},
				},
			}))
			Expect(p.Pos).To(Equal(3))
		})

		It("parses nested items", func() {
			got, ok := doParseUList([]string{
				"- hello0",
				"  - hello00",
				"- hello1",
				"  - hello10",
			})
			Expect(ok).To(BeTrue())
			Expect(got).To(Equal(&Element{
				Name: "ul",
				Children: []Ast{
					&Element{
						Name: "li",
						Children: []Ast{
							&Inline{Value: "hello0"},
							&Element{
								Name: "ul",
								Children: []Ast{
									&Element{
										Name: "li",
										Children: []Ast{
											&Inline{Value: "hello00"},
										},
									},
								},
							},
						},
					},
					&Element{
						Name: "li",
						Children: []Ast{
							&Inline{Value: "hello1"},
							&Element{
								Name: "ul",
								Children: []Ast{
									&Element{
										Name: "li",
										Children: []Ast{
											&Inline{Value: "hello10"},
										},
									},
								},
							},
						},
					},
				},
			}))
			Expect(p.Pos).To(Equal(4))
		})
	})
})
