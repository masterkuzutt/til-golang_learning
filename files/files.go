package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const BSIZE = 100

func readTextFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open failed:", filename)
	}
	defer file.Close()

	buf := make([]byte, BSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]), n)
		fmt.Println(string(buf), "------")
	}
}

func writeTextFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("failed to create file:", filename)
	}
	defer file.Close()
	output := "output"
	file.Write(([]byte)(output))
	//　追記は自分で実装する必要がありそう。

}

func readTextFileUtil(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
	}
	fmt.Println(string(data))
	//サイズが大きいファイルには向いてないらしいが本当かは不明。
}

func writeTextFileUtil(filepath string) {
	msg := []byte("Hello, Gophers!")
	err := ioutil.WriteFile(filepath, msg, 0644)
	if err != nil {
		// log.Fatal(err)
	}
	// 追記も対応している。⇒してない。ファイル作らないだけ？パーミッションに影響あるだけ？
}

func readTextFileEachRow(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
	}
	defer file.Close()

	//　標準入力から読む時と同じやつか？
	sc := bufio.NewScanner(file)

	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
		}
		fmt.Printf("%v %v", i, sc.Text())
	}

}

func main() {
	exe, _ := os.Executable()
	dir := filepath.Dir(exe)
	fmt.Println(dir)

	path := "C:\\Users\\t_sakai\\go\\src\\local\\til-golang_learning\\files\\test.txt"
	// writeTextFile(path)　//ファイル上書きしてしまう。
	readTextFile(path)

	writeTextFileUtil(path)
	readTextFileUtil(path)
	// 一行ずつ読み込み
	readTextFileEachRow(path)
}
