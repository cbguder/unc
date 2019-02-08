package main

import (
	"flag"
	"log"
	"os"

	"github.com/cbguder/v2e/usecases"
)

func main() {
	fromFormat := flag.String("f", "", "from format")
	toFormat := flag.String("t", "", "to format")
	inputPath := flag.String("i", "", "input file path")
	outputPath := flag.String("o", "", "output file path")
	flag.Parse()

	checkArgument(*fromFormat)
	checkArgument(*toFormat)
	checkArgument(*outputPath)
	checkArgument(*inputPath)

	useCase := usecases.ConvertNotesUseCase{}
	err := useCase.ConvertNotes(*fromFormat, *toFormat, *inputPath, *outputPath)
	if err != nil {
		log.Fatalln(err)
	}
}

func checkArgument(arg string) {
	if arg == "" {
		flag.Usage()
		os.Exit(2)
	}
}
