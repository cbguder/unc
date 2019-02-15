package converter

import "github.com/cbguder/unc/models"

type Exporter interface {
	Export(string, []models.Note) error
}
