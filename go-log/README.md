使用zap日志库和lumberjack进行日志切割

可用于gin，kratos框架中

需要安装：

go get -u go.uber.org/zap
go get -u github.com/natefinch/lumberjack


1. 能够将事件记录到文件中，而不是应用程序控制台;
2. 日志切割-能够根据文件大小来切割日志文件;
3. 支持不同的日志级别。例如INFO，DEBUG，ERROR等;
4. 能够打印基本信息，如调用文件/函数名和行号，日志时间等;

以上就是本文所有内容了