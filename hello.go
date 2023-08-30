package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	// go build -o getargs hello.go
	/* os.Argsの要素数を表示 */
	fmt.Printf("length=%d\n", len(os.Args))

	/* os.Argsの内容を表示 */
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}

	/* ファイルの読み込み */
	bs, err := readFile("./go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

	cdZooAndBack()

	currentDirs()

	createHelloText("./hello.txt")

	moveFile("./hello.txt", "./hellotmp/hello.txt")

	openFile("./hellotmp/hello.txt")

	removeFile("./hellotmp")
}

// os.Open()でファイルを開き、その内容を読み込む
func readFile(filename string) ([]byte, error) {
	/* 変数fは*os.File型 */
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/* []byte型のスライスにファイルの内容を読み込む */
	bs := make([]byte, 128)
	n, err := f.Read(bs) // nは実際に読み込んだバイト数
	fmt.Printf("%d bytes read\n", n)

	/* ファイルのオフセットを指定して読み込む */
	// n, err := f.ReadAt(bs, 10) // 10バイト目から読み込む

	/* ファイル内のシーク */
	offset, err := f.Seek(10, io.SeekStart) // ファイルの先頭から10バイト目にシーク
	fmt.Printf("offset = %d err=%v\n", offset, err)
	offset, err = f.Seek(-2, io.SeekCurrent) // 現在位置から-2バイト先にシーク
	fmt.Printf("offset = %d err=%v\n", offset, err)
	offset, err = f.Seek(0, io.SeekEnd) // ファイルの末尾から0バイト目にシーク
	fmt.Printf("offset = %d err=%v\n", offset, err)

	/* ファイルのステータスを取得 */
	fi, err := f.Stat()                        // fiはos.FileInfo型
	fmt.Printf("name = %s\n", fi.Name())       // ファイル名(string型)
	fmt.Printf("size = %d\n", fi.Size())       // ファイルサイズ(int64型)
	fmt.Printf("mode = %v\n", fi.Mode())       // ファイルのモード(os.FileMode型)
	fmt.Printf("modtime = %v\n", fi.ModTime()) // 最終更新日時(time.Time型)
	fmt.Printf("isdir = %v\n", fi.IsDir())     // ディレクトリかどうか(bool型)
	return bs, err
}

// hello.txtを作成して文字列を書き込む
func createHelloText(filename string) {
	f, _ := os.Create(filename)
	defer f.Close()

	/* 新規作成したファイルのステータス */
	fi, _ := f.Stat()
	fmt.Printf("name = %s\n", fi.Name())   // ファイル名(string型)
	fmt.Printf("size = %d\n", fi.Size())   // ファイルサイズ(int64型)
	fmt.Printf("isdir = %v\n", fi.IsDir()) // ディレクトリかどうか(bool型)

	f.Write([]byte("Hello, World!\n")) // ファイルに[]byte型のスライスを書き込む
	f.WriteAt([]byte("Golang!"), 7)    // ファイルの7バイト目から[]byte型のスライスを書き込む
	f.Seek(0, io.SeekEnd)              // ファイルの末尾にシーク
	f.WriteString("Yeah!\n")           // ファイルに文字列を書き込む
}

func openFile(filename string) {
	/* ファイルを読み込み専用でオープン */
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	fmt.Printf("err=%v\n", err)
	defer f.Close()
	bs := make([]byte, 128)
	n, err := f.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes read\n", n)
}

func removeFile(filename string) {
	err := os.RemoveAll(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func moveFile(from, to string) {
	/* カレントディレクトリ下にhellotmpディレクトリを作成 */
	err := os.Mkdir("hellotmp", 0777)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename(from, to)
	if err != nil {
		log.Fatal(err)
	}
}

func cdZooAndBack() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	err = os.Chdir("zoo")
	if err != nil {
		log.Fatal(err)
	}

	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	err = os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func currentDirs() {
	/* カレントディレクトリのオープン */
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	/* カレントディレクトリ下のディレクトリを列挙 */
	fis, err := f.Readdir(0) // 0はすべてのファイルを表す
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----")
	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", fi.Name())
		} else {
			fmt.Printf("%s\n", fi.Name())
		}
	}
}
