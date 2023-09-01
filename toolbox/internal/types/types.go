// Code generated by goctl. DO NOT EDIT.
package types

type MatchRequest struct {
	UserId  int64 `json:"userId"`
	Rank    int64 `json:"rank"`
	NeedNum int64 `json:"needNum"`
}

type MatchResponse struct {
	RoomId int64 `json:"roomId"`
}

type OcrRequest struct {
	File     string `json:"file"`
	Type     string `json:"type"`
	FileType int    `json:"fileType"`
}

type OcrResponse struct {
	Result interface{} `json:"result"`
}

type LogRequest struct {
	Action     string `json:"action"`
	ActionUser string `json:"actionUser"`
	Ip         string `json:"ip"`
}

type LogResponse struct {
	Result interface{} `json:"result"`
}

type CompressionRequest struct {
	Nmae string `json:"nmae"`
	Type int64  `json:"type,omitempty"`
}

type CompressionResponse struct {
	Path string `json:"path"`
}
