package main

import (
	"fmt"
	"strings"
)

// Words : 함수나 구조체 전역은 첫글자 대문자, 지역은 소문자
type Words struct {
	word string
}

func (w Words) swap() string {
	tmpRune := []rune(w.word)
	a := []string{string(tmpRune[0])}
	b := []string{string(tmpRune[1])}

	result := append(b, a...)
	return strings.Join(result, "")
}

func main() {
	myStr := Words{word: "새개"}
	fmt.Println("hello " + myStr.word)
	fmt.Println("What the " + myStr.swap())
}
