package xgif

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
)

var CurrentPath string
var CurrentOS string

//解析出gif的所有图片
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
	imgOption := &jpeg.Options{Quality: 70}

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

func EncodeImagesToGif(listimage []*image.Image, outpath string) {
	log.Println("EncodeImagesToGif")
	dir := path.Dir(outpath)
	err := os.MkdirAll(dir, 0666)

	if err != nil {
		log.Fatalln(err)
	}

	outGif := &gif.GIF{}
	for _, itemValue := range listimage {
		gifItemWriter := &bytes.Buffer{}

		gif.Encode(gifItemWriter, *itemValue, nil)

		gifImg, err := gif.Decode(gifItemWriter)
		if err != nil {
			log.Println("gif.Decode()", err)
		}
		outGif.Image = append(outGif.Image, gifImg.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
		//		gifItemWriter.Reset()
	}

	outFile, err := os.OpenFile(outpath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer outFile.Close()
	err = gif.EncodeAll(outFile, outGif)
	if err != nil {
		log.Fatalln("gif.EncodeAll", err)
	}
}

func EncodeFileToGif(filelist []string, outpath string) {
	imageList := make([]*image.Image, 0)

	for _, itemPic := range filelist {
		itemFile, err := os.Open(itemPic)
		if err != nil {
			log.Println("os.Open", err, itemPic)
			continue
		}
		defer itemFile.Close()

		itemImage, _, err := image.Decode(itemFile)
		if err != nil {
			log.Println("Decode", err, itemPic)
			continue
		}
		imageList = append(imageList, &itemImage)
	}
	EncodeImagesToGif(imageList, outpath)
}

//压缩gif
func CompressGif(sourceGif string, outGif string, quality int) {
	sourceFile, err := os.Open(sourceGif)
	if err != nil {
		log.Fatalln("Open", err, sourceGif)
	}
	defer sourceFile.Close()

	outFile, err := os.OpenFile(outGif, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("OpenFile", err, outGif)
	}
	defer outFile.Close()

	gifData, err := gif.DecodeAll(sourceFile)
	if err != nil {
		log.Fatalln("Decode", err, sourceGif)
	}
	outGifData := gif.GIF{}
	outGifData.Delay = gifData.Delay

	option := &jpeg.Options{quality}
	//	gifOption := &gif.Options{}
	//	gifOption.NumColors = 2

	for _, itemImage := range gifData.Image {
		itemJPG := &bytes.Buffer{}
		err = jpeg.Encode(itemJPG, itemImage, option)
		if err != nil {
			log.Println("jpeg.Encode()", err)
			continue
		}
		gifImage, _, err := image.Decode(itemJPG)
		if err != nil {
			log.Println("image.Decode", err)
		}
		//image.NewPaletted(gifImage.Bounds(),)
		itemGif := &bytes.Buffer{}
		err = gif.Encode(itemGif, gifImage, nil)
		if err != nil {
			log.Println("gif.Encode", err)
		}

		gifFrame, err := gif.Decode(itemGif)
		if err != nil {
			log.Println("gif.Decode", err)
			continue
		}
		outGifData.Image = append(outGifData.Image, gifFrame.(*image.Paletted))
	}

	err = gif.EncodeAll(outFile, &outGifData)
	if err != nil {
		log.Fatalln("gif.Encode", err)
	}
}

//使用gifsicle,压缩单张gif
//colorNum范围1-256，最佳128
func CompressByGifsicle(source string, out string, colorNum string) {
	if CurrentPath == "" {
		CurrentPath, _ = os.Getwd()
	}

	outDir := path.Dir(out)
	err := os.MkdirAll(outDir, os.ModePerm)

	if err != nil {
		log.Println("MkdirAll", outDir, err)
		return
	}
	gifSoft := path.Join(CurrentPath, "gifsicle."+runtime.GOOS)

	gifCmd := exec.Command(gifSoft, "--colors", colorNum, "-O3", source, "-o", out)
	resByte, err := gifCmd.CombinedOutput()
	gifName := path.Base(source)

	if err != nil {
		resStr := string(resByte)
		log.Println(gifName, err, resStr, gifSoft+"\n")
	}
}

//压缩文件夹下的所有gif文件
func CompressGifDir(fromDir string, outDir string, colorNum string) {

	fromFile, err := os.Open(fromDir)
	if err != nil {
		log.Println("os.Open()", err, fromDir)
		return
	}
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		log.Println("MkdirAll", err, outDir)
		return
	}
	fileList, err := fromFile.Readdir(0)
	if err != nil {
		log.Println("Readdir", err, fromDir)
		return
	}
	if fileList == nil || len(fileList) <= 0 {
		log.Println("empty dir ", fromFile)
		return
	}

	for _, itemFile := range fileList {
		if itemFile.IsDir() || itemFile.Size() <= 0 {
			continue
		}
		name := itemFile.Name()
		CompressByGifsicle(path.Join(fromDir, name), path.Join(outDir, name), colorNum)
	}
}
