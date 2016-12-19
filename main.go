package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getExtensionFromFileName(filename string) string {
	splitedString := strings.Split(filename, ".")

	if len(splitedString) > 1 {
		return splitedString[1]
	}
	return ""
}

func main() {
	var inputFile string
	var outputFile string
	var extension string

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
	extension = getExtensionFromFileName(inputFile)

	imagick.Initialize()
	defer imagick.Terminate()

	source := imagick.NewMagickWand()
	lgtm := imagick.NewMagickWand()
	result := imagick.NewMagickWand()

	source.ReadImage(inputFile)
	lgtm.ReadImage("lgtm.gif")

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

	result.WriteImages(outputFile+"."+extension, true)
	result.Destroy()
}
