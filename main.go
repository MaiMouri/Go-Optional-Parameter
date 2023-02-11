package main

import (
	"flag"
	"fmt"
	"reflect"
)

var code string
var option string

func init() {
	flag.StringVar(&code, "code", "World", "コード")
	flag.StringVar(&option, "o", "", "実行オプション")
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

type CountryOpts struct {
	CountryCode *string
	runOption   *string
}

func Country(opts *CountryOpts) {
	// 引数がnilだったら未指定なのでデフォルト値で埋める
	countryCode := "World"

	// 引数がnilでない場合は、値を取り出す
	if opts.CountryCode != nil {
		fmt.Printf("countryCode passed: %v\n", countryCode)
		countryCode = *opts.CountryCode
		fmt.Printf("countryCode passed: %v\n", countryCode)
	}

	fmt.Printf("opt.runOption: %v\n", opts.runOption)
	val := reflect.ValueOf(opts.runOption)
	value := reflect.Indirect(val)
	fmt.Printf("opt.runOption(reflectを使用): %v\n", value)

	// 引数がnilでない場合は、このブロックが実行される
	if opts.runOption != nil {
		fmt.Println("With run option")
	}
	// flagを使用するとdefaultでnilにはならない
	var blank string = ""
	if opts.runOption == &blank {
		fmt.Println("No run option")
	}

	fmt.Printf("Hello, %s!\n", countryCode)
}

func main() {
	fmt.Println("-----")
	fmt.Println("-----")

	Greet("gopher", &GreetOpts{}) // Hello, gopher!

	word := "Hey"
	Greet("gopher", &GreetOpts{GreetingWord: &word}) // Hey, gopher!

	// 変数を渡す
	fmt.Println("------------------------------")
	fmt.Println("定義した変数を渡す")
	fmt.Println("------------------------------")
	country := "JP"
	Country(&CountryOpts{CountryCode: &country}) // Hello, JP!
	Country(&CountryOpts{})                      // Hello, World!

	// flagを使用してコマンドライン引数を渡す
	flag.Parse()

	fmt.Println("------------------------------")
	fmt.Println("(3) flag: コードのみ渡す")
	fmt.Println("------------------------------")
	Country(&CountryOpts{CountryCode: &code}) // Hello, UK!

	// opt.runOptionはnilになる

	fmt.Println("------------------------------")
	fmt.Println("(4)flag: 1.コード, 2.実行オプション")
	fmt.Println("------------------------------")
	Country(&CountryOpts{CountryCode: &code, runOption: &option})
	// This is run option
	// Hello, UK!

	// 実行オプションとして-oをコマンドラインに渡さなくても
	// flagからデフォルト値が渡されるため
	// opt.runOptionはnilにならない

	// 1. 実行コマンド
	// go run main.go -code UK -o dry-run

	// 2. 実行コマンド (コードを渡さない)
	// go run main.go -o dry-run
}
