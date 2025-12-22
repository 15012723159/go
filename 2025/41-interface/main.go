package main

import "fmt"

type Video interface {
	scan()
}
type Store interface {
	GetPreSignUrl(filename string) string
	PutPreSignUrl(url string)
	Video
}

type Aws struct {
	Name          string
	Access_key    string
	Access_Secret string
}

func (a *Aws) GetPreSignUrl(filename string) string {
	return filename + "?type=aws"
}

func (a *Aws) PutPreSignUrl(url string) {
	fmt.Println("开始put")
}
func (a *Aws) Scan() {
	fmt.Println("观看视频")
}

type Hw struct {
	Name          string
	Access_key    string
	Access_Secret string
}

func (a *Hw) GetPreSignUrl(filename string) string {
	return filename + "?type=hw"
}

func (a *Hw) PutPreSignUrl(url string) {
	fmt.Println("开始put")
}

func (a *Hw) Scan() {
	fmt.Println(a.Name + "开启浏览视频")
}
func main() {

	hw := new(Hw)
	hw.Name = "hw"
	hw.Access_Secret = "232"
	hw.Access_Secret = "322"
	fmt.Println(hw.GetPreSignUrl("test.stl"))
	hw.Scan()
}
