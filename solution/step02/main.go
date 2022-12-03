package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}

func run() error {

	var tmpl string
	// TODO: コマンドライン引数がある場合は変数tmplに代入
	// 0番目はプログラム名なので注意
	if len(os.Args) > 1 {
		tmpl = os.Args[1]
	}

	switch tmpl {
	// TODO: "greeting"の場合
	case "greeting":
		return printGreeting()
	// TODO: それ以外の場合
	default:
		return printDefault()
	}

	return nil
}

func printGreeting() error {
	msg := "hello"
	name := "tenntenn"

	var src bytes.Buffer
	fmt.Fprintln(&src, "package main")
	fmt.Fprintln(&src, `import "fmt"`)
	fmt.Fprintln(&src, "func main() {")
	fmt.Fprintf(&src, `fmt.Println("%s %s")`, msg, name)
	fmt.Fprintln(&src)
	fmt.Fprintln(&src, "}")

	fmtSrc, err := format.Source(src.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(fmtSrc))

	return nil
}

func printDefault() error {
	var src bytes.Buffer
	fmt.Fprintln(&src, "package main")
	fmt.Fprintln(&src, `import "fmt"`)
	fmt.Fprintln(&src, "func main() {")
	fmt.Fprintln(&src, `fmt.Println("Hello, 世界")`)
	fmt.Fprintln(&src, "}")

	fmtSrc, err := format.Source(src.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(fmtSrc))

	return nil
}
