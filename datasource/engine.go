package datasource

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"myweb/model"
	"sync"
)

var engine *xorm.Engine
var once sync.Once

func GetEngine() (*xorm.Engine, error) {
	var err error
	once.Do(func() {
		engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			log.Fatalf("连接数据库失败：%v", err)
		}
		//自动同步表结构
		if err = engine.Sync2(new(model.User), new(model.Project), new(model.TesetCase)); err != nil {
			log.Fatalf("同步表结构失败: %v", err)
		}
	})
	if engine == nil {
		return nil, errors.New("初始化数据库引擎失败")
	}
	//显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)

	return engine, nil
}
