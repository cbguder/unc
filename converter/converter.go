package converter

type Converter struct {
	Importer Importer
	Exporter Exporter
}

func NewConverter(importer Importer, exporter Exporter) Converter {
	return Converter{
		Importer: importer,
		Exporter: exporter,
	}
}

func (c Converter) Convert(inputPath, outputPath string) error {
	notes, err := c.Importer.Import(inputPath)

	if err != nil {
		return err
	}

	return c.Exporter.Export(outputPath, notes)
}
