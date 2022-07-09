package internal

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// loadImage 载入图片
func loadImage(photo string) (img image.Image, err error) {
	// 打开图片文件
	file, err := os.Open(photo)
	// 检查错误
	if err != nil {
		return
	}
	// 关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 图片解码
	img, _, err = image.Decode(file)

	return
}

// clipImage 剪切图片
func clipImage(srcFile, newFile string, x, y, w, h int) error {
	// 载入图片
	src, err := loadImage(srcFile)
	// 检查错误
	if err != nil {
		return err
	}

	// 剪切图片
	img := src.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x, y, x+w, y+h))

	// 保存图片
	saveErr := saveImage(newFile, img)
	// 检查错误
	if saveErr != nil {
		return err
	}

	return nil
}

// saveImage 保存图片
func saveImage(path string, img image.Image) error {
	// 新建并打开文件
	f, err := os.OpenFile(path, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	// 检查错误
	if err != nil {
		return err
	}
	// 关闭
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// 获取文件后缀
	ext := filepath.Ext(path)

	// 如果是jpeg类型
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		// jpeg图片编码
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") { // png类型
		// png图片编码
		err = png.Encode(f, img)
	}

	return err
}

// transformImage 将指定图片进行裁剪，并返回错误信息
func transformImage(inputPath, outputPath string) error {
	var width, height, x int
	img, err := loadImage(inputPath)
	if err != nil {
		return err
	}
	// 获取图片边界
	b := img.Bounds()
	// 获取图片宽度并设置为一半
	width = b.Max.X / 2
	// 获取图片高度
	height = b.Max.Y
	// 将x坐标设置为0
	x = width

	return clipImage(inputPath, outputPath, x, 0, width, height)
}
