package converter

import "github.com/cbguder/unc/models"

type Importer interface {
	Import(string) ([]models.Note, error)
}
