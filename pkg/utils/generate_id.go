package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gofiber/fiber/v2/log"
)

var SnowflakeNode *snowflake.Node

func InitSnowflakeNode() error {
	node := os.Getenv("NODE")
	startTime := time.Date(2023, 12, 2, 10, 31, 32, 0, time.UTC)

	snowflake.Epoch = startTime.UnixMilli()
	nodeNum, err := strconv.Atoi(node)
	if err != nil {
		return fmt.Errorf("node is not a valid integer: %s", err)
	}
	snowflakeNode, err := snowflake.NewNode(int64(nodeNum))
	if err != nil {
		return fmt.Errorf("error initializing Snowflake node: %s", err)
	}
	SnowflakeNode = snowflakeNode
	return nil
}

func GenerateId() string {
	if SnowflakeNode == nil {
		log.Errorf("SnowflakeNode is not initialized")
		return ""
	}
	id := SnowflakeNode.Generate()
	return fmt.Sprint(id)
}