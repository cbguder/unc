package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cbguder/unc/helpers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	BeforeEach(func() {
		os.Setenv("TZ", "America/Los_Angeles")
	})

	AfterEach(func() {
		os.Unsetenv("TZ")
	})

	It("produces the expected output", func() {
		vesperNotesPath := helpers.GetFixturePath("vesper")

		outputDir, err := ioutil.TempDir("", "evernote")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		outputFile := filepath.Join(outputDir, "notes.enex")

		session := Run("-f", "vesper", "-t", "evernote", "-i", vesperNotesPath, "-o", outputFile)
		Eventually(session).Should(Exit(0))

		actualEnex, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		expectedEnex := helpers.ReadFixtureString("evernote.enex")
		Expect(string(actualEnex)).To(Equal(expectedEnex))
	})
})
