package main

import (
	"github.com/hibiken/asynq"
	"github.com/wanglixianyii/go-middleware/go-asynq/task"
	"log"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     "101.42.237.244:6379",
			Password: "",
			DB:       0,
		},
		asynq.Config{
			Concurrency: 3,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeExample, task.HandleTaskExample)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

	log.Println("[rocketmq] server listening on: %s")

}
