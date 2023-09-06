package main

import (
	"fmt"
	"math"
	"os"

	"github.com/kakikubo/starting-go/zoo/animals"
)

// n1はパッケージ変数
var n1 = 100

func main() {
	fmt.Println(AppName()) /* 関数AppNameの呼び出しを追加 */
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	// パッケージ変数n1の値を用いて計算
	n := n1 + 1
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
	someString()
	someArray()
	someInterface()
	someMath()
	fmt.Println(plus(1, 2))
	hello()
	q, r := div(19, 7)
	fmt.Printf("商=%d 余り=%d\n", q, r) // => "商=2 余り=5" quotient, remainder
	q2, _ := div(19, 7)              // 余りは捨てる
	fmt.Printf("商=%d\n", q2)         // => "商=2"

	fmt.Println(doSomethingA())
	fmt.Println(doSomethingXY())
	fmt.Println(ignoreArgs(1, 2))

	// fmt.Println(RequiredFunction(1))

	// 無名関数
	fn := func(x, y int) int { return x + y }
	fmt.Println(fn(2, 3))
	fmt.Printf("%T\n", fn) // => "func(int, int) int"
	ClosureSample()

	// 定義済みの関数に別名をつけているかのような記述
	var plusAlias = plus
	fmt.Println(plusAlias(10, 20))

	// 関数を返す関数
	rfn := returnFunc()
	rfn()
	fmt.Printf("%T\n", rfn) // => "func()"
	returnFunc()()          // こうやっても実行できる

	// 関数を引数にとる関数
	callFunction(func() {
		fmt.Println("I'm a callFunction")
	})

	// クロージャ(関数閉包)
	later := later()
	fmt.Println(later("Golang"))
	fmt.Println(later("is"))
	fmt.Println(later("awesome!"))
	fmt.Println(later("awesome!"))

	// クロージャを利用してGeneratorを実装する
	ints := integers()

	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3

	otherInts := integers()
	fmt.Println(otherInts()) // 1 (otherIntsの状態は別)

	// 定数
	// const X = 1
	const (
		X = 1
		Y = 2
		Z = 3
	)

	// 関数内での定数宣言
	x, y := onetwo()
	fmt.Printf("x=%d, y=%d\n", x, y) // => "x=1, y=2"

	// 定数定義(値の省略)
	const (
		XX = 10
		YY
		ZZ
		S1 = "あ"
		S2
	)
	fmt.Println(XX, YY, ZZ, S1, S2) // => "10 10 10 あ あ"

	const (
		XXX = 2
		YYY = 7
		ZZZ = XXX + YYY // ZZZ = 9

		S3  = "今日"
		S4  = "晴れ"
		S34 = S3 + "は" + S4 // S = "今日は晴れ"
	)

	fmt.Println(ZZZ, S34)
	fmt.Printf("%T\n", ZZZ) // => "int"
	// 以前はこの書き方でもOKだったみたい(X1, X2が定義されていない状態でX12を定義している)
	// const (
	// 	X12 = X1 + X2
	// 	X1 = 1
	// 	X2 = 2
	// )

	const (
		// I64 int64   = -1
		// F64 float64 = 1.2
		I64 = int64(-1)
		F64 = float64(1.2)
	)
	fmt.Printf("%T %v\n%T %v\n", I64, I64, F64, F64) // => "int64 -1\nfloat64 1.2"

	// 以下はオーバーフローする
	/*
		const (
			N = 999999999999999999999999999999999999999999999999999999999999999999999999999999999999999
		)
		n999 := N
		fmt.Printf("%T %v\n", n999, n999) // => "int64 999999
	*/

	// 型が違う場合はコンパイルエラーにしてくれる
	/*
		const (
			UI64 = uint64(12345)
		)

		var i64 int64
		i64 = UI64 // cannot use UI64 (constant 12345 of type uint64) as int64 value in assignment
	*/

	// コンパイル時に演算が処理される為、以下はコンパイルエラーにならない
	const (
		// uint64の最大値に1を足した値
		MAXUI64PLUS1 = math.MaxUint64 + 1
	)
	muint64 := uint64(MAXUI64PLUS1 - 1)     // コンパイル時に値が決定される為、18446744073709551615(uint64の最大値)になる
	fmt.Printf("%T %v\n", muint64, muint64) // => "uint64 18446744073709551615"

	// 浮動小数点数の定数
	const (
		Pi = 3.14
	)
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	fmt.Printf("%v\n", f32)
	fmt.Printf("%v\n", f64)
	const F = 1.0000000000001
	fmt.Println(float64(F) * 10000)
	fmt.Println(F * 10000)

	// 複素数の定数
	const (
		C = 4.7 + 1.3i
	)
	fmt.Printf("%T %v\n", C, C) // => "complex128 (4.7+1.3i)"

	// ルーン、文字列の定数
	const (
		R  = 'あ'
		S  = "Go言語"
		RS = `秋の田のかりほの庵の苫をあらみ
わが衣手は露にぬれつつ`
	)
	fmt.Printf("%v\n", R)  // => "12354"
	fmt.Printf("%v\n", S)  // => "Go言語"
	fmt.Printf("%v\n", RS) // => "秋の田のかりほの庵の苫をあらみ\nわが衣手は露にぬれつつ"

	// iota
	const (
		A1 = iota + 1
		B1
		C1
		N = iota
	)
	// iotaは定数の宣言ごとにリセットされる
	const (
		A2 = iota
		B2
		C2
	)
	fmt.Println(A1, B1, C1, N) // => "1 2 3 3"
	fmt.Println(A2, B2, C2)    // => "0 1 2"
	// やらないけど言語仕様上は以下のようにもかける

	const (
		朝の挨拶 = "おはよう"
		昼の挨拶 = "こんにちは"
		夜の挨拶 = "こんばんは"
	)
	あいさつ(昼の挨拶) // => "こんにちは"

	// 文字と認められていない〒(記号とされている)を使ってみる
	// const 〒 = "郵便番号"

	fmt.Println(animals.MAX)
	// fmt.Println(animals.internal_const) // コンパイルエラー

	fmt.Println(animals.FooFunc(5)) // => "6"
	// fmt.Println(animals.internalFunc(5)) // コンパイルエラー
	fmain()

	someCondition()
	funcSwitch()
	typeAssertion()
	typeAssertion2()
	typeSwitch()
	typeSwitch2()
	gotoSample()
	labelSample()
	labelSample2()
	runDefer()
	// runPanic()
	runRecover()

	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
	defer fmt.Println("!")

	fmt.Println("os.Exit")
	os.Exit(0)
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
	fmt.Printf("b = %b\n", b)
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

func someString() {
	s := "Goの文字列"
	fmt.Printf("%v\n", s) // 出力: Goの文字列

	s2 := `
GOの
RAW文字列リテラルによる
複数行に渡る
文字列
`
	fmt.Printf("%v\n", s2)
	s3 := `abc`
	fmt.Printf("%v\n", s3)
	s4 := `\n
\n`
	fmt.Printf("%v\n", s4)
}

func someArray() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", a[0]) // 出力: 1
	a1 := [5]int{}
	fmt.Printf("%v\n", a1)
	a2 := [5]int{1, 2, 3}
	fmt.Printf("%v\n", a2)
	// a3 := [5]int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("%v\n", a3) // ./main.go:281:30: index 5 is out of bounds (>= 5)
	var a3 [5]int
	a4 := [5]int{}
	fmt.Printf("%v\n", a3 == a4) // 出力: true

	ia := [3]int{}
	fmt.Printf("int array %v\n", ia) // 出力: [0 0 0]
	ua := [3]uint{}
	fmt.Printf("unsigned int array %v\n", ua) // 出力: [0 0 0]
	ba := [3]bool{}
	fmt.Printf("bool array %v\n", ba) // 出力: [false false false]
	fa := [3]float64{}
	fmt.Printf("float64 array %v\n", fa) // 出力: [0 0 0]
	ca := [3]complex128{}
	fmt.Printf("complex128 array %v\n", ca) // 出力: [(0+0i) (0+0i) (0+0i)]
	ra := [3]rune{}
	fmt.Printf("rune array %v\n", ra) // 出力: [0 0 0]
	sa := [3]string{}
	fmt.Printf("string array %v\n", sa) // 出力: [  ]
	zeroa := [0]int{}
	fmt.Printf("zero array %v\n", zeroa) // 出力: []

	a10 := [...]int{1, 2, 3}
	a11 := [...]int{1, 2, 3, 4, 5}
	a12 := [...]int{}
	fmt.Printf("%v\n%v\n%v\n", a10, a11, a12)

	a5 := [...]int{1, 2, 3}
	a5[0] = 0
	a5[2] = 0
	fmt.Printf("%v\n", a5) // 出力: [0 2 0]
	// 以下はエラーになる
	// var (
	// 	a6 [3]int
	// 	a7 [5]int
	// )
	// a6 = a7 // ./main.go:317:7: cannot use a7 (variable of type [5]int) as type [3]int value in assignment(exit status 1)

	// 異なる型の配列は代入できません
	// var (
	// 	a8 [5]int
	// 	a9 [5]uint
	// )
	// a8 = a9 // ./main.go:325:7: cannot use a9 (variable of type [5]uint) as [5]int value in assignment (exit status 1)

	a13 := [3]int{1, 2, 3}
	a14 := [3]int{4, 5, 6}
	a13 = a14
	a13[0] = 0
	a13[2] = 0
	fmt.Printf("a13 = %v\n", a13) // a13の値は[0 5 0]
	fmt.Printf("a14 = %v\n", a14) // a14の値は[4 5 6]※a13とa14は別の配列
}

