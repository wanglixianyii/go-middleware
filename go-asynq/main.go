package main

import (
	"github.com/wanglixianyii/go-middleware/go-asynq/client"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		client.TaskAdd(i)
		time.Sleep(time.Second * 3)
	}
}
