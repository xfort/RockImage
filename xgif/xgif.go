package xgif

import (
	"image/gif"
	"image/jpeg"
	"log"
	"os"
	"path"
	"strconv"
)

func DecodeGifFrame(gifPath string, outDir string) {
	giffile, err := os.Open(gifPath)
	if err != nil {
		log.Fatalln(err, "getAllFrame", gifPath)
	}
	//	log.Println(giffile.Name(), path.Dir(filepath), path.Base(filepath))
	if outDir == "" {
		//		log.Println(filepath.Base(gifPath), filepath.Ext(gifPath), filepath.Clean(gifPath))
		filename := path.Base(gifPath)
		outDir = path.Join(path.Dir(gifPath), filename[:len(filename)-len(path.Ext(filename))])
	}
	err = os.MkdirAll(outDir, os.ModePerm)

	if err != nil {
		log.Fatalln(err)
	}
	gifData, err := gif.DecodeAll(giffile)
	if err != nil {
		log.Fatalln(err, "getAllFrame_DecodeAll")
	}
	//	dir, _ := os.Getwd()
	imgOption := &jpeg.Options{Quality: 50}

	for index, item := range gifData.Image {
		itempath := path.Join(outDir, strconv.Itoa(index)+"_"+strconv.Itoa(imgOption.Quality)+".jpg")
		//		os.MkdirAll(itempath, 0666)
		itemFile, err := os.OpenFile(itempath, os.O_RDWR|os.O_CREATE, 0666)

		defer itemFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
		err = jpeg.Encode(itemFile, item, imgOption)
		if err != nil {
			log.Fatalln("getAllFrame_item", err)
		}
	}
}

//func EncodeImagesToGif(listimage []*image.Image, outpath string) {
//	gif.EncodeAll()
//}
