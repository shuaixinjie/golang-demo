package main

import (
	"flag"
	"golang-demo/goroutine"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	//https://books.studygolang.com/gopl-zh/ch5/ch5-10.html
	goroutine.Demo1()
}


