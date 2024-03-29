package image

type UploadImage struct {
	Path     string `json:"path"`
	Id       int64  `json:"id"`
	ConfigId int64  `json:"configId"`
}

type MqMessage struct {
	Path     string `json:"path"`
	Id       int64  `json:"id"`
	ConfigId string `json:"configId"`
	Action   string `json:"action"`
}
