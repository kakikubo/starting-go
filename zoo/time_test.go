package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeStandard(t *testing.T) {
	tm := time.Date(2023, 8, 31, 7, 15, 30, 0, time.Local)
	expect := "2023-08-31 07:15:30"
	actual := tm.Format(fixedTimeString)
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
	fmt.Println(tm)
	fmt.Println(tm.Year())
	fmt.Println(tm.YearDay())
	fmt.Println(tm.Month())
	fmt.Println(tm.Day())
	fmt.Println(tm.Hour())
	fmt.Println(tm.Minute())
	fmt.Println(tm.Second())
	fmt.Println(tm.Nanosecond())
	fmt.Println(tm.Location())
	fmt.Println(tm.Weekday())
	fmt.Println(tm.Zone())
	fmt.Println(time.July.String())
	fmt.Println(time.Sunday.String())
}

func TestTimeDuration(t *testing.T) {
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)
	/* 「2時間30分」を表すDurationを文字列から作成 */
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	/* 特定の日時を表すTimeを作成 */
	tm := time.Date(2023, 9, 1, 4, 44, 30, 0, time.Local)
	tm = tm.Add(2*time.Hour + 15*time.Minute + 30*time.Second)

	expect := "2023-09-01 07:00:00"
	actual := tm.Format(fixedTimeString)

	// 日本時間2023/09/01 7:00:00を表すTimeが生成された事を確認
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
}

func TestTimeDurationSub(t *testing.T) {
	tm := time.Date(2023, 9, 1, 7, 0, 0, 0, time.Local)
	tm2 := time.Date(2018, 8, 8, 7, 0, 0, 0, time.Local)
	d := tm.Sub(tm2)

	expect := "44400h0m0s"
	actual := d.String()
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
}

func TestTimeBeforeAfterEqual(t *testing.T) {
	t0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.Local)
	t1 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)

	// t0 < t1 → true
	if t1.Before(t0) {
		t.Errorf("t1.Before(t0) is false")
	}
	// t1 < t0 → false
	if !t1.After(t0) {
		t.Errorf("!t1.After(t0) is false")
	}
	// t0 > t1 → false
	if !t0.Before(t1) {
		t.Errorf("!t0.Before(t1) is false")
	}
	// t0 < t1 → true
	if t0.After(t1) {
		t.Errorf("t0.After(t1) is false")
	}

	// JSTを生成する
	time.Local = time.FixedZone("Local", 9*60*60)

	jst9 := time.Date(2015, 10, 1, 9, 0, 0, 0, time.Local) // 2015/10/1 9:00:00(JST)
	utc0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)   // 2015/10/1 0:00:00(UTC)

	if !jst9.Equal(utc0) {
		t.Errorf("!jst9.Equal(utc0) is false")
	}
}

func TestTimeAddDate(t *testing.T) {
	tm := time.Date(2023, 9, 1, 7, 0, 0, 0, time.Local)
	tm2 := tm.AddDate(1, 2, 3)

	expect := "2024-11-04 07:00:00"
	actual := tm2.Format("2006-01-02 15:04:05")
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}

	tm3 := tm2.AddDate(-1, -2, -3)
	if tm != tm3 {
		t.Errorf("tm != tm3")
	}
}

