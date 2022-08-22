package gedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var redisClient  *redis.Client

var redisClient_Once sync.Once


// 使用连接池是保护mysql和redis，但是并不增加性能
func Redis()  *redis.Client {
	redisClient_Once.Do(func() {  // 实现单例模式
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "192.168.19.146:6379",
			Password: "", // 密码
			DB:       0,  // 数据库
			PoolSize: 20, // 连接池大小
		})
		pong,err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("connect error:%s",err))
		}
		log.Println(pong)
	})
	return redisClient
}