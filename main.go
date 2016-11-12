package main

import (
	"flag"
	"os"

	"github.com/cbguder/v2e/converter"
	"github.com/cbguder/v2e/evernote"
	"github.com/cbguder/v2e/vesper"
)

func main() {
	inputPath := flag.String("i", "", "input file path")
	outputPath := flag.String("o", "", "output file path")
	flag.Parse()

	checkArgument(*outputPath)
	checkArgument(*inputPath)

	vesperImporter := vesper.NewImporter()
	evernoteExporter := evernote.NewExporter()

	conv := converter.NewConverter(vesperImporter, evernoteExporter)
	conv.Convert(*inputPath, *outputPath)
}

func checkArgument(arg string) {
	if arg == "" {
		flag.Usage()
		os.Exit(2)
	}
}
