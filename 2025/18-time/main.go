package main

import (
	"fmt"
	"time"
)

func main() {
	/* t := time.Now()

	   fmt.Println(t)

	   t1 := time.Unix(0, t.UnixNano())

	   fmt.Println(t1)*/

	/* t := time.Date(2024, 1, 12, 1, 2, 3, 1234, time.Local)
	   fmt.Println(t)
	   fmt.Println(t.Year())
	   fmt.Println(t.Month())
	   fmt.Println(t.Day())
	   fmt.Println(t.Hour())
	   fmt.Println(t.Minute())
	   fmt.Println(t.Second())
	   fmt.Println(t.Nanosecond())
	   fmt.Println(t.UnixNano())

	   y, m, d := t.Date()
	   fmt.Println(y, m, d)

	   h, min, s := t.Clock()
	   fmt.Println(h, min, s)*/
	// 时间向字符串转化
	t := time.Now()
	s := t.Format("2006-01-02 15:04:05")
	fmt.Println(s)

	//字符串转化为时间
	timer := "2025-12-11 01:12:23"

	t2, _ := time.Parse("2006-01-02 15:04:05", timer)
	fmt.Println(t2)

}
