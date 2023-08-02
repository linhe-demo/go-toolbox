package oauth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"toolbox/common"
	"toolbox/exception"
	"toolbox/internal/svc"
	"toolbox/pkg/watchdog"
	"toolbox/tools"
)

func GetBaiduOauthToken(ctx context.Context, l *svc.ServiceContext, w http.ResponseWriter, r *http.Request) string {
	var param []watchdog.LogInfo
	cmd := l.RedisClient.Get(ctx, "baiduAccessTokenKey")
	if err := cmd.Err(); err != nil {
		log.Print("get baidu access token err!")
	} else {
		param = append(param, watchdog.LogInfo{Action: "getRedisToken"})
		watchdog.Save(ctx, l, param, w, r)
		token, _ := cmd.Result()
		if len(token) > common.Zero {
			return token
		}
	}

	param = append(param, watchdog.LogInfo{Action: "getBaiduToken"})
	watchdog.Save(ctx, l, param, w, r)
	url := fmt.Sprintf("%soauth/2.0/token?client_id=%s&client_secret=%s&grant_type=client_credentials",
		l.Config.BaiduOauth.OauthUrl, l.Config.BaiduOauth.AppKey, l.Config.BaiduOauth.AppSecret)
	res, _ := tools.Post(url, "", map[string]string{"Content-Type": "application/json", "Accept": "application/json"})
	data := &AccessTokenRes{}
	err := json.Unmarshal(res, data)
	if err != nil {
		log.Print("JSON 数据解析失败！")
	}
	// 将数据存入redis
	cmd2 := l.RedisClient.Set(ctx, "baiduAccessTokenKey", data.AccessToken, time.Duration(int64(time.Second)*data.ExpiresIn))
	if err := cmd2.Err(); err != nil {
		log.Print("baidu access token save err!")
	}
	return data.AccessToken
}

func AnalysisPictureText(ctx context.Context, l *svc.ServiceContext, requestType string, file string, fileType int, w http.ResponseWriter, r *http.Request) (out *IdentifyPictureRes, err error) {
	// 获取请求地址
	method, err := tools.GetBaiduUrl(requestType)
	if err != nil {
		return out, err
	}
	var (
		image []byte
		sEnc  string
	)
	// 获取accessToken
	accessToken := GetBaiduOauthToken(ctx, l, w, r)
	if fileType == common.FileName {
		//读取文件
		wd, _ := os.Getwd()
		image, err = ioutil.ReadFile(wd + "/data/" + file)
		if err != nil {
			return out, exception.NewCodeError(exception.ApiCode, "get image fail!")
		}
		sEnc = base64.StdEncoding.EncodeToString(image)
	} else if fileType == common.BinaryStream {
		sEnc = file
	}

	url := fmt.Sprintf("%srest/2.0/ocr/v1/%s?access_token=%s", l.Config.BaiduOauth.OauthUrl, method, accessToken)
	param, err := tools.GetBaiduParam(requestType, sEnc)
	if err != nil {
		return out, err
	}
	res, _ := tools.PostForm(url, param)
	data := &IdentifyPictureRes{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return out, exception.NewCodeError(exception.ApiCode, "Analysis picture fail!")
	}
	return data, nil
}
