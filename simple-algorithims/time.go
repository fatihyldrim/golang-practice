/*
	Tarih türleri
	1 saniye delay
	Tarihde RFC formatlarına göre işlem yapma
	iki tarih arasındaki fark
	Özel format çıktısı alma
	Bir zamanı özel bir bölgeye göre ayarlama
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Unix time: ", time.Now().Unix())
	fmt.Println("String time: ", time.Now().String())
	fmt.Println("UnixNano time: ", time.Now().UnixNano())
	fmt.Println("Weekday time: ", time.Now().Weekday())
	fmt.Println("Yearday time: ", time.Now().YearDay())
	fmt.Println("Nanosecond time: ", time.Now().Nanosecond())

	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339Nano))

	time.Sleep(time.Second)
	tt := time.Now()
	fmt.Println("Difference: ", tt.Sub(t))

	format := t.Format("02 January 2006")
	fmt.Println(format)

	loc, _ := time.LoadLocation("Europe/Paris")
	parisTime := t.In(loc)
	fmt.Println("Paris ", parisTime)

}
