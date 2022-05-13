package model

import "errors"

var (
	ERROR_USER_NOTEXISTES = errors.New("用户不存在...")
	ERROR_USER_EXISTES    = errors.New("用户已存在...")
	ERROR_USER_PWD        = errors.New("密码出错...")
)
