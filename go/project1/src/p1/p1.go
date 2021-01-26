package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	/*
	가~갛 까지 Unicode Rune 순서 확인
	 */
	/*
	DecodeRuneInString is like DecodeRune but its input is a string. If s is empty it returns (RuneError, 0). Otherwise, if the encoding is invalid, it returns (RuneError, 1). Both are impossible results for correct, non-empty UTF-8.
	An encoding is invalid if it is incorrect UTF-8, encodes a rune that is out of range, or is not the shortest possible UTF-8 encoding for the value. No other validation is performed.

	Example
	Code:
	str := "Hello, 世界"  for len(str) > 0 {     r, size := utf8.DecodeRuneInString(str)     fmt.Printf("%c %v\n", r, size)
	*/
	r1, size := utf8.DecodeRuneInString("가")
	fmt.Println(r1, size)
	r2, size := utf8.DecodeRuneInString("갛")
	fmt.Println(r2, size)
	for i:= r1; i<=r2; i++{
		fmt.Println(string(i), i)
	}
}

