package xgif

import (
	"log"
	"os"

	"runtime"
	"os/exec"

	"errors"
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

//gif 压缩，加水印
type GifHandler struct {
	ImageMagickDirPath string
	GifsicleDirPath    string
}

//添加水印
func (gif *GifHandler) AddWatermark(gifsourcepath string, watermarkpath string, outgif string, converPath string) error {
	//convert test7.gif -coalesce -gravity southeast -geometry +10+10 null: xs_watermark.png -layers composite -layers optimize test7-watermarked.gif

	if converPath == "" {
		if gif.ImageMagickDirPath == "" {
			return errors.New("错误:ImageMagick文件夹路径为空")
		}
		if runtime.GOOS == "windows" {
			converPath = filepath.Join(gif.ImageMagickDirPath, "bin", "convert.exe")
		} else {
			return errors.New("错误:非windows系统必须指定ImageMagick的convert程序路径")
		}
	}
	log.Println("开始添加水印", gifsourcepath, watermarkpath)
	gifCmd := exec.Command(converPath, gifsourcepath, "-coalesce", "-gravity", "southeast", "-geometry", "+10+10",
		"null:", watermarkpath, "-layers", "composite", "-layers", "optimize", outgif)
	resByte, err := gifCmd.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintln("错误:添加水印失败", gifCmd.Args, string(resByte)))
	}
	return nil
}

//压缩,使用gifsicle
//targetSize 单位为M，例如2.0
//colornum （1-256）色值越小图片越小，质量越低，最佳值128，若targetSize>0,则忽略此参数
func (gif *GifHandler) Compress(sourcegif string, outgif string, colornum string, targetSize float32) error {
	log.Println("开始压缩图片", sourcegif)
	if gif.GifsicleDirPath == "" {
		return errors.New("错误:ImageMagick文件夹路径为空")
	}

	gifsiclePath := filepath.Join(gif.GifsicleDirPath, "gifsicle."+runtime.GOOS)

	if targetSize > 0 {
		return gif.CompressByTargetsize(sourcegif, outgif, targetSize, gifsiclePath)
	} else {
		return gif.CompressByColornum(sourcegif, outgif, colornum, gifsiclePath)
	}
	return nil
}

func (gif *GifHandler) CompressByColornum(sourcegif string, outgif string, colornum string, gifsiclepath string) error {
	gifCmd := exec.Command(gifsiclepath, "--colors", colornum, "-O3", sourcegif, "-o", outgif)
	resByte, err := gifCmd.CombinedOutput()

	if err != nil {
		return errors.New("错误：执行压缩命令失败," + err.Error() + fmt.Sprintln(sourcegif, string(resByte)))
	}
	return nil
}

//targtesize 单位为M,例如2.3M
func (gif *GifHandler) CompressByTargetsize(sourcegif string, outgif string, targtesize float32, gifsiclepath string) error {

	colorNum := 256
	metaColor := 256
	targtesize = targtesize * 1000 * 1000
	gifName := filepath.Base(sourcegif)

	sourceGifFile, err := os.Stat(sourcegif)
	if err != nil {
		return errors.New(fmt.Sprintln("错误：读取gif文件出错", gifName, err))
	}

	sourcesize := sourceGifFile.Size()

	if sourcesize < int64(targtesize) {
		fileData, err := ioutil.ReadFile(sourcegif)
		if err != nil {
			return errors.New(fmt.Sprintln("警告：源gif文件小于目标大小，不做处理,复制文件失败", gifName, err))
		}
		err = ioutil.WriteFile(outgif, fileData, 0644)
		if err != nil {
			return errors.New(fmt.Sprintln("警告：源gif文件小于目标大小，不做处理,复制文件时失败", gifName, err))
		}
		return errors.New(fmt.Sprintln("警告：源gif文件小于目标大小，不做处理", gifName))
	}

	for {
		gifCmd := exec.Command(gifsiclepath, "--colors", strconv.Itoa(colorNum), "-O3", sourcegif, "-o", outgif)
		resByte, err := gifCmd.CombinedOutput()

		if err != nil {
			resStr := string(resByte)
			return errors.New(fmt.Sprintln("错误：压缩出错", gifName, err, resStr))
		}

		outGifFile, err := os.Stat(outgif)
		if err != nil {
			return errors.New(fmt.Sprintln("错误：读取gif文件出错", gifName, err))
		}

		outsize := outGifFile.Size()
		diffCurrent := int64(targtesize) - outsize
		if diffCurrent >= 0 && colorNum >= 256 {
			break
		}

		metaColor = metaColor / 2

		if metaColor <= 0 {
			if diffCurrent < 0 {
				colorNum = colorNum - 1
				continue
			}
			outGifName := filepath.Base(outgif)
			outGifName = strings.Replace(outGifName, ".gif", "_"+strconv.Itoa(colorNum)+".gif", -1)
			err = os.Rename(outgif, filepath.Join(filepath.Dir(outgif), outGifName))

			if err != nil {
				log.Println("重命名文件失败", outgif)
			}
			break
		}

		if diffCurrent < 0 {
			colorNum = colorNum - metaColor
		} else {
			if colorNum >= 256 {
				colorNum = 0
			}
			colorNum = colorNum + metaColor
		}
	}
	return nil
}

//先添加水印，然后压缩
func (gif *GifHandler) AddWatermarkCompress(sourceGif string, outgif string, watermark string, colornum string,
	tarhetSize float32, convertPath string, gifsiclepath string) error {

	err := gif.AddWatermark(sourceGif, watermark, outgif, convertPath)
	if err != nil {
		return err
	}
	return gif.Compress(outgif, outgif, colornum, tarhetSize)
}

//添加水印，压缩文件夹下的所有gif文件
func (gif *GifHandler) WatermarkCompressDir(fromDir string, outDir string, watermark string, colornum string, targetSize string) error {
	var targetFileSize float32

	if targetSize != "" {
		targetFileSizeTmp, err := strconv.ParseFloat(targetSize, 32)
		if err != nil {
			return errors.New(fmt.Sprintln("错误:解析文件体积参数失败", targetSize, err))
		}
		targetFileSize = float32(targetFileSizeTmp)

	}

	fromFile, err := os.Open(fromDir)
	if err != nil {
		return errors.New(fmt.Sprintln("错误：打开gif源文件夹失败", err, fromDir))
	}
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintln("错误：创建gif输出文件夹失败", err, outDir))
	}
	fileList, err := fromFile.Readdir(0)
	if err != nil {

		return errors.New(fmt.Sprintln("错误：读取gif源文件夹内文件失败", err, fromDir))
	}
	if fileList == nil || len(fileList) <= 0 {
		return errors.New(fmt.Sprintln("错误：gif源文件夹内无文件", err, fromDir))
	}

	convertPath := ""
	gifsiclePath := ""
	if gif.ImageMagickDirPath == "" {
		execpath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return errors.New(fmt.Sprintln("错误：读取当前文件路径失败", err))
		}
		gif.ImageMagickDirPath = filepath.Join(execpath, "bin", "ImageMagick_windows")
	}
	errMsg := ""

	for _, itemFile := range fileList {
		if itemFile.IsDir() || itemFile.Size() <= 0 {
			continue
		}
		name := itemFile.Name()
		if strings.HasSuffix(name, "gif") {
			err = gif.AddWatermarkCompress(filepath.Join(fromDir, name), filepath.Join(outDir, name), watermark, colornum, targetFileSize, convertPath, gifsiclePath)
			if err != nil {
				errMsg = errMsg + err.Error()
				//log.Println(err)
				err = nil
			}
		}
	}
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
