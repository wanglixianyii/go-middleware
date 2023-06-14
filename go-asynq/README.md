Asynq是一个go语言实现的分布式任务队列和异步处理库，基于Redis

特点
- 保证至少执行一次任务
- 任务写入Redis后可以持久化
- 任务失败之后，会自动重试
- worker崩溃自动恢复
- 可是实现任务的优先级
- 任务可以进行编排
- 任务可以设定执行时间或者最长可执行的时间
- 支持中间件
- 可以使用 unique-option 来避免任务重复执行，实现唯一性
- 支持 Redis Cluster 和 Redis Sentinels 以达成高可用性
- 作者提供了Web UI & CLI Tool让大家查看任务的执行情况