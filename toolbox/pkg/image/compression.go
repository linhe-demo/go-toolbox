package image

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"toolbox/common"
)

func CompressionImage(path string, compressionRatio float64, name string) {
	defer os.Remove(path)
	// 读取原始图片
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 压缩图片
	compressedImg := compressImage(img, compressionRatio) // 压缩到原来的一半大小
	imageName := fmt.Sprintf("%s.jpg", name)

	_, err = os.ReadDir(common.FilePath)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(common.FilePath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 保存压缩后的图片
	outputFile, err := os.Create(common.FilePath + imageName)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()
	jpeg.Encode(outputFile, compressedImg, &jpeg.Options{Quality: 80})
}

// 压缩图片
func compressImage(img image.Image, compression float64) image.Image {
	width := uint(float64(img.Bounds().Dx()) * 1)
	height := uint(float64(img.Bounds().Dy()) * 1)

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	return resizedImg
}
