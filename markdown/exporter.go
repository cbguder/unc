package markdown

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbguder/unc/models"
)

type Exporter struct {
}

func (e Exporter) Export(outputPath string, notes []models.Note) error {
	info, err := os.Stat(outputPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outputPath, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if !info.IsDir() {
		return errors.New("output path exists, but is not a directory")
	}

	for _, note := range notes {
		noteFilename := strings.Replace(note.Title, "/", "", -1)
		noteFilename = strings.Replace(noteFilename, ":", "", -1)
		noteFilename = fmt.Sprintf("%s.md", noteFilename)

		notePath := filepath.Join(outputPath, noteFilename)

		f, err := os.OpenFile(notePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}

		fmt.Fprintf(f, "# %s\n", note.Title)
		fmt.Fprintln(f)
		fmt.Fprintln(f, note.Body)

		formattedTags := make([]string, len(note.Tags))
		for i, tag := range note.Tags {
			formattedTags[i] = fmt.Sprintf("#%s", tag)
		}

		if len(formattedTags) > 0 {
			fmt.Fprintln(f)
			fmt.Fprintln(f, strings.Join(formattedTags, " "))
		}

		err = f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
