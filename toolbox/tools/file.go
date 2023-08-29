package tools

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"toolbox/common"
)

func SaveFile(file multipart.File, header *multipart.FileHeader) string {
	// 保存上传的文件到 临时文件下
	_, err := os.ReadDir(common.TempFilePath)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(common.TempFilePath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	tempFile, err := ioutil.TempFile(common.TempFilePath, "*"+header.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	return tempFile.Name()
}

func DownloadFile(path string, w http.ResponseWriter, filename string) {
	w.Header().Add("Content-Type", "image/jpg")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\".jpg")
	file, err := os.Open(path)
	defer file.Close()
	stat, _ := file.Stat()
	size := stat.Size()
	data := make([]byte, size)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}