func TestTimeFormat(t *testing.T) {
	tm, err := time.Parse(fixedDateString, "2015/10/01")
	if err != nil {
		t.Errorf("err != nil")
	}
	expect := fixedDateTimeString
	actual := tm.Format(fixedTimeString)
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}

	// tm.Format で生成したtimeはUTCで生成される
	expect = "UTC"
	actual, _ = tm.Zone()
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}

	// time.ParseでRFC822形式の文字列をパースする
	tm2, err := time.Parse(time.RFC822, "01 Oct 15 00:00 JST")
	if err != nil {
		t.Errorf("err != nil")
	}

	expect = fixedDateTimeString
	actual = tm2.Format(fixedTimeString) // 時刻形式に%Yなどは使わない
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}

	// 日本語を混ぜた文字列もパース可能
	jt, _ := time.Parse("2006年01月02日", "2015年10月01日")
	expect = fixedDateTimeString
	actual = jt.Format(fixedTimeString)
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}

	// Parseに足りない要素(時,分,秒)は初期値が入る
	tm3, _ := time.Parse(fixedDateString, "2015/10/01")

	// ここは与えられた文字列からパースした部分
	if tm3.Year() != 2015 {
		t.Errorf("tm3.Year() != 2015")
	}
	if tm3.Month() != time.October {
		t.Errorf("tm3.Month() != time.October")
	}
	if tm3.Day() != 1 {
		t.Errorf("tm3.Day() != 1")
	}
	// ここからは初期値
	if tm3.Hour() != 0 {
		t.Errorf("tm3.Hour() != 0")
	}
	if tm3.Minute() != 0 {
		t.Errorf("tm3.Minute() != 0")
	}
}

func TestGenerateStringFromTime(t *testing.T) {
	// JSTを生成する
	time.Local = time.FixedZone("JST", 9*60*60)

	tm := time.Date(2023, 9, 1, 7, 0, 0, 0, time.Local)

	expectRFC822 := "01 Sep 23 07:00 JST"
	actualRFC822 := tm.Format(time.RFC822)
	if expectRFC822 != actualRFC822 {
		t.Errorf(stringExpectFormat, expectRFC822, actualRFC822)
	}

	expectRFC3339Nano := "2023-09-01T07:00:00+09:00"
	actualRFC3339Nano := tm.Format(time.RFC3339Nano)
	if expectRFC3339Nano != actualRFC3339Nano {
		t.Errorf(stringExpectFormat, expectRFC3339Nano, actualRFC3339Nano)
	}

	expectTokyo := "2023年9月1日 07時00分00秒"
	actualTokyo := tm.Format("2006年1月2日 15時04分05秒")
	if expectTokyo != actualTokyo {
		t.Errorf(stringExpectFormat, expectTokyo, actualTokyo)
	}

	expect := "2023/09/01"
	actual := tm.Format(fixedDateString)
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
}

func TestTimeUTC(t *testing.T) {
	// JSTを生成する
	time.Local = time.FixedZone("JST", 9*60*60)

	tm := time.Date(2023, 9, 1, 9, 0, 0, 0, time.Local)
	tmUTC := tm.UTC()

	expect := "2023-09-01 00:00:00 +0000 UTC"
	actual := tmUTC.String()
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
}

func TestTimeJST(t *testing.T) {
	// JSTを生成する
	time.Local = time.FixedZone("JST", 9*60*60)

	tm := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
	tmJST := tm.Local()

	expect := "2023-09-01 09:00:00 +0900 JST"
	actual := tmJST.String()
	if expect != actual {
		t.Errorf(stringExpectFormat, expect, actual)
	}
}

func TestTimeUnix(t *testing.T) {
	tm := time.Date(2015, 11, 27, 15, 0, 0, 0, time.UTC)
	expect := int64(1448636400)
	// timeからUnix時間を取得
	actual := tm.Unix()
	if expect != actual {
		t.Errorf("%d != %d", expect, actual)
	}

	tm2 := time.Unix(1448636400, 0)

	// Unix時間からtimeを生成
	expect2 := "2015-11-27 15:00:00 +0000 UTC"
	actual2 := tm2.UTC().String()
	if expect2 != actual2 {
		t.Errorf(stringExpectFormat, expect2, actual2)
	}
}

func TestTimeSleep(t *testing.T) {
	fmt.Println("start")
	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

func TestTimeTicker(t *testing.T) {
	t.Skip("skip TestTimeTicker")
	// https://devlights.hatenablog.com/entry/2020/11/13/135630
	// time.Tick だとgarbage collectionで回収されないのでtime.NewTickerを使う
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
