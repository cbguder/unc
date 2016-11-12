package evernote_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEvernote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evernote Suite")
}
