package service

import (
	"errors"
	"myweb/datasource"
	"myweb/model"
)

type UserService struct {
}

// 用户登录
func (us *UserService) Login(username, password string) (*model.User, error) {
	user := &model.User{}
	engine, _ := datasource.GetEngine()

	has, err := engine.Where("user_name =?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("不存在该用户")
	}
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

// 用户注册
func (us *UserService) Register(user *model.User) error {
	if user.UserName == "" {
		return errors.New("用户名不能为空")
	}
	if user.UserName == "" || user.Password == "" {
		return errors.New("用户名或密码不能为空")

	}
	engine, _ := datasource.GetEngine()
	has, err := engine.Where("user_name =?", user.UserName).Exist(&model.User{})
	if err != nil {
		return err
	}
	if has {
		return errors.New("用户名已存在")
	}
	_, err = engine.Insert(user)
	return err
}

func (us *UserService) SyncTable() error {
	engine, _ := datasource.GetEngine()
	return engine.Sync2(new(model.User))
}
