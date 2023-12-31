package main

import (
	"fmt"
	"testing"
	"time"
)

// https://www.wakuwakubank.com/posts/792-go-goroutine-channel/ 参照

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func TestChannel(t *testing.T) {
	t.Skip()

	// TODO: このテストは何故かうまくいってない
	// <https://gist.github.com/kakikubo/0782223ec2d69360c689c54db3af0537>
	ch := make(chan int)

	go receiver(ch) // チャネルを受信専用で受け取る関数をゴルーチンで実行

	i := 0
	for i < 10 {
		ch <- i
		i++
	}
	// Unordered Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}

func TestChannelLen(t *testing.T) {
	ch := make(chan int, 3)
	if len(ch) != 0 {
		t.Errorf("チャネルの長さが異なります。")
	}
	ch <- 1
	ch <- 1
	ch <- 1
	if len(ch) != 3 {
		t.Errorf("チャネルの長さが異なります。")
	}
	// 書いてはいけないコード。0は保証されない
	// if len(ch) > 0 {
	// 	i := <-ch
	// }
	fmt.Println(len(ch))
}

func TestChannelCap(t *testing.T) {
	ch := make(chan int, 3)
	if cap(ch) != 3 {
		t.Errorf("チャネルの容量が異なります。")
	}
	ch <- 1
	ch <- 1
	ch <- 1
	if cap(ch) != 3 {
		t.Errorf("チャネルの容量が異なります。")
	}
	fmt.Println(cap(ch))
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	// ch <- 1 // クローズしているので送信はできない

	var (
		i  int
		ok bool
	)
	i, ok = <-ch
	fmt.Println(i, ok) // 1 true
	i, ok = <-ch
	fmt.Println(i, ok) // 2 true
	i, ok = <-ch
	fmt.Println(i, ok) // 3 true
	i, ok = <-ch
	fmt.Println(i, ok) // 0 false
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			/* 受信できなくなったら終了 */
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func TestGoRoutine(t *testing.T) {
	ch := make(chan int, 20)

	// どれが受信するかはわからない
	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(1 * time.Second)
}

func TestChannelFor(t *testing.T) {
	ch := make(chan int, 20)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}

	time.Sleep(3 * time.Second)
}

func TestChannelSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 20

	for i := 0; i < 3; i++ {
		select {
		case <-ch1:
			fmt.Println("ch1から受信しました。")
		case <-ch2:
			fmt.Println("ch2から受信しました。")
		case ch3 <- 30:
			fmt.Println("ch3へ送信しました。")
		default:
			fmt.Println("ここへは到達しない。")
		}
	}
}

func TestChannelSelect2(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	/* ch1から受信した整数を2倍してch2へ送信 */
	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()
	/* ch2から受信した整数を1減算してch3へ送信 */
	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		/* 整数を増分させてch1へ送信 */
		case ch1 <- n:
			n++
		/* ch3から受信した整数を出力 */
		case i := <-ch3:
			fmt.Println("received", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}

// func TestChannelAssignDummy(t *testing.T) {
// 	var ch0 chan int // 送受信可能チャネル
// 	var ch1 <-chan int // 受信専用チャネル
// 	var ch2 chan<- int // 送信専用チャネル
//
// 	送受信可能チャネルは、送信専用チャネル、受信専用チャネルに代入可能
// 	ch1 = ch0
// 	ch2 = ch0
// 	送信専用チャネルや受信専用チャネルは、送受信可能チャネルに代入不可
// 	ch0 = ch1 // コンパイルエラー
// 	ch0 = ch2 // コンパイルエラー
// 	受信専用チャネルは、送信専用チャネルには代入できず、その逆も同じ
// 	ch1 = ch2 // コンパイルエラー
// 	ch2 = ch1 // コンパイルエラー
// }

// func TestPanicChannel(t *testing.T) {
// 	ch := make(chan int, 3)
// 	ch <- 1
// 	ch <- 1
// 	ch <- 1
// 	ch <- 1 // デッドロック発生
// 	close(ch)
/* バッファサイズを指定しない場合は、バッファサイズ0のチャネルとなる */
// ch := make(chan int)
// ch <- 1 // 1をチャネルに送信
// i := <-ch
// fmt.Println(i)

// /* バッファサイズを指定する場合は、バッファサイズ分の要素を持つチャネルとなる */
// ch10 := make(chan int, 10)
// ch10 <- 1
// ch10 <- 2
// ch10 <- 3
// i10 := <-ch10
// fmt.Println(i10)
// i10 = <-ch10
// fmt.Println(i10)
// i10 = <-ch10
// fmt.Println(i10)

// ch4 := make(chan int, 8)
// /* チャネル4に整数5を「送信」 */
// ch4 <- 5
// /* チャネル4から整数を「受信」*/
// i := <-ch4
// fmt.Println(ch1)
// fmt.Println(ch2)
// fmt.Println(i)
// }
