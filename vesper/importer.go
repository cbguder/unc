package vesper

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/cbguder/v2e/models"
)

const timeLayout = "Jan 2, 2006, 3:04 PM"

func NewImporter() Importer {
	return Importer{}
}

type Importer struct {
}

func (i Importer) Import(inputPath string) ([]models.Note, error) {
	files, _ := ioutil.ReadDir(inputPath)

	notes := []models.Note{}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			path := filepath.Join(inputPath, file.Name())
			note := convertNote(path)
			notes = append(notes, note)
		}
	}

	return notes, nil
}

func convertNote(notePath string) models.Note {
	noteBytes, _ := ioutil.ReadFile(notePath)
	note := string(noteBytes)

	lines := strings.Split(note, "\n")

	nlines := len(lines)
	footerStart := nlines - 5

	title := lines[0]
	body := strings.TrimLeft(strings.Join(lines[1:footerStart-1], "\n"), "\n")
	tags := parseTags(lines[footerStart])
	created, _ := parseTime(lines[footerStart+2])
	modified, _ := parseTime(lines[footerStart+3])

	return models.Note{
		Title:    title,
		Body:     body,
		Tags:     tags,
		Created:  created,
		Modified: modified,
	}
}

func splitLine(line string) (string, string) {
	parts := strings.Split(line, ": ")
	return parts[0], parts[1]
}

func parseTags(line string) []string {
	_, rawTags := splitLine(line)
	tags := strings.Split(rawTags, ",")

	for idx, tag := range tags {
		tags[idx] = strings.TrimSpace(tag)
	}

	return tags
}

func parseTime(line string) (time.Time, error) {
	_, timeString := splitLine(line)

	return time.ParseInLocation(timeLayout, timeString, time.Local)
}
