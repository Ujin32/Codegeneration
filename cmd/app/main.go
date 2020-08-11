package main

import "module07/internal/analys"

func main() {
	if err := analys.Analys("/Users/i.shvyryalkin/Projects/rebrain/module07/internal/convertor/convertor.go"); err != nil {
		panic(err)
	}
}
