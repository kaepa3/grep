package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SearchList []string
}

var config Config

func main() {
	filename := os.Args[1]
	println("read:" + filename)
	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		// エラー処理
	}
	defer fp.Close()
	readConfig()
	analyze(fp)
}

func readConfig() {
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}

func analyze(fp *os.File) {
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// ここで一行ずつ処理
		analyzeLine(scanner.Text())
	}
}

func analyzeLine(line string) {
	flg := false
	for _, v := range config.SearchList {
		r := regexp.MustCompile(v)
		for _, str := range r.FindAllStringSubmatch(line, -1) {
			for i, txt := range str {
				if !flg {
					flg = true
					fmt.Println("start->")
				}
				fmt.Println(strconv.Itoa(i) + ":" + txt)
			}
		}
	}
}
