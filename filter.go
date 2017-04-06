package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/agonopol/go-stem"
	"github.com/astaxie/beego/orm"
)

type wordMap map[string]bool

//训练好的过滤器对外接口
func Filter(data string, o orm.Ormer) bool {

	return false
}

//生成词干map
func genWordMap(data string) wordMap {
	pattern := `([a-zA-Z]+['’][a-zA-Z]\b)|\$?\d+(\.\d+)?%?|(\b[A-Za-z]\.)+[A-Za-z]\b|\w+(-\w+)*`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("regexp compile errors:%v\n", err)
		os.Exit(0)
	}
	bb := re.FindAll([]byte(strings.ToLower(data)), -1) //分词
	wordmap := make(wordMap)                            //词干map;词干提取
	for _, b := range bb {
		word := string(stemmer.Stem(b))
		if len(word) <= 1 {
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

func TrainFilter() {
	d, err := os.Open("E:\\mygo\\src\\bysj.GoMFilter\\train\\spam")
	if err != nil {
		log.Println("spam dir:", err)
		os.Exit(0)
	}
	defer d.Close()
	fi, err := d.Readdir()
	if err != nil {
		log.Println("spam dir fi:", err)
		os.Exit(0)
	}
	for _, f := range fi {
		if !f.IsDir() {
			log.Println(f.Name())
		}
	}
}
