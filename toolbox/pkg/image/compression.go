package image

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"strings"
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
	jpeg.Encode(outputFile, compressedImg, &jpeg.Options{Quality: 90})
}

// 压缩图片
func compressImage(img image.Image, compression float64) image.Image {
	width := uint(float64(img.Bounds().Dx()))
	height := uint(float64(img.Bounds().Dy()))

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	return resizedImg
}

// TransferToPdf 图片转PDF
func TransferToPdf(filepath string, name string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.Image(filepath, 10, 10, pdf.GetPageSizeStr("A4").Wd-20, 0, false, "", 0, "")
	err := pdf.OutputFileAndClose(fmt.Sprintf("%s%s.pdf", common.FilePath, name))
	if err != nil {
		log.Fatalf("output failed,err:%s", err)
	}
}

// RemoveWatermark 去除水印
func RemoveWatermark(filepath string, name string) string {
	tmpSlice := strings.Split(filepath, ".")
	suffix := tmpSlice[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	//判断水印位置
	bounds := img.Bounds()
	x := bounds.Dx() - 100
	y := bounds.Dy() - 100
	fmt.Println(x)
	fmt.Println(y)
	//去除水印
	img = imaging.Crop(img, image.Rect(0, 0, x, y))
	newFilePath := fmt.Sprintf("%s%s.%s", common.FilePath, name, suffix)
	//保存处理后的图片
	err = imaging.Save(img, newFilePath)
	return fmt.Sprintf("%s%s.%s", common.DownloadFilePath, name, suffix)
}
