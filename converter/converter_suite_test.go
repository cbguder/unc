package converter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConverter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Converter Suite")
}
