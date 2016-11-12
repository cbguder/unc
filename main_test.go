package main_test

import (
	"github.com/cbguder/v2e/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("produces the expected output", func() {
		vesperNotesPath := helpers.GetFixturePath("vesper")

		enexFile := helpers.CreateTempFile("enex")
		defer helpers.DiscardTempFile(enexFile)

		session := Run("-i", vesperNotesPath, "-o", enexFile.Name())
		Eventually(session).Should(Exit(0))

		actualEnex := helpers.ReadFileString(enexFile.Name())
		expectedEnex := helpers.ReadFixtureString("evernote.enex")
		Expect(actualEnex).To(Equal(expectedEnex))
	})
})