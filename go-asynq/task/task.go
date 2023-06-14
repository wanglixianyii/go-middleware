package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeExample = "asynq:example"
)

// PayloadExample 异步任务携带的负载
type PayloadExample struct {
	Order string
	Req   int
}

func NewExampleTask(payload *PayloadExample) (*asynq.Task, error) {

	fmt.Println("###初始化任务:", time.Now().Format("2006-01-02 15:04:05"))

	data, _ := json.Marshal(payload)

	return asynq.NewTask(TypeExample, data), nil
}

func HandleTaskExample(ctx context.Context, t *asynq.Task) error {

	fmt.Println("###执行任务 start:", time.Now().Format("2006-01-02 15:04:05"))

	var payload PayloadExample
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	fmt.Println(payload.Order, payload.Req)

	return nil
}
