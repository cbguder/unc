package converter

import "github.com/cbguder/v2e/models"

type Importer interface {
	Import(string) ([]models.Note, error)
}

