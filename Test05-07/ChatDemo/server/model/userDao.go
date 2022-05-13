package model

import (
	"ChatDemo/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// MyUserDao 在服务器启动后，就初始化一个userDao实例
// 将其作为全局变量，在需要和redis操作时，直接使用
var (
	MyUserDao *UserDao
)

// UserDao 定义一个UserDao 的结构体
// 完成对user结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式，创建一个userDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1.根据用户ID返回一个user实例 + err
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	//通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			//表示users没有找到对应的id
			err = ERROR_USER_NOTEXISTES
		}
		return
	}

	//把res反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal fail =", err)
	}

	return
}

// 完成登录的校验login
//1. login完成对用户的验证
//2. 如果用户id和pwd正确，则返回一个正确的user实例
//3. 如果id和pwd错误，则返回对应的错误信息

func (ud *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	//先从redis连接池中取出一个连接
	conn := ud.pool.Get()
	defer conn.Close()

	user, err = ud.getUserById(conn, userID)
	if err != nil {
		fmt.Println("getUserById fail =", err)
		return
	}

	//此时用户存在，开始匹配密码
	if user.UserPwd == userPwd {

	} else {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (ud *UserDao) Register(user *message.User) (err error) {
	//先从redis连接池中取出一个连接
	conn := ud.pool.Get()
	defer conn.Close()

	_, err = ud.getUserById(conn, user.UserID)
	if err == nil {
		err = ERROR_USER_EXISTES //用户已存在
		return
	}

	//此用户还未注册，开始注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//写入redis
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("conn.DO fail =", err)
		return
	}

	return
}
