package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var (
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

	var data map[string]any
	dec := json.NewDecoder(strings.NewReader(flagData))
	if err := dec.Decode(&data); err != nil {
		return err
	}

	// TODO: ホームディレクトリを取得し、結果を変数homeと変数errに代入
	
	if err != nil {
		return err
	}

	// TODO: 変数homeと".mydevtool", "template", "*.tmpl"をfilepath.Joinで結合
	
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	templateName := flag.Arg(0)
	var src bytes.Buffer
	// TODO: &srcに対してテンプレートにデータを埋め込んで書き込む
	// テンプレート名は変数templateNameの値
	// データは変数dataの値
	if /* ココにコードを書く */; err != nil {
		return err
	}

	fmtSrc, err := format.Source(src.Bytes())
	if err != nil {
		return err
	}

	fmt.Print(string(fmtSrc))

	return nil
}

func parseFlag() error {

	flag.StringVar(&flagData, "data", "null", "埋め込まれるデータ")

	flag.Parse()

	return nil
}
