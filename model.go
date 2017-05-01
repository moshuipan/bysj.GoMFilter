package main

import (
	"time"
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
type AllMap struct {
	Id         int `orm:"pk;auto"`
	Key        string
	SpamN      int
	HamN       int
	SpamEmails int
	HamEmails  int
}
