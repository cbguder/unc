package usecases

import (
	"fmt"

	"github.com/cbguder/unc/converter"
	"github.com/cbguder/unc/evernote"
	"github.com/cbguder/unc/markdown"
	"github.com/cbguder/unc/paper"
	"github.com/cbguder/unc/vesper"
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
	if format == "paper" {
		return paper.Importer{}, nil
	} else if format == "vesper" {
		return vesper.Importer{}, nil
	}

	return nil, fmt.Errorf("unrecognized format: %s", format)
}

func getExporter(format string) (converter.Exporter, error) {
	if format == "evernote" {
		return evernote.Exporter{}, nil
	} else if format == "markdown" {
		return markdown.Exporter{}, nil
	}

	return nil, fmt.Errorf("unrecognized format: %s", format)
}
