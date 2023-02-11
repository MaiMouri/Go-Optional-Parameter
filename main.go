package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	// gorm

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var code string

func init() {
	flag.StringVar(&code, "code", "default", "code")
}

type GreetOpts struct {
	GreetingWord *string
}

// オプショナルパラメータを構造体で受け取る
func Greet(name string, opts *GreetOpts) {
	greetingWord := "Hello"
	if opts.GreetingWord != nil {
		// 引数がnilだったら未指定なのでデフォルト値で埋める
		greetingWord = *opts.GreetingWord
	}
	fmt.Printf("%s, %s!\n", greetingWord, name)
}

func main() {
	flag.Parse()

	Greet("gopher", &GreetOpts{}) // Hello, gopher!

	word := "Hey"
	Greet("gopher", &GreetOpts{GreetingWord: &word}) // Hey, gopher!

}
