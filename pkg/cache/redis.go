package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"golang-gin-todolist/jwt"
	"log"
	"os"
	"reflect"
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
	_, err := Redis.Do("SET", td.AccessUid, strconv.Itoa(userid), "EX", int(at.Sub(now).Seconds()))
	if err != nil {
		return err
	}

	_, err = Redis.Do("SET", td.RefreshUid, strconv.Itoa(userid), "EX", int(rt.Sub(now).Seconds()))
	if err != nil {
		return err
	}
	return nil
}

type ConsumeFunc func(channel string, message []byte) error

func Publish(channel string, message string ) (int, error){
	n, err := redis.Int(Redis.Do("PUBLISH", channel, message))
	if err != nil {
		return 0, fmt.Errorf("redis publish %s %s, err: %v", channel, message, err)
	}
	return n, nil
}

func Subscribe(consume ConsumeFunc, channel ...string) error {
	psc := redis.PubSubConn{Conn: Redis}

	if err := psc.Subscribe(redis.Args{}.AddFlat(channel)...); err != nil {
		return err
	}

	done := make(chan error, 1)

	go func() {
		defer psc.Close()
		for  {
			fmt.Println(reflect.TypeOf(psc.Receive()))
			switch msg := psc.Receive().(type) {
				case error:
					done <- fmt.Errorf("redis pubsub receive error : %v", msg)
					return
				case redis.Message:
					fmt.Println("channel: ", msg.Channel)
					fmt.Println("msg: ", string(msg.Data))
					return
				case redis.Subscription:
					if msg.Count == 0 {
						done <- nil
						return
					}

			}
		}
	}()


	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	// Wait for goroutine to complete.
	return <-done
}