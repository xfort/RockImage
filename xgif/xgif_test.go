package xgif

import (
	"testing"
)

func TestGetAllGifFrame(t *testing.T) {
	DecodeGifFrame("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol.gif", "")
}

func TestEncodeFileToGif(t *testing.T) {
	//	fileList := make([]string, 0)
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/0.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/0_0.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/1.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/1_0.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/2.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/2_0.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/3.jpg")
	//	fileList = append(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/3_0.jpg")
	//	EncodeFileToGif(fileList, "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/res.gif")
}

func TestCompressGif(t *testing.T) {
	//	CompressGif("/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol.gif", "/Users/Mac/Work/go/code/work/src/github.com/xfort/RockImage/xgif/gif/lol_256_100.gif", 50)
}
