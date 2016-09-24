package main

import (
	"bufio"
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
				}
			}
		}
	}

	log.Println(fromDirPath+"\n", outDirPath+"\n", colorNum)

	if fromDirPath == "" {
		log.Fatalln("fromDir 错误")
	}

	if outDirPath == "" {
		log.Fatalln("outDir 错误")
	}

	if colorNum == "" {
		colorNum = "256"
	}

	xgif.CompressGifDir(fromDirPath, outDirPath, colorNum)

	log.Println("任务结束\n", "保存路径="+outDirPath)
	time.Sleep(5 * time.Minute)

}
