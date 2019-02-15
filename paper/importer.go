package paper

import (
	"bytes"
	"log"

	"github.com/cbguder/unc/models"
)

type Importer struct {
}

func (i Importer) Import(inputPath string) ([]models.Note, error) {
	cl := client{AuthToken: inputPath}

	docs, err := cl.ListDocuments()
	if err != nil {
		return nil, err
	}

	notes := []models.Note{}

	for _, docId := range docs.DocIds {
		res, body, err := cl.DownloadDocument(docId)
		if err != nil {
			return nil, err
		}

		log.Printf("Imported note: %s", res.Title)

		parts := bytes.SplitN(body, []byte("\n"), 2)
		body = bytes.TrimSpace(parts[1])
		body = append(body, byte('\n'))

		notes = append(notes, models.Note{
			Title: res.Title,
			Body:  string(body),
		})
	}

	return notes, nil
}
