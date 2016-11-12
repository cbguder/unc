package converter_test

import (
	"errors"

	"github.com/cbguder/v2e/converter"
	"github.com/cbguder/v2e/converter/converterfakes"
	"github.com/cbguder/v2e/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converter", func() {
	var (
		conv     converter.Converter
		importer *converterfakes.FakeImporter
		exporter *converterfakes.FakeExporter
	)

	BeforeEach(func() {
		importer = &converterfakes.FakeImporter{}
		exporter = &converterfakes.FakeExporter{}

		conv = converter.NewConverter(importer, exporter)
	})

	It("converts notes from one format to another", func() {
		notes := []models.Note{
			{Title: "My Special Note 1"},
			{Title: "My Special Note 2"},
		}

		importer.ImportReturns(notes, nil)

		conv.Convert("/from/path", "/to/path")

		Expect(importer.ImportCallCount()).To(Equal(1))

		inputPath := importer.ImportArgsForCall(0)
		Expect(inputPath).To(Equal("/from/path"))

		Expect(exporter.ExportCallCount()).To(Equal(1))

		outputPath, exportedNotes := exporter.ExportArgsForCall(0)
		Expect(outputPath).To(Equal("/to/path"))
		Expect(exportedNotes).To(Equal(notes))
	})

	It("returns an error when the import fails", func() {
		importError := errors.New("My Special Error")

		importer.ImportReturns([]models.Note{}, importError)

		err := conv.Convert("/from/path", "/to/path")

		Expect(err).To(Equal(importError))
	})

	It("returns an error when the export fails", func() {
		exportError := errors.New("My Special Error")

		exporter.ExportReturns(exportError)

		err := conv.Convert("/from/path", "/to/path")

		Expect(err).To(Equal(exportError))
	})
})