func someInterface() {
	var x interface{}
	fmt.Printf("%#v\n", x) // 出力: <nil>
	x = 1
	x = 3.14
	x = '山'
	x = "文字列"
	x = [...]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", x) // 出力: [5]uint8{0x1, 0x2, 0x3, 0x4, 0x5}

	// interfaceはすべての型の値を汎用的に表す手段である為、演算の対象としては利用できない
	// var xx, yy interface{}
	// xx, yy = 1, 2 // 代入
	// z := xx + yy  // 演算できない
}

func someMath() {
	var n int
	s := "Go言語"
	x := 10
	n += 5
	fmt.Println(n)
	s += "の解説"
	fmt.Println(s)
	n *= 10
	fmt.Println(n)
	n &= x
	fmt.Println(n)
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello, World!")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

// 戻り値を表す変数
func doSomethingA() (a int) {
	return
	// 以下のコードと同じ
	// var a int
	// return a
}

func doSomethingXY() (x, y int) {
	y = 5
	return
	// 以下のコードと同じ
	// var x, y int
	// y = 5
	// return x, y
}

func ignoreArgs(_, _ int) int {
	return 1
}

// 型Tの定義
type T struct {
	value int
}

// インターフェース型I
type I interface {
	// 引数が2つ必要であると定義
	RequiredFunction(a, b int) int
}

// T型のインターフェースIを満たす関数(メソッド)
func (*T) RequiredFunction(a, _ int) int {
	// 実装に2番目の引数は不要
	return a
}

func ClosureSample() {
	var f func(int, int) int
	f = func(x, y int) int { return x + y }
	fmt.Println(f(1, 2))

	// fmt.Printf("ClosureSample %#v\n", func(x, y int) int { return x + y })
	fmt.Printf("ClosureSample %#v\n", func(x, y int) int { return x + y }(2, 3))
}

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	// 1つ前に与えられた文字列を保存する変数
	var store string
	// 引数に文字列を取り、文字列を返す関数を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

const ONE = 1

func onetwo() (int, int) {
	const TWO = 2
	return ONE, TWO
}

func あいさつ(m string) {
	fmt.Println(m)
}

/* パッケージ変数 */
var (
	m = 256 // パッケージ内のみで参照出来る変数
	N = 512 // 公開される変数
)

/* 公開される関数 */
func DoSomethingDo() {
	fmt.Println("DoSomething Do")
}

/* パッケージ内のみで参照できる関数 */
func doSomethingDo() {
	fmt.Println("doSomething Do")
}

func someCondition() {
	for {
		fmt.Println("loop")
		break // このbreakがないと無限ループになる
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	/* 変数iが100になるまでループ */
	i := 0
	for i < 100 {
		// continue使うとdlvが停止してしまうのでコメントアウト
		// if i == 50 {
		// 	continue
		// }
		fmt.Println(i)
		i++
	}
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	/* rangeを伴うfor */
	for i, s := range fruits {
		// iはインデックス、sは要素(値)を取る
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
	/* 文字列の場合はrune型になる */
	for i, r := range "あいう" {
		fmt.Printf("[%d]=%d %v\n", i, r, string(r))
	}

}

func funcSwitch() {
	n := 5
	// golangのswitchはbreakが不要(フォールスルーしない)
	switch n := 3; n { // ここで定義したnはswitch内でのみ有効
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4です")
		fallthrough // フォールスルーさせる場合は明示的に記述する
	// case "3": // 型が違うとコンパイルエラー
	// 	fmt.Println("3")
	case 5.0: // 整数5と互換性がある為エラーにならない
		fmt.Println("5.0です")
	case 6 + 0i: // 整数6と複素数6+0iは同じである為エラーにならない
		fmt.Println("6+0i")
	default:
		fmt.Println("unknown")
	}
	fmt.Printf("結局nは%d\n", n) // ここで定義したnはswitch外でも有効

	nn := 4
	switch {
	case nn > 0 && nn < 3:
		fmt.Println("0 < nn < 3")
	case nn > 3 && nn < 6:
		fmt.Println("3 < nn < 6")
	}

	/* caseに定数と式が混在するswitch文はコンパイルエラー(正確には型が異なるとエラー) */
	// switch x := 1; x {
	// case 1, 2, 3:
	// 	fmt.Println(x)
	// case x > 3:
	// 	fmt.Println("x > 3")
	// }
}

func typeAssertion() {
	anything(1)
	anything(3.14)
	anything(4 + 5i)
	anything('海')
	anything("日本語")
	anything([...]int{1, 2, 3, 4, 5})

	var x interface{} = 3
	i := x.(int)
	// f := x.(float64) // panic: interface conversion: interface {} is int, not float64
	fmt.Println(i)

	var y interface{} = 3.14

	ii, isInt := y.(int)
	ff, isFloat64 := y.(float64)
	s, isString := y.(string)
	fmt.Println(ii, isInt)     // 0 false
	fmt.Println(ff, isFloat64) // 3.14 true
	fmt.Println(s, isString)   //  false

}

func typeAssertion2() {
	// 2つの値を返す型アサーションの結果で条件分岐
	var x interface{} = "abc"
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("int %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Printf("string %s\n", s)
	} else {
		fmt.Println("Unsupported type!")
	}
}

func anything(a interface{}) {
	fmt.Println(a)
}

func typeSwitch() {
	var x interface{} = 3

	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}
}

func typeSwitch2() {
	var x interface{} = 3

	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("%#v\n", v)
	}
}

func gotoSample() {
	fmt.Println("A")
	goto L
	fmt.Println("B")
L: /* ラベル */
	fmt.Println("C")
}

func labelSample() {
LOOP:
	for {
		for {
			for {
				fmt.Println("START")
				break LOOP
			}
			fmt.Println("ここは通らない")
		}
		fmt.Println("ここは通らない")
	}
	fmt.Println("END")
}

func labelSample2() {
L:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j > 1 {
				continue L
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println("ここは処理されない")
	}
}

func runDefer() {
	/* deferに登録された式は関数の終了時に評価される。実行順序は最後に登録したものから先に実行される */
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}() // deferに登録する関数を即時実行する場合は()を付ける
	fmt.Println("done")
}

func runPanic() {
	/* panic時でもdeferは実行される */
	defer fmt.Println("defer on runPanic")

	panic("runtime error") // ここでエラー終了
	fmt.Println("Hello, World!")
}

func runRecover() {
	/* recoverはpanic時に発生したエラー情報を取得する */
	defer func() {
		if x := recover(); x != nil {
			/* 変数xはpanicに渡されたinterface{}型の値 */
			fmt.Println("recover:", x)
		}
	}()
	panic("runtime error")
	/* これは実行されない */
	fmt.Println("Hello, World!")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			/* panicによるinterface{}型の値に応じて処理を分岐 */
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Printf("panic: unknown=%v\n", v)
			}
		}
	}()
	panic(src)
	return
}
