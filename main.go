package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
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
	CountryCode string
}
type RunOpts struct {
	runOption *string
}

func Country(opts *CountryOpts) {
	// 引数がnilだったら未指定なのでデフォルト値で埋める
	countryCode := "World"

	// 引数がnilでない場合は、値を取り出す
	if opts.CountryCode != "" {
		fmt.Printf("引数を渡す前のcountryCode: %v\n", countryCode)
		countryCode = opts.CountryCode
		fmt.Printf("countryCode = *opts.CountryCode後のcountryCode: %v\n", countryCode)
	}

	// // flag使用のため、opt.runOptionはnilにならない
	// fmt.Printf("opt.runOption: %v\n", opts.runOption)
	// value := reflect.Indirect(reflect.ValueOf(opts.runOption))
	// fmt.Printf("reflect.Indirect(reflect.ValueOf(opts.runOption): %v\n", value)

	// fmt.Println("")

	// // 引数がnilでない場合は、このブロックが実行される
	// if opts.runOption != nil {
	// 	fmt.Printf("opts.runOption != nil: %v\n", opts.runOption != nil)
	// }

	// // 構造体比較のためストリング""を定義
	// var bl string = ""
	// type Blank struct {
	// 	blank *string
	// }
	// blank := Blank{
	// 	blank: &bl,
	// }

	// // flagを使用するとdefaultでnilにはならない
	// if &opts.runOption == &blank.blank {
	// 	fmt.Printf("&opts.runOption == &blank.blank: %v\n", &opts.runOption == &blank.blank)
	// 	fmt.Println("No run option")
	// }

	// fmt.Printf("reflect.DeepEqual(opts.runOption, blank.blank): %v\n\n", reflect.DeepEqual(opts.runOption, blank.blank))

	// if reflect.DeepEqual(opts.runOption, blank.blank) {
	// 	fmt.Printf("reflect.DeepEqual(opts.runOption, blank.blank): %v\n", reflect.DeepEqual(opts.runOption, blank.blank))
	// }

	fmt.Printf("Hello, %s!\n", countryCode)
}

func RunOption(opts *RunOpts) {

	// flag使用のため、opt.runOptionはnilにならない
	fmt.Printf("opt.runOption: %v\n", opts.runOption)
	value := reflect.Indirect(reflect.ValueOf(opts.runOption))
	fmt.Printf("reflect.Indirect(reflect.ValueOf(opts.runOption): %v\n", value)

	fmt.Println("")

	// 引数がnilでない場合は、このブロックが実行される
	if opts.runOption != nil {
		fmt.Printf("opts.runOption != nil: %v\n", opts.runOption != nil)
	}

	// 構造体比較のためストリング""を定義
	var bl string = ""
	var dryRun string = "dry-run"
	var force string = "force"
	type Blank struct {
		blank *string
	}
	blank := Blank{
		blank: &bl,
	}

	type RunOption struct {
		blank  *string
		dryRun *string
		force  *string
	}
	// cmp.Equalを使用するためには、
	// 構造体のフィールドが全て比較対象と同じである必要がある
	// この場合はポインタ型
	runOption := RunOption{
		blank:  &bl,
		dryRun: &dryRun,
		force:  &force,
	}

	// flagを使用するとdefaultでnilにはならない
	if &opts.runOption == &blank.blank {
		fmt.Printf("&opts.runOption == &blank.blank: %v\n", &opts.runOption == &blank.blank)
		fmt.Println("No run option")
	}

	fmt.Printf("reflect.DeepEqual(opts.runOption, blank.blank): %v\n\n", reflect.DeepEqual(opts.runOption, blank.blank))

	if reflect.DeepEqual(opts.runOption, blank.blank) {
		fmt.Printf("reflect.DeepEqual(opts.runOption, blank.blank): %v\n", reflect.DeepEqual(opts.runOption, blank.blank))
	}

	// go-cmpを使用した場合
	if cmp.Equal(opts.runOption, blank.blank) {
		fmt.Printf("cmp.Equal(opts.runOption, blank.blank): %v\n", cmp.Equal(opts.runOption, blank.blank))
	}

	if cmp.Equal(opts.runOption, runOption.dryRun) {
		fmt.Printf("i) cmp.Equal(opts.runOption, runOption.dryRun): %v\n", cmp.Equal(opts.runOption, runOption.dryRun))
	} else if cmp.Equal(opts.runOption, runOption.force) {
		fmt.Printf("ii) cmp.Equal(opts.runOption, runOption.force): %v\n", cmp.Equal(opts.runOption, runOption.force))
	} else if cmp.Equal(opts.runOption, runOption.blank) {
		fmt.Printf("iii) cmp.Equal(opts.runOption, runOption.blank): %v	", cmp.Equal(opts.runOption, runOption.blank))
	} else {
		fmt.Printf("No run option")
	}

}

func main() {
	// オプショナルパラメータの基本
	fmt.Println("------------------------------")
	fmt.Println("引数があってもなくても動く")
	fmt.Println("------------------------------")

	Greet("gopher", &GreetOpts{}) // Hello, gopher!

	word := "Hey"
	Greet("gopher", &GreetOpts{GreetingWord: &word}) // Hey, gopher!

	// 変数を渡す
	fmt.Println("-------------------------------------------")
	fmt.Println("コマンドライン引数からでも定義した変数でも渡せる")
	fmt.Println("-------------------------------------------")
	country := "JP"
	fmt.Printf("◆変数を渡す: %v\n", country)
	Country(&CountryOpts{CountryCode: country}) // Hello, JP!
	// Country(&CountryOpts{})                      // Hello, World!

	// flagを使用してコマンドライン引数を渡す
	flag.Parse()
	fmt.Printf("◆flagを使用してコマンドライン引数を渡す: %v\n", code)
	Country(&CountryOpts{CountryCode: code}) // Hello, UK!

	// opt.runOptionはnilになる

	fmt.Println("------------------------------")
	fmt.Println("実行オプション -o")
	fmt.Println("------------------------------")
	fmt.Printf("option == \"\": %v\n", option == "")

	// (A) main.goで分岐するパターン
	if option == "" {
		fmt.Println("実行オプションが渡されなかった場合の処理:")
		RunOption(&RunOpts{})
	} else {
		fmt.Println("実行オプションが渡された場合の処理:")
		RunOption(&RunOpts{runOption: &option})
	}

	// (B) 呼び先の関数内で分岐するパターン
	fmt.Println("呼び先の関数内で分岐するパターン:")
	RunOption(&RunOpts{runOption: &option})

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
