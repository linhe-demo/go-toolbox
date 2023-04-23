package oauth

type AccessTokenRes struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int64  `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}

type IdentifyPictureParam struct {
	Image string `json:"image"`
}

type IdentifyPictureRes struct {
	LogId               uint64             `json:"log_id"`
	WordsResultNum      uint32             `json:"words_result_num"`
	WordsResult         []Probability      `json:"words_result"`
	ParagraphsResultNum int                `json:"paragraphs_result_num"`
	ParagraphsResult    []ParagraphsResult `json:"paragraphs_result"`
}

type Probability struct {
	Words string `json:"words"`
}

type BaiDuAbility struct {
	Average  float64 `json:"average"`
	Min      float64 `json:"min"`
	Variance float64 `json:"variance"`
}

type ParagraphsResult struct {
	WordsResultIdx []int `json:"words_result_idx"`
}
