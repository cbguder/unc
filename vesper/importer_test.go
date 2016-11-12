package vesper_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/cbguder/v2e/vesper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vesper Importer", func() {
	var (
		importer vesper.Importer
	)

	BeforeEach(func() {
		importer = vesper.NewImporter()
	})

	It("imports notes", func() {
		vesperNote1 := `My Special Note Title 1

My Special Note Body Line 1
My Special Note Body Line 2

Tags: Work, Life

Created: Apr 20, 2016, 4:20 PM
Modified: Jan 1, 2017, 8:00 AM
`

		vesperNote2 := `My Special Note Title 2

My Special Note Body Line 3
My Special Note Body Line 4

Tags: Harder, Better, Faster, Stronger

Created: May 9, 2015, 5:55 PM
Modified: Nov 8, 2016, 6:00 AM
`

		tempDir := writeNotesToTempDir(vesperNote1, vesperNote2)
		defer os.RemoveAll(tempDir)

		notes, err := importer.Import(tempDir)
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
		vesperNote := `My Special Note Title
My Special Note Body Line 1
My Special Note Body Line 2

Tags: Work, Life

Created: Apr 20, 2016, 4:20 PM
Modified: Jan 1, 2017, 8:00 AM
`

		tempDir := writeNotesToTempDir(vesperNote)
		defer os.RemoveAll(tempDir)

		notes, err := importer.Import(tempDir)
		Expect(err).NotTo(HaveOccurred())

		Expect(notes).To(HaveLen(1))

		Expect(notes[0].Title).To(Equal("My Special Note Title"))
		Expect(notes[0].Body).To(Equal("My Special Note Body Line 1\nMy Special Note Body Line 2"))
	})
})

func writeNotesToTempDir(notes ...string) string {
	tempDir, err := ioutil.TempDir("", "vesper")
	if err != nil {
		panic(err)
	}

	for i, note := range notes {
		filename := fmt.Sprintf("note_%d.txt", i)
		err = ioutil.WriteFile(filepath.Join(tempDir, filename), []byte(note), 0644)
		if err != nil {
			panic(err)
		}
	}

	return tempDir
}
