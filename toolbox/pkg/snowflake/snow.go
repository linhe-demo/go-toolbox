package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"toolbox/common"
)

func GetSnowId() (int64, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return common.Zero, fmt.Errorf("雪花算法实例化异常！")
	}
	id := node.Generate().Int64()
	return id, nil
}
