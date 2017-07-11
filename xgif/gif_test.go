package xgif

import (
	"testing"
	"log"
)

func TestGifHandler_Compress(t *testing.T) {
	gifHandler := &GifHandler{}

	sourceGif := "/Users/xs/Documents/out_water_test7.gif"
	outGif := "/Users/xs/Documents/out_test7.gif"
	gifsoftpath := "/Users/xs/work/go/code/work/src/github.com/xfort/RockImage/bin/gifsicle.darwin"

	err := gifHandler.CompressByTargetsize(sourceGif, outGif, 2.0, gifsoftpath)
	if err != nil {
		log.Fatalln(err)
	}
	err = gifHandler.CompressByColornum(sourceGif, outGif, "128", gifsoftpath)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestGifHandler_AddWatermark(t *testing.T) {
	gifHandler := &GifHandler{}
	gifHandler.ImageMagickDirPath = ""

	sourceGif := "/Users/xs/Documents/test7.gif"
	watermarkFile := "/Users/xs/Documents/qqkuaibao.png"
	outGif := "/Users/xs/Documents/out_water_test7.gif"

	err := gifHandler.AddWatermark(sourceGif, watermarkFile, outGif, "convert")
	if err != nil {
		log.Fatalln(err)
	}
}
