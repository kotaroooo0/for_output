package main

import (
	"io"
	"os"
)

func main() {
	// エラー処理を省略

	// のちにio.Readerとして使われるos.File
	oldFile, _ := os.Open("old.txt")

	// 終了したらファイルディスクリプタを閉じる
	defer oldFile.Close()

	// のちにio.Writerとして使われるos.File
	newFile, _ := os.Create("new.txt")

	// 終了したらファイルディスクリプタを閉じる
	defer newFile.Close()

	// oldFileからnewFileへコピー
	io.Copy(newFile, oldFile)
}
