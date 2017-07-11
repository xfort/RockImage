package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/xfort/RockImage/xgif"
)

func main() {
	compressGif()
}

func compressGif() {

	currentPath, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	configFile, err := os.Open(path.Join(currentPath, "config.txt"))
	if err != nil {
		log.Fatalln(err, "无法读取config文件", currentPath)
	}
	defer configFile.Close()
	configReader := bufio.NewReader(configFile)

	var fromDirPath string
	var outDirPath string
	var colorNum string
	var maxSize string
	var watermarkpath string
	for {
		lineByte, _, err := configReader.ReadLine()
		if err == io.EOF {
			break
		}
		linStr := string(lineByte)
		if len(linStr) > 3 || !strings.HasPrefix(linStr, "//") {
			itemArray := strings.Split(linStr, "=")
			if itemArray != nil && len(itemArray) >= 2 {
				switch itemArray[0] {
				case "fromDir":
					fromDirPath = itemArray[1]
				case "outDir":
					outDirPath = itemArray[1]
				case "colorNum":
					colorNum = itemArray[1]
				case "maxSize":
					maxSize = itemArray[1]
				case "watermark":
					watermarkpath = itemArray[1]
				}
			}
		}
	}

	log.Println("读取配置\n", fromDirPath+"\n", outDirPath+"\n", colorNum+"\n", maxSize+"\n", watermarkpath)

	if fromDirPath == "" {
		log.Fatalln("fromDir 错误")
	}

	if outDirPath == "" {
		log.Fatalln("outDir 错误")
	}

	if colorNum == "" {
		colorNum = "256"
	}

	gifHandler := &xgif.GifHandler{}

	err = gifHandler.WatermarkCompressDir(fromDirPath, outDirPath, watermarkpath, colorNum, maxSize)
	if err != nil {
		log.Println(err.Error())
	}
	//xgif.CompressGifDir(fromDirPath, outDirPath, maxSize, colorNum)

	fmt.Println("任务结束\n", "文件保存路径="+outDirPath)
	time.Sleep(20 * time.Second)
}
