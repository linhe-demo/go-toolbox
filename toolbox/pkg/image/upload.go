package image

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"toolbox/common"
	"toolbox/internal/config"
	"toolbox/internal/models"
	"toolbox/internal/svc"
)

func DealImageFile(c config.Config, context context.Context, ctx *svc.ServiceContext, data string) {
	param := MqMessage{}
	err := json.Unmarshal([]byte(data), &param)
	if err != nil {
		log.Fatal(err)
		return
	}
	if param.Action == "add-image" {
		DealUploadImage(c, context, ctx, param)
	}
	if param.Action == "remove-image" {
		DealRemoveImage(c, param)
	}
}

func DealUploadImage(c config.Config, context context.Context, ctx *svc.ServiceContext, param MqMessage) {

	defer os.Remove(param.Path)
	//将图片压缩
	CompressionImage(param.Path, 0.5, strconv.FormatInt(param.Id, 10))
	newPath := fmt.Sprintf("%s%s.jpg", common.FilePath, strconv.FormatInt(param.Id, 10))

	//将压缩后的文件传输 七牛云
	qiniuPath, err := UploadToQiNiu(c, context, newPath, param.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(newPath)
	//将正常处理的图片信息保存
	info := models.LifeConfig{
		ConfigId:          param.ConfigId,
		ImgUrl:            qiniuPath,
		Text:              sql.NullString{Valid: true},
		Status:            2,
		HorizontalVersion: 0,
		CreateTime:        time.Now(),
		UpdateTime:        sql.NullTime{Valid: true},
	}
	_, err = ctx.LifeConfigModel.Insert(context, info)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Remove(param.Path)
	os.Remove(newPath)
}

func DealRemoveImage(c config.Config, param MqMessage) {
	mac := qbox.NewMac(c.QiNiuConf.AccessKey, c.QiNiuConf.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region = &storage.ZoneHuadongZheJiang2
	bucketManager := storage.NewBucketManager(mac, &cfg)
	// 删除七牛云上的图片
	keys := []string{
		param.Path,
	}
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(c.QiNiuConf.Bucket, key))
	}
	rets, err := bucketManager.Batch(deleteOps)
	if len(rets) == 0 {
		// 处理错误
		if e, ok := err.(*storage.ErrorInfo); ok {
			log.Printf("batch error, code:%s", e.Code)
		} else {
			log.Printf("batch error, %s", err)
		}
		return
	}
	// 返回 rets，先判断 rets 是否
	for _, ret := range rets {
		// 200 为成功
		if ret.Code == 200 {
			log.Printf("%d\n", ret.Code)
		} else {
			log.Printf("%s\n", ret.Data.Error)
		}
	}
}

func UploadToQiNiu(c config.Config, context context.Context, path string, name int64) (out string, err error) {

	putPolicy := storage.PutPolicy{
		Scope: c.QiNiuConf.Bucket,
	}
	mac := qbox.NewMac(c.QiNiuConf.AccessKey, c.QiNiuConf.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"name": "图片",
		},
	}
	data, err := ioutil.ReadFile(path) //read the content of file
	if err != nil {
		return out, err
	}
	dataLen := int64(len(data))

	key := fmt.Sprintf("images/life/%d", name)

	err = formUploader.Put(context, &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return out, err
	}
	log.Printf("七牛云返回处理结果: %s, hashKey: %s", ret.Key, ret.Hash)
	return ret.Key, nil
}
