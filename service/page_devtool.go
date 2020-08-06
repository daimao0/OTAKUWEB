package service

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"youblog/common/utils"
	"youblog/mapper"
)

func ListDevTools() string {
	redisConn := utils.InitRedis().Get()

	if err := redisConn.Err(); err != nil {
		log.Fatal("redis连接出现错误")
	}
	defer redisConn.Close()
	softJson, _ := redis.Bytes(redisConn.Do("get", "json:devtool"))
	if softJson == nil {
		softJson = getDevToolJsonFromDB()
		n, err := redisConn.Do("set", "json:devtool", softJson)
		if err != nil {
			log.Println("redis存入string类型失败", err)
		}
		if n == int64(1) {
			log.Println("redis存储string类型成功")
		}
	}
	return string(softJson)
}

func getDevToolJsonFromDB() []byte {
	listSoftWares := mapper.ListDevTools()
	bytes, err := json.Marshal(listSoftWares)
	if err != nil {
		log.Println("序列化json错误", err.Error())
	}
	return bytes
}
