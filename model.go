package main

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type EmailObject struct {
	Id   int `orm:"pk;auto"`
	From string
	Rcpt string
	Data string    `orm:"type(text)"`
	Date time.Time `orm:"auto_now_add"`
	Ham  bool      `orm:"default(true)"`
}

func init() {
	orm.RegisterModel(new(EmailObject))
}
