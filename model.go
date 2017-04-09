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
type TrainSpamResult struct {
	Id         int `orm:"pk;auto"`
	Key        string
	Pro        float64
	LastUptate time.Time `orm:"auto_now"`
}
type TrainHamResult struct {
	Id         int `orm:"pk;auto"`
	Key        string
	Pro        float64
	LastUptate time.Time `orm:"auto_now"`
}
type AllPro struct {
	Id         int `orm:"pk;auto"`
	Key        string
	Pros       float64
	Proh       float64
	LastUptate time.Time `orm:"auto_now"`
}

func init() {
	orm.RegisterModel(new(EmailObject), new(TrainHamResult), new(TrainSpamResult), new(AllPro))
}
