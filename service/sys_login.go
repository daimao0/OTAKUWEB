package service

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"youblog/common/utils"
)

func CheckTokenFromRedis(tokenFromCookie string) bool {
	redisConn := utils.InitRedis().Get()
	defer redisConn.Close()
	userId, err := utils.GetUserIDFromToken(tokenFromCookie)
	if err != nil {
		log.Println(err)
	}
	token, _ := redis.String(redisConn.Do("get", "token:"+string(userId)))
	if tokenFromCookie == token {
		return true
	}
	return false
}
