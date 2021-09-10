package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"os"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
	//刚才设置的环境变量 想要生效 我们必须得重启goland
}


func main() {
	r := gin.Default()


	rdb:= redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "ABC123", // no password set
		DB:       0,  // use default DB
	})






	r.GET("/ping", func(c *gin.Context) {

		val1,err:=rdb.Incr(context.Background(),"number").Result()

		if err!=nil{
			panic(err)
		}
		fmt.Println(val1)
		val, err := rdb.Get(context.Background(), "number").Result()
		if err!=nil{
			panic(err)
		}


		name, err := os.Hostname()

		if err!=nil{
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": "pong1",
			"number": val,
			"host":name,
			"pass":GetEnvInfo("REDIS_PASSWORD"),
		})
	})
	r.Run(":80")
}
