package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := `このプログラムはstrings.NewReader()関数を使って与えられた文字列を
io.Readerインターフェースと互換性のあるstrings.Reader型を生成します
なにか文字列を入力してEnterキーをおしてみてください。終了はCtrl+Dです`

	/* 文字列からstrings.Readerを生成 */
	r := strings.NewReader(s)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/* 標準入力をソースにしたスキャナの生成 */
	scanner = bufio.NewScanner(os.Stdin)

	/* 入力のスキャンが成功する限り繰り返すループ */
	for scanner.Scan() {
		fmt.Println("=> ", scanner.Text()) // Ctrl+Dで終了
	}

	/* スキャン中にエラーが発生した場合の処理 */
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
