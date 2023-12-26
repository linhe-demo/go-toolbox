package common

const (
	Zero               = 0
	One                = 1
	DefaultHashArrSize = 16 // 默认hash数组大小

	TimeFormat       = "2006-01-02 15:04:05"
	WaitSecond       = 10
	Success          = "success"
	FilePath         = "/var/www/html/"
	UploadFilePath   = "/home/static/upload"
	TempFilePath     = "/home/static/temp/"
	DownloadFilePath = "http://150.158.82.218/images/"
)

const (
	DefaultSecond    = 1.5 //匹配时间基数
	DefaultWaiteTime = 5   //改变匹配范围所需要等待的时间
	DefaultRankRange = 5   //增加匹配rank分范围

	MaxWaitTime = 3600000 //最大等待时间
)

const (
	FileName     = 1
	BinaryStream = 2
)
