package main

import "Go-Programming-Cookbook/chapter1/bytestrings"

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
