package vesper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVesper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vesper Suite")
}
