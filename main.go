package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"gopkg.in/gographics/imagick.v1/imagick"
)

//go:embed lgtm.gif
var lgtmImage []byte

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	var inputFile string
	var outputFile string

	flag.StringVar(&inputFile, "i", "", "input file")
	flag.StringVar(&outputFile, "o", "result", "output file")
	flag.Parse()

	if len(inputFile) == 0 {
		fmt.Printf("Please specify input file.\n")
		os.Exit(1)
	}

	if !exists(inputFile) {
		fmt.Printf("'%s' not exists.\n", inputFile)
		os.Exit(1)
	}
	extension := filepath.Ext(inputFile)

	imagick.Initialize()
	defer imagick.Terminate()

	source := imagick.NewMagickWand()
	lgtm := imagick.NewMagickWand()
	result := imagick.NewMagickWand()

	source.ReadImage(inputFile)
	lgtm.ReadImageBlob(lgtmImage)

	sourceWidth := source.GetImageWidth()
	sourceHeight := source.GetImageHeight()
	lgtm.ScaleImage(sourceWidth, sourceHeight)

	coalescedImages := source.CoalesceImages()
	source.Destroy()

	for i := 1; i <= int(coalescedImages.GetNumberImages()); i++ {
		coalescedImages.SetIteratorIndex(i)
		tmpImage := coalescedImages.GetImage()
		tmpImage.CompositeImage(lgtm, imagick.COMPOSITE_OP_OVER, 0, 0)
		result.AddImage(tmpImage)
		tmpImage.Destroy()
	}

	lgtm.Destroy()
	coalescedImages.Destroy()

	if err := result.WriteImages(outputFile+extension, true); err != nil {
		log.Fatal(err)
	}
	result.Destroy()
}
