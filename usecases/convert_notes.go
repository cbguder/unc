package usecases

import (
	"fmt"

	"github.com/cbguder/v2e/converter"
	"github.com/cbguder/v2e/evernote"
	"github.com/cbguder/v2e/vesper"
)

type ConvertNotesUseCase struct {
}

func (c ConvertNotesUseCase) ConvertNotes(fromFormat, toFormat, inputPath, outputPath string) error {
	importer, err := getImporter(fromFormat)
	if err != nil {
		return err
	}

	exporter, err := getExporter(toFormat)
	if err != nil {
		return err
	}

	conv := converter.NewConverter(importer, exporter)
	return conv.Convert(inputPath, outputPath)
}

func getImporter(format string) (converter.Importer, error) {
	if format == "vesper" {
		return vesper.Importer{}, nil
	}

	return nil, fmt.Errorf("unrecognized format: %s", format)
}

func getExporter(format string) (converter.Exporter, error) {
	if format == "evernote" {
		return evernote.Exporter{}, nil
	}

	return nil, fmt.Errorf("unrecognized format: %s", format)
}
