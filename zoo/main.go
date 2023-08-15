package main

import (
	"fmt"
	"math"

	"github.com/kakikubo/starting-go/zoo/animals"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	// コメントです。
	a := [3]string{
		"Yamada Taro",
		"Sato Hanako",
		"Suzuki Kenji",
	}
	fmt.Println(a[0], a[1], a[2])
	// 10進数の形式で数値5を%dの箇所へ埋め込む
	fmt.Printf("%d\n", 5) // => "数値=5"
	// 数値用の書式いろいろ
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17) // => "10進数=17 2進数=10001 8進数=21 16進数=11"
	// 埋め込むパラメータが足りない
	fmt.Printf("%d年%d月%d日\n", 2015, 7) // => "2015年7月%!d(MISSING)日"
	// 埋め込むパラメータが過剰
	fmt.Printf("%d年%d月%d日\n", 2015, 7, 15, 23)
	/* => "2015年7月15日
	%!(EXTRA int=23)" */

	// %vはさまざまな型のデータを埋め込む
	fmt.Printf("\n数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列=Golang 配列=[1 2 3]"
	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=5 文字列="Golang" 配列=[3]int{1, 2, 3}"
	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3}) // => "数値=int 文字列=string 配列=[3]int"

	print("Hello, World!\n") // => "Hello, World!" を標準エラー出力へ出力
	println("Hello, World!") // => "Hello, World!\n" を標準エラー出力へ出力
	print(1, 2, 3, "\n")     // => "123\n" を標準エラー出力へ出力
	println(1, 2, 3)         // => "1 2 3\n" を標準エラー出力へ出力

	var n int
	n = 5
	// n == 5
	// n = "string value" // コンパイルエラー
	fmt.Println(n)
	// int型の変数iを定義して1を代入(型推論)
	i := 1
	fmt.Println(i)
	// bool型の変数bを定義して真偽値trueを代入(型推論)
	b := true
	fmt.Println(b)
	// float64型の変数fを定義して浮動小数点数3.14を代入(型推論)
	f := 3.14
	fmt.Println(f)
	// string型の変数sを定義して文字列"abc"を代入(型推論)
	s := "abc"
	fmt.Println(s)

	// 関数の戻り値を元に型推論を行って変数を定義、初期化
	n2 := one()
	fmt.Println(n2)

	// varで変数定義をまとめる書き方
	var (
		n3 = 1
		s3 = "string value"
		b3 = true
	)
	// 暗黙的な定義を並べる書き方だと以下
	// n3 := 1
	// s3 := "string value"
	// b3 := true
	fmt.Println(n3, s3, b3)

	// パッケージ変数n1の値を用いて計算
	n = n1 + 1
	fmt.Printf("n=%d\n", n)

	var_sample()
	var_sample3()
	var_sample4()
	var_sample5()
	var_sample6()
	wrap_around()

	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)

	// 浮動小数点数
	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan)

	fmt.Println(1.0e2)  // 1.0 * 10^2
	fmt.Println(1.0e+2) // 1.0 * 10^2
	fmt.Println(1.0e-2) // 1.0 / 10^2

	doubleValueSample()
	complexSomething()
	runeSomething()
}

func one() int {
	return 1
}

func var_sample() {
	a := 1
	// 以下を定義しているとエラーになってしまう
	// b := 2
	// c := 3
	// # command-line-arguments
	// ./main.go:93:2: b declared and not used
	// ./main.go:94:2: c declared and not used

	fmt.Println(a)
}

func var_sample2() {
	n4 := 9223372036854775807 // 符号付き64ビット整数で表現可能である最大値
	fmt.Println(n4)

	// var (
	// 	n1 int
	// 	n2 int64
	// )
	// n1 = 1
	// n2 = n1 // コンパイルエラー
	// fmt.Println(n1, n2)

}

func var_sample3() {
	// n := uint(17)
	n := 1
	b := byte(n)
	i64 := int64(n)
	u32 := uint32(n)
	fmt.Println(b, i64, u32)
	fmt.Printf("%T %T %T\n", b, i64, u32)
}

func var_sample4() {
	n := 256
	b := byte(n)
	fmt.Printf("b = %s\n", b)
}

func var_sample5() {
	b := byte(255)
	b = b + 1
	fmt.Println(b)
}

func var_sample6() {
	n := -1
	b := byte(n)
	fmt.Println(b)
}

func wrap_around() {
	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	if !doSomething(ui_1, ui_2) {
		fmt.Println("エラーが発生しました")
		return
	}

	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d\n", ui_1, ui_2, sum)
	// 400000000 + 4000000000 = 105032704 オーバーフローして1億弱になってしまう
}

func doSomething(a, b uint32) bool {
	if (math.MaxInt32 - a) < b {
		return false
	} else {
		// チェック済みの為、問題なし
		sum := a + b
		fmt.Println(sum)
		return true
	}
}

func doubleValueSample() {
	fmt.Printf("value = %v\n", 1.0000000000000000)
	fmt.Printf("value = %v\n", 1.0000000000000001)
	fmt.Printf("value = %v\n", 1.0000000000000002)
	fmt.Printf("value = %v\n", 1.0000000000000003)
	fmt.Printf("value = %v\n", 1.0000000000000004)
	fmt.Printf("value = %v\n", 1.0000000000000005)
	fmt.Printf("value = %v\n", 1.0000000000000006)
	fmt.Printf("value = %v\n", 1.0000000000000007)
	fmt.Printf("value = %v\n", 1.0000000000000008)
	fmt.Printf("value = %v\n", 1.0000000000000009)
	fmt.Printf("value = %v\n", float32(1.0000000000000000))
	fmt.Printf("value = %v\n", float32(1.0000000000000001))
	fmt.Printf("value = %v\n", float32(1.0000000000000002))
	fmt.Printf("value = %v\n", float32(1.0000000000000003))
	fmt.Printf("value = %v\n", float32(1.0000000000000004))
	fmt.Printf("value = %v\n", float32(1.0000000000000005))
	fmt.Printf("value = %v\n", float32(1.0000000000000006))
	fmt.Printf("value = %v\n", float32(1.0000000000000007))
	fmt.Printf("value = %v\n", float32(1.0000000000000008))
	fmt.Printf("value = %v\n", float32(1.0000000000000009))
	fmt.Println(float32(1.0) / float32(3.0))
	fmt.Println(float64(1.0) / float64(3.0))

	f := 3.14
	n := int(f)
	fmt.Println(n)

	nf := -3.14
	nn := int(nf)
	fmt.Println(nn)
}

func complexSomething() {
	// 複素数型のサンプル
	c := 1.0 + 3i         // complex128型の変数cを定義して、1.0 + 3iを代入
	fmt.Println(c)        // 出力: (1+3i)
	c2 := complex(1.0, 3) // 別の定義方法
	fmt.Println(c2 == c)  // 出力: true
	// 複素数リテラル
	fmt.Println(0i)
	fmt.Println(11i)
	fmt.Println(0.i)
	fmt.Println(2.71828i)
	fmt.Println(6.67428e-11i)
	fmt.Println(1e6i)
	fmt.Println(.25i)
	fmt.Println(.12345e+5i)

	// 複素数の実部と虚部
	c3 := 1.3 + 4.2i
	fmt.Println(real(c3)) // real number (実数)
	fmt.Println(imag(c3)) // imaginary number (虚数)
}

func runeSomething() {
	r := '松'
	fmt.Printf("%v\n", r) // 出力: 26494
}