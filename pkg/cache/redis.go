package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golang-gin-todolist/jwt"
	"log"
	"os"
	"strconv"
	"time"
)

var Redis redis.Conn

func init() {
	var err error
	ip := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	Redis, err = redis.Dial("tcp", ip)
	if err != nil {
		log.Fatal("redis can not connect")
	}
}

// 設定 token 到快取
func SetTokenCache(userid int, td *jwt.TokenDetails) error {
	at := time.Unix(td.AtExpiredAt, 0) // convert unix to utc
	rt := time.Unix(td.RfExpiredAt, 0)
	now := time.Now()

	// 設定 access token 到快取
	_, err := Redis.Do("SET", td.AccessUid, strconv.Itoa(userid), "EX", at.Sub(now))
	if err != nil {
		return err
	}

	_, err = Redis.Do("SET", td.RefreshUid, strconv.Itoa(userid), "EX", rt.Sub(now))
	if err != nil {
		return err
	}
	return nil
}