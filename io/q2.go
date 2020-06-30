package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// エラー処理を省略

	// ブラウザにzipファイルであることを伝える
	w.Header().Set("Content-Type", "application/zip")
	// ファイル名を指定
	w.Header().Set("Content-Disposition", "attachment; filename-ascii_sample.zip")

	// 出力先のWriterをzip.NewWriterに渡すとzipファイル書き込み用の構造体ができる
	zipWriter := zip.NewWriter(w)
	// 終了したらファイルディスクリプタを閉じる
	defer zipWriter.Close()

	// zipファイルに含めるファイルを作成
	a, _ := zipWriter.Create("a.txt")
	io.Copy(a, strings.NewReader("1つめのファイルのテキストです"))

	// zipファイルに含めるファイルを作成
	b, _ := zipWriter.Create("b.txt")
	io.Copy(b, strings.NewReader("2つめのファイルのテキストです"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
