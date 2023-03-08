package model

import "time"

type TesetCase struct {
	ID           int64     `xorm:"pk autoincr" json:"id"`
	TestCaseName string    `xorm:"varchar(255) " json:"project"` //测试用例名称
	Description  string    `xorm:"varchar(1024)" json:"des"`     //测试用例描述
	CreatedAt    time.Time `xorm:"created_at" json:"created"`    //创建时间
	UpdatedAt    time.Time `xorm:"updated_at" json:"updated"`    //更新时间
}
