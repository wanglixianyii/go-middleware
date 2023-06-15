package main

import "github.com/wanglixianyii/go-middleware/go-wire/wire"

func main() {

	userHandler := wire.InitializeHandler()

	userHandler.Run()
}
