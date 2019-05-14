# goapp
go 通用app类型，提供一些app常用方法

## 接口
```go
// App is a common App interface
type App interface {
  // Name 获取App名称
	Name() string
  
  // Version 返回App当前版本信息
	Version() string
  
  // Start 启动函数，成功返回true
	Start() bool
}
```

## 方法
```go
// CreatePidFile 用于创建PID文件
func CreatePidFile(app App) bool

// Stop 用于停止App，会向App发送SIGQUIT信号
func Stop()

// WaitExit 用于阻塞主线程,直到收到SIGQUIT/SIGINT/SIGTERM/SIGKILL/SIGSEGV信号
func WaitExit()
```
