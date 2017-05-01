package main

import (
	"go-smtpd/smtpd"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	DEBUG      = true
	DEBUG_MODE = true
	TEST       = true
)

var server smtpd.Server

type env struct {
	rcpts []smtpd.MailAddress //接收者
	from  smtpd.MailAddress   //发送者
	srv   smtpd.Server
	data  string //邮件内容
}

func (e *env) AddRecipient(rcpt smtpd.MailAddress) error {
	//	添加收件人
	e.rcpts = append(e.rcpts, rcpt)
	return nil
}

func (e *env) BeginData() error {
	if len(e.rcpts) == 0 {
		return smtpd.SMTPError("554 5.5.1 Error: no valid recipients")
	}
	return nil
}
func (e *env) Write(line []byte) error {
	e.data += string(line)
	return nil
}
func (e *env) Close() error {
	var emailOb *EmailObject
	o := orm.NewOrm()
	if !Filter2(e.data) {
		for _, v := range e.rcpts {
			if DEBUG && !times {
				if DEBUG_MODE {
					DeCrease(e.data)
				}
				return smtpd.SMTPError("554 Error: DEBUG MODE!")
			}
			emailOb = &EmailObject{
				From: e.from.Email(),
				Rcpt: v.Email(),
				Data: e.data,
				Ham:  false,
			}
			_, err := o.Insert(emailOb)
			if err != nil {
				log.Printf("%s---insert error:err=%V,--from %s to %s\n", time.Now().Format("2006-01-02 15:04:05"), err, emailOb.From, emailOb.Rcpt)
			}
		}
		return smtpd.SMTPError("554 Error: rejecting Spam Mail!")
	}
	//暂时只保存邮件
	for _, v := range e.rcpts {
		emailOb = &EmailObject{
			From: e.from.Email(),
			Rcpt: v.Email(),
			Data: e.data,
			Ham:  true,
		}
		_, err := o.Insert(emailOb)
		if err != nil {
			log.Printf("%s---insert error:err=%V,--from %s to %s\n", time.Now().Format("2006-01-02 15:04:05"), err, emailOb.From, emailOb.Rcpt)
		}
	}
	return nil
}

//返回一个新的信封对象
func onNewMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Printf("ajas: new mail from %q", from)
	//检查from...
	newmail := &env{
		rcpts: nil,
		from:  from,
		srv:   server,
		data:  "",
	}
	return newmail, nil
}

//处理新连接e.g.ip黑名单检查
func OnNewConnection(c smtpd.Connection) error {
	return nil
}

func main() {
	//	InitTrain()
	s := &smtpd.Server{
		Addr:            ":8000",
		OnNewMail:       onNewMail,
		OnNewConnection: OnNewConnection,
	}
	log.Println("====listen=====")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
