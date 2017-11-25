package tvcptl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTvCptl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TvCptl Suite")
}
