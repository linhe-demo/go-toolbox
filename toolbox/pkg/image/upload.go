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
	"strconv"
	"time"
	"toolbox/common"
	"toolbox/internal/config"
	"toolbox/internal/models"
	"toolbox/internal/svc"
)

func DealUploadImage(c config.Config, context context.Context, ctx *svc.ServiceContext, data string) {
	param := UploadImage{}
	err := json.Unmarshal([]byte(data), &param)
	if err != nil {
		log.Fatal(err)
		return
	}
	//将图片压缩
	CompressionImage(param.Path, 0.5, strconv.FormatInt(param.Id, 10))
	newPath := fmt.Sprintf("%s%s.jpg", common.FilePath, strconv.FormatInt(param.Id, 10))

	//将压缩后的文件传输 七牛云
	qiniuPath, err := UploadToQiNiu(c, context, newPath, param.Id)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	return ret.Key, nil
}
