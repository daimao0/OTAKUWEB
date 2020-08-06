package test

import (
	"fmt"
	"log"
	"testing"
	"youblog/common/utils"
	"youblog/mapper"
	"youblog/service"
)

func TestMysql(t *testing.T) {
	listSoftWares := mapper.ListSoftWares()
	fmt.Println(listSoftWares)
}

func TestService(t *testing.T) {
	service.ListSoftWares()
}

func TestSetRedisString(t *testing.T) {
	conn := utils.InitRedis().Get()
	_, err := conn.Do("set", "name", "张三")
	if err != nil {
		log.Fatal(err)
	}
}
