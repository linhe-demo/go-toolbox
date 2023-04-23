package tools

import (
	"fmt"
	"github.com/axgle/mahonia"
	"toolbox/common"
	"toolbox/exception"
)

func GetBaiduUrl(target string) (out string, err error) {
	if v, ok := common.BaiDuApiUrlMap[target]; !ok {
		return out, exception.NewCodeError(exception.ApiCode, fmt.Sprintf("图文识别暂未开放此功能：%s", target))
	} else {
		out = v
	}
	return out, nil
}

func GetBaiduParam(target string, data string) (out map[string][]string, err error) {
	switch target {
	case "text":
		out = map[string][]string{"image": {data}, "probability": {"false"}}
	case "form":
		out = map[string][]string{"image": {data}}
	default:
		err = exception.NewCodeError(exception.ApiCode, fmt.Sprintf("图文识别配置信息获取失败！：%s", target))
	}
	return out, err
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
