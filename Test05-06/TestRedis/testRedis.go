package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func testHash(conn redis.Conn) {
	defer conn.Close()

	//向redis中写入hash
	_, err := conn.Do("HSET", "user2", "name", "wusong")
	if err != nil {
		fmt.Println("hset err =", err)
	}

	hashName, err := redis.String(conn.Do("Hget", "user2", "name"))
	if err != nil {
		fmt.Println("hget err =", err)
	}

	fmt.Println(hashName)

	hashMap1, err := redis.StringMap(conn.Do("Hgetall", "user1"))
	for _, v := range hashMap1 {
		fmt.Println(v)
	}

}

//当程序启动时就初始化连接池
func init() {
	pool = &redis.Pool{
		//最大空闲连接数
		MaxIdle: 8,
		//表示和数据库的最大链接数，0表示没有限制
		MaxActive: 0,
		//最大空闲时间
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) { //初始化连接
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	////通过go向reds写入数据和读取数据
	////1.连接到redis
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println("redis.Dial err =", err)
	//	return
	//}
	//
	////defer conn.Close()
	//
	//////2.通过go向redis写入数据string[key - val]
	////_, err = conn.Do("Set", "name", "tom1")
	////if err != nil {
	////	fmt.Println("set err =", err)
	////}
	////
	//////fmt.Println("conn =", conn)
	//////3.通过go获取redis中的数据，获取到的数据类型是interface{},redis本身提供了类型转换的方法
	////data, err := redis.String(conn.Do("get", "name"))
	////fmt.Println(data)
	//
	//testHash(conn)

	//从连接池中取出一个连接
	conn := pool.Get()
	//连接池关闭之后就不能再取出连接了
	defer conn.Close()

}
