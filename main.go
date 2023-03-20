package main

import (
	"my.exbot/internal/config"
)

var (
	Env    = config.NewViper("./")
	Logger = config.NewLogger()
)

func main() {

}
