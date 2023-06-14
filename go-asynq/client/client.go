package client

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/wanglixianyii/go-middleware/go-asynq/task"

	"time"
)

func TaskAdd(i int) {

	// 初始化客户端
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     "101.42.237.244:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	// 初始化任务
	t, err := task.NewExampleTask(&task.PayloadExample{
		Order: "测试复合数据结构",
		Req:   i,
	})
	if err != nil {
		fmt.Println(err)
	}

	// 用ProcessIn延迟执行任务 10秒后
	info, err := client.Enqueue(t, asynq.ProcessIn(time.Second*30))
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}

}
