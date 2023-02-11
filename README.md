# Go-Optionao-Parameter

```go
package main
import "fmt"
func main(){

    type Person struct {
        Name string
        Age int
    }

    PersonMap := map[int]Person{}

    PersonMap[1] = Person{Name: "John", Age: 10}
    PersonMap[2] = Person{Name: "Tom", Age: 20}
    PersonMap[3] = Person{Name: "Mike"}                      // 未指定のフィールドはゼロ値が入る
    fmt.Println("構造体の未指定フィールドはゼロ値が自動で入れられる")
    fmt.Println("ゼロを入れた未指定だったかわからない")
    fmt.Println(PersonMap[1], PersonMap[2], PersonMap[3])　　// {John 10} {Tom 20} {Mike 0}
    // mapを再構成するから出力時にアドレスが毎回変わるため値変更ができない
    // PersonMap[1].Name = "Change"

    type Member struct {
        Name string
        Age *int
    }

    // 構造体はポインタで定義
    MemberMap := map[int]*Member{}

    var age int = 10
    MemberMap[1] = &Member{Name: "John", Age: &age}
    MemberMap[2] = &Member{Name: "Tom", Age: &age}
    MemberMap[3] = &Member{Name: "Mike"}               // 未指定のフィールドはnilが入る

    // ポインタのため値変更ができる
    MemberMap[1].Name = "Change"

    //
    fmt.Println("構造体の未指定フィールドがポインタだとnilが自動で入れられて明示的")
    fmt.Println(MemberMap[1], MemberMap[2].Age, MemberMap[3])
}
```



# 実行方法別のアウトプットの違い

## 1. 実行コマンドで引数を渡す

```shell
go run main.go -code UK -o dry-run
```

```shell
With run option
Hello, UK!
```

## 2. 実行コマンドで-codeを渡さない

```shell
go run main.go -o dry-run
```

```shell
With run option
Hello, World!
```
