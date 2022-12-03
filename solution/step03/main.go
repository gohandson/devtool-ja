package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"os"
	"strings"
)

var (
	// TODO: flagDataというstring型の変数を宣言
	flagData string
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー：", err)
		os.Exit(1)
	}
}

func run() error {

	if err := parseFlag(); err != nil {
		return err
	}

	// TODO: キーがstring型、値はany型のマップ型の変数dataを宣言
	var data map[string]any
	dec := json.NewDecoder(strings.NewReader(flagData))
	// TODO: JSONをデコードして変数dataに入れる
	if err := dec.Decode(&data); err != nil {
		return err
	}

	tmpl := flag.Arg(0)
	switch tmpl {
	case "greeting":
		return printGreeting(data)
	default:
		return printDefault(data)
	}

	return nil
}

func parseFlag() error {

	// TODO: flag.StringVarでflagDataに-dataオプションでもらったデータを渡す
	// デフォルト値は"null"
	// 説明は"埋め込まれるデータ"
	flag.StringVar(&flagData, "data", "null", "埋め込まれるデータ")

	flag.Parse()

	return nil
}

func printGreeting(data map[string]any) error {
	var src bytes.Buffer
	fmt.Fprintln(&src, "package main")
	fmt.Fprintln(&src, `import "fmt"`)
	fmt.Fprintln(&src, "func main() {")
	fmt.Fprintf(&src, `fmt.Println("%s %s")`, data["msg"], data["name"])
	fmt.Fprintln(&src)
	fmt.Fprintln(&src, "}")

	fmtSrc, err := format.Source(src.Bytes())
	if err != nil {
		return err
	}

	fmt.Println(string(fmtSrc))

	return nil
}

func printDefault(data map[string]any) error {
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
