package service

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"youblog/common/utils"
	"youblog/mapper"
)

func ListSoftWares() string {
	redisConn := utils.InitRedis().Get()

	if err := redisConn.Err(); err != nil {
		log.Fatal("redis连接出现错误")
	}
	defer redisConn.Close()
	softJson, _ := redis.Bytes(redisConn.Do("get", "json:soft"))
	if softJson == nil {
		softJson = getSoftJsonFromDB()
		n, err := redisConn.Do("set", "json:soft", softJson)
		if err != nil {
			log.Println("redis存入string类型失败", err)
		}
		if n == int64(1) {
			log.Println("redis存储string类型成功")
		}
	}

	return string(softJson)
}

func getSoftJsonFromDB() []byte {
	listSoftWares := mapper.ListSoftWares()
	bytes, err := json.Marshal(listSoftWares)
	if err != nil {
		log.Println("序列化json错误", err.Error())
	}
	return bytes
}
