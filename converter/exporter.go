package converter

import "github.com/cbguder/v2e/models"

type Exporter interface {
	Export(string, []models.Note) error
}
