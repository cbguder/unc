package paper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPaper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paper Suite")
}
