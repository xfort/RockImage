package xgif

import (
	"testing"
)

func TestGetAllGifFrame(t *testing.T) {
	//	DecodeGifFrame("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/test_2.gif", "")
}

func TestEncodeFileToGif(t *testing.T) {
	//	fileList := make([]string, 0)
	//	dir := "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/test_2/"
	//	index := 0
	//	for index < 20 {
	//		indexStr := strconv.Itoa(index)
	//		fileList = append(fileList, dir+indexStr+"_70.jpg")
	//		index++
	//	}
	//	EncodeFileToGif(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/test_2_70.gif")
}

func TestCompressGif(t *testing.T) {
	//	CompressGif("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol.gif", "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol_90.gif", 50)
}

func TestCompressByGifsicle(t *testing.T) {
	//	CompressByGifsicle("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol.gif", "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/test/lol_one_10.gif", 10)
}

func TestCompressGifDir(t *testing.T) {
	CompressGifSize("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/", "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/test/", "2.0")
}
