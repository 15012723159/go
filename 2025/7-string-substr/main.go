package main

import "fmt"

func main() {

	s := "smalling向"
	fmt.Println(len(s))
	s1 := s[:5]
	fmt.Println(s1)

	arr := []rune(s)
	s2 := fmt.Sprintf("%c", arr[8])
	fmt.Println(s2)
	fmt.Printf("%T", s2)

	for s3, _ := range s {
		fmt.Printf("%T", s[s3])
	}
}
