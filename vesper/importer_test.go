package vesper_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cbguder/unc/helpers"
	"github.com/cbguder/unc/vesper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vesper Importer", func() {
	var (
		importer vesper.Importer
	)

	BeforeEach(func() {
		importer = vesper.Importer{}
	})

	It("imports notes", func() {
		vesperNotesPath := helpers.GetFixturePath("vesper")

		notes, err := importer.Import(vesperNotesPath)
		Expect(err).NotTo(HaveOccurred())

		Expect(notes).To(HaveLen(2))

		Expect(notes[0].Title).To(Equal("My Special Note Title 1"))
		Expect(notes[0].Body).To(Equal("My Special Note Body Line 1\nMy Special Note Body Line 2"))
		Expect(notes[0].Tags).To(Equal([]string{"Work", "Life"}))

		expectedCreated := time.Date(2016, time.April, 20, 16, 20, 0, 0, time.Local)
		Expect(notes[0].Created).To(BeTemporally("==", expectedCreated))

		expectedModified := time.Date(2017, time.January, 1, 8, 0, 0, 0, time.Local)
		Expect(notes[0].Modified).To(BeTemporally("==", expectedModified))

		Expect(notes[1].Title).To(Equal("My Special Note Title 2"))
		Expect(notes[1].Body).To(Equal("My Special Note Body Line 3\nMy Special Note Body Line 4"))
		Expect(notes[1].Tags).To(Equal([]string{"Harder", "Better", "Faster", "Stronger"}))

		expectedCreated = time.Date(2015, time.May, 9, 17, 55, 0, 0, time.Local)
		Expect(notes[1].Created).To(BeTemporally("==", expectedCreated))

		expectedModified = time.Date(2016, time.November, 8, 6, 0, 0, 0, time.Local)
		Expect(notes[1].Modified).To(BeTemporally("==", expectedModified))
	})

	It("imports a note with no space between title and body", func() {
		vesperNote := helpers.ReadFixtureString("no_space_note.txt")

		noteFiles := map[string]string{
			"Note 1.txt": vesperNote,
		}

		tempDir := writeFilesToTempDir(noteFiles)
		defer os.RemoveAll(tempDir)

		notes, err := importer.Import(tempDir)
		Expect(err).NotTo(HaveOccurred())

		Expect(notes).To(HaveLen(1))

		Expect(notes[0].Title).To(Equal("My Special Note Title"))
		Expect(notes[0].Body).To(Equal("My Special Note Body Line 1\nMy Special Note Body Line 2"))
	})

	It("imports a note with no tags", func() {
		vesperNote := helpers.ReadFixtureString("no_tags_note.txt")

		noteFiles := map[string]string{
			"Note 1.txt": vesperNote,
		}

		tempDir := writeFilesToTempDir(noteFiles)
		defer os.RemoveAll(tempDir)

		notes, err := importer.Import(tempDir)
		Expect(err).NotTo(HaveOccurred())

		Expect(notes).To(HaveLen(1))

		Expect(notes[0].Tags).To(BeEmpty())
	})

	It("ignores files not ending in .txt", func() {
		noteFiles := map[string]string{
			"Not a Note.com": "",
		}

		tempDir := writeFilesToTempDir(noteFiles)
		defer os.RemoveAll(tempDir)

		notes, err := importer.Import(tempDir)
		Expect(err).NotTo(HaveOccurred())

		Expect(notes).To(BeEmpty())
	})
})

func writeFilesToTempDir(notes map[string]string) string {
	tempDir, err := ioutil.TempDir("", "vesper")
	Expect(err).NotTo(HaveOccurred())

	for filename, note := range notes {
		err = ioutil.WriteFile(filepath.Join(tempDir, filename), []byte(note), 0644)
		Expect(err).NotTo(HaveOccurred())
	}

	return tempDir
}
