package markdown_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cbguder/v2e/markdown"
	"github.com/cbguder/v2e/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Markdown Exporter", func() {
	var (
		exporter markdown.Exporter
	)

	BeforeEach(func() {
		os.Setenv("TZ", "America/Los_Angeles")

		exporter = markdown.Exporter{}
	})

	AfterEach(func() {
		os.Unsetenv("TZ")
	})

	It("exports notes", func() {
		notes := []models.Note{
			{
				Title:    "My Special Title 1",
				Body:     "My Special Body 1",
				Tags:     []string{"Work", "Life"},
				Created:  time.Date(2016, time.April, 20, 16, 20, 0, 0, time.Local),
				Modified: time.Date(2017, time.January, 1, 8, 0, 0, 0, time.Local),
			},
			{
				Title:    "My Special Title 2",
				Body:     "My Special Body 2",
				Tags:     []string{"Harder", "Better", "Faster", "Stronger"},
				Created:  time.Date(2012, time.July, 4, 9, 5, 0, 0, time.Local),
				Modified: time.Date(2014, time.February, 1, 15, 30, 0, 0, time.Local),
			},
		}

		outputDir, err := ioutil.TempDir("", "markdown")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		err = exporter.Export(outputDir, notes)
		Expect(err).NotTo(HaveOccurred())

		outputPath1 := filepath.Join(outputDir, "My Special Title 1.md")
		outputBytes1, err := ioutil.ReadFile(outputPath1)
		Expect(err).NotTo(HaveOccurred())

		Expect(outputBytes1).To(ContainSubstring("# My Special Title 1"))
		Expect(outputBytes1).To(ContainSubstring("My Special Body 1"))
		Expect(outputBytes1).To(ContainSubstring("#Work #Life"))

		outputPath2 := filepath.Join(outputDir, "My Special Title 2.md")
		outputBytes2, err := ioutil.ReadFile(outputPath2)

		Expect(outputBytes2).To(ContainSubstring("# My Special Title 2"))
		Expect(outputBytes2).To(ContainSubstring("My Special Body 2"))
		Expect(outputBytes2).To(ContainSubstring("#Harder #Better #Faster #Stronger"))
	})
})
