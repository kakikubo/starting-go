package main

import (
	"fmt"
	"testing"
)

// https://www.wakuwakubank.com/posts/792-go-goroutine-channel/ 参照

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan int)

	go receiver(ch) // チャネルを受信専用で受け取る関数をゴルーチンで実行

	i := 0
	for i < 100 {
		ch <- i
		i++
	}

}

func TestChannelLen(t *testing.T) {
	ch := make(chan int, 3)
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
