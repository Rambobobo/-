package model

type User struct {
	ID       int    `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(50)" json:"username"`
	Password string `xorm:"varchar(25)" json:"password"`
}
