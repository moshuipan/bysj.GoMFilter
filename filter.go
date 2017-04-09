package main

import (
	"bufio"
	//	"math"
	"strconv"
	//	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/agonopol/go-stem"
	"github.com/astaxie/beego/orm"
)

type wordMap map[string]bool

var (
	re  *regexp.Regexp
	all = make(map[string]*AllPro)
)

const (
	Spamdir = "E:\\mygo\\src\\bysj.GoMFilter\\train\\spam"
	Hamdir  = "E:\\mygo\\src\\bysj.GoMFilter\\train\\ham"
)

func init() {
	pattern := `([a-zA-Z]+['’][a-zA-Z]\b)|\$?\d+(\.\d+)?%?|(\b[A-Za-z]\.)+[A-Za-z]\b|\w+(-\w+)*`
	re = regexp.MustCompile(pattern)
}

//训练好的过滤器对外接口
func Filter(data string) bool {
	wordmap := genWordMap(data)
	//=======================
	ps, ph := 0.5, 0.5
	var pwis float64 = 0.0
	var pwih float64 = 0.0
	for k, _ := range wordmap {
		if pro, ok := all[k]; ok {
			//			pwis += math.Log(pro.Pros)
			if pwis == 0 {
				pwis = pro.Pros
			} else {
				pwis *= pro.Pros
			}
		}
	}
	for k, _ := range wordmap {
		if pro, ok := all[k]; ok {
			//			pwih += math.Log(pro.Proh)
			if pwih == 0 {
				pwih = pro.Proh
			} else {
				pwih *= pro.Proh
			}
		}
	}
	pros := (ps * pwis) / (ps*pwis + ph*pwih)
	proh := (ph * pwih) / (ps*pwis + ph*pwih)
	log.Println("pros:", pros, "\tproh:", proh)
	return (pros / proh) < 1.2
}

//生成词干map
func genWordMap(data string) wordMap {
	bb := re.FindAll([]byte(strings.ToLower(data)), -1) //分词
	wordmap := make(wordMap)                            //词干map;词干提取?不提取词干
	for _, b := range bb {
		word := string(stemmer.Stem(b))
		//		word := string(b)
		if len(word) <= 1 { //长度小于=1
			continue
		}
		if _, exist := wordmap[word]; exist {
			wordmap[word] = true
		} else {
			wordmap[word] = false
		}
	}
	return wordmap
}

//E:\\mygo\\src\\bysj.GoMFilter\\train\\spam
func CountWordMap(dir string) (map[string]int, int) {
	log.Printf("==========open dir %s======\n", dir)
	d, err := os.Open(dir)
	if err != nil {
		log.Println("email ham dir:", err)
		os.Exit(0)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		log.Println("email ham dir fi:", err)
		os.Exit(0)
	}
	var reader *bufio.Reader
	var spamd string = ""
	fn := 0
	temp := make(map[string]int)
	for _, f := range fi {
		if !f.IsDir() {
			ff, err := os.Open(dir + "\\" + f.Name())
			if err != nil {
				log.Println("read email ham:", err)
				os.Exit(0)
			}
			fn++
			reader = bufio.NewReader(ff)
			for {
				s, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
					break
				}
				if err == io.EOF {
					spamd = spamd + s
					break
				}
				spamd = spamd + s
			}
			ff.Close()
			wordmap := genWordMap(spamd)
			for w, _ := range wordmap {
				if n, ok := temp[w]; ok {
					temp[w] = n + 1
				} else {
					temp[w] = 1
				}
			}
			spamd = ""
		}
	}
	//	隐藏单词频率超过fn，?小于=1的
	logf, err := os.OpenFile("trainlog.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if os.IsNotExist(err) {
		logf, err = os.Create("trainlog.txt")
		if err != nil {
			log.Println("create logfile:", err)
			os.Exit(0)
		}
	} else if err != nil {
		log.Println("open logfile:", err)
		os.Exit(0)
	}
	defer logf.Close()
	l1, _ := logf.Seek(0, os.SEEK_END)
	logf.WriteAt([]byte(dir+"\n"), l1)
	for w, n := range temp {
		if n <= 1 || n > fn {
			//		if n > fn {
			l, _ := logf.Seek(0, os.SEEK_END)
			msg := "word:" + w + "\t\t\ttimes:" + strconv.FormatInt(int64(n), 32) + "\n"
			logf.WriteAt([]byte(msg), l)
			delete(temp, w)
		}
	}
	log.Println("file nums:", fn)
	return temp, fn
}

func InitTrain() {
	o := orm.NewOrm()
	prospam := make(map[string]float64)
	proham := make(map[string]float64)
	swordmap, nums := CountWordMap(Spamdir)
	log.Println("====swordmap length:", len(swordmap), "======")
	for k, v := range swordmap {
		prospam[k] = float64(v) / float64(nums)
		//o.Insert(&TrainSpamResult{Key: k, Pro: prospam[k]})
	}
	go func() {
		for k, v := range prospam {
			o.Insert(&TrainSpamResult{Key: k, Pro: v})
		}
	}()
	hwordmap, nums := CountWordMap(Hamdir)
	log.Println("====hwordmap length:", len(hwordmap), "======")
	for k, v := range hwordmap {
		proham[k] = float64(v) / float64(nums)
		//o.Insert(&TrainHamResult{Key: k, Pro: proham[k]})
	}
	go func() {
		for k, v := range proham {
			o.Insert(&TrainHamResult{Key: k, Pro: v})
		}
	}()
	//	all := make(map[string]*AllPro)
	log.Println("========begin count all=====")
	for k, v := range prospam {
		if pro, ok := proham[k]; ok {
			all[k] = &AllPro{Key: k, Pros: v, Proh: pro}
		} else {
			all[k] = &AllPro{Key: k, Pros: v, Proh: 0.01}
		}
		//		o.Insert(all[k])
	}
	for k, v := range proham {
		if _, ok := prospam[k]; !ok {
			all[k] = &AllPro{Key: k, Pros: 0.01, Proh: v}
		}
		//		o.Insert(all[k])
	}
	go func() {
		for _, v := range all {
			o.Insert(v)
		}
	}()
}
