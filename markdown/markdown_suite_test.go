package markdown_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMarkdown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Markdown Suite")
}
