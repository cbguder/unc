package evernote_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cbguder/v2e/evernote"
	"github.com/cbguder/v2e/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Evernote Exporter", func() {
	var (
		exporter evernote.Exporter
	)

	BeforeEach(func() {
		os.Setenv("TZ", "America/Los_Angeles")

		exporter = evernote.NewExporter()
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

		outputDir, err := ioutil.TempDir("", "evernote")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		outputFile := filepath.Join(outputDir, "notes.enex")

		err = exporter.Export(outputFile, notes)
		Expect(err).NotTo(HaveOccurred())

		outputBytes, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(outputBytes).To(ContainSubstring("<title>My Special Title 1</title>"))
		Expect(outputBytes).To(ContainSubstring(`<en-note><div>My Special Body 1</div></en-note>`))
		Expect(outputBytes).To(ContainSubstring("<created>20160420T232000Z</created>"))
		Expect(outputBytes).To(ContainSubstring("<updated>20170101T160000Z</updated>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Work</tag>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Life</tag>"))

		Expect(outputBytes).To(ContainSubstring("<title>My Special Title 2</title>"))
		Expect(outputBytes).To(ContainSubstring(`<en-note><div>My Special Body 2</div></en-note>`))
		Expect(outputBytes).To(ContainSubstring("<created>20120704T160500Z</created>"))
		Expect(outputBytes).To(ContainSubstring("<updated>20140201T233000Z</updated>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Harder</tag>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Better</tag>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Faster</tag>"))
		Expect(outputBytes).To(ContainSubstring("<tag>Stronger</tag>"))
	})

	It("exports notes with empty lines", func() {
		notes := []models.Note{
			{
				Title:    "My Special Title",
				Body:     "Line 1\n\nLine 2",
				Tags:     []string{},
				Created:  time.Now(),
				Modified: time.Now(),
			},
		}

		outputDir, err := ioutil.TempDir("", "evernote")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		outputFile := filepath.Join(outputDir, "notes.enex")

		err = exporter.Export(outputFile, notes)
		Expect(err).NotTo(HaveOccurred())

		outputBytes, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(outputBytes)).To(ContainSubstring(`<en-note><div>Line 1</div><div><br/></div><div>Line 2</div></en-note>`))
	})

	It("exports notes with special characters", func() {
		notes := []models.Note{
			{
				Title:    "My Special Title",
				Body:     "& < >",
				Tags:     []string{},
				Created:  time.Now(),
				Modified: time.Now(),
			},
		}

		outputDir, err := ioutil.TempDir("", "evernote")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		outputFile := filepath.Join(outputDir, "notes.enex")

		err = exporter.Export(outputFile, notes)
		Expect(err).NotTo(HaveOccurred())

		outputBytes, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(outputBytes).To(ContainSubstring(`<div>&amp; &lt; &gt;</div>`))
	})

	It("truncates the destination file", func() {
		notes := []models.Note{}

		outputDir, err := ioutil.TempDir("", "evernote")
		Expect(err).NotTo(HaveOccurred())
		defer os.RemoveAll(outputDir)

		data := ""
		for i := 0; i < 20; i++ {
			data += "Lorem Ipsum "
		}

		outputFile := filepath.Join(outputDir, "notes.enex")
		ioutil.WriteFile(outputFile, []byte(data), 0644)

		err = exporter.Export(outputFile, notes)
		Expect(err).NotTo(HaveOccurred())

		outputBytes, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(outputBytes).NotTo(ContainSubstring("Lorem Ipsum"))
	})
})
