package tvcptl_test

import (
	. "github.com/keitax/tv-cptl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RenderBlocks()", func() {
	It("renders ast blocks", func() {
		got := RenderBlocks(nil)
		Expect(got).To(Equal("hello"))
	})
})
