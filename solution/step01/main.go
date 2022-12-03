package main

import (
	"fmt"
)

func main() {
	msg := "hello"
	// TODO: 変数nameに"tenntenn"を代入して宣言する
	name := "tenntenn"

	fmt.Println("package main")
	fmt.Println(`import "fmt"`)
	fmt.Println("func main() {")
	// TODO: "hello tenntenn"のようにfmt.Printfを使って表示する
	fmt.Printf(`fmt.Println("%s %s")`, msg, name)
	fmt.Println()
	fmt.Println("}")
}
