package main

import (
	"os"
	"playing-around/pkg/bootstrap"
	"strconv"
)

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	app := bootstrap.Create(port, os.Args[2])
	_ = app.Start()
}
