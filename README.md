<h1 align='center'>gribbon</h1>
<div align=center><img src="https://github.com/sugtex/gribbon/blob/main/workGo.jpg"/></div>
<h2 align='center'>A Goroutine Pool For Go</h2>

## 📖 简介

`gribbon` 是一个简单高效的协程池三方库，通过模拟上发条的方式实现。采用单向链表维护创建的节点 `node`，节点数据域为核心工作协程 `worker`。 `worker`中含有标示位用于区分，用户提交 `task` 会遍历链表，判断标示位决定是否使用该 `worker` 。理论上，对于短作业的函数，会有相当大的优势。允许使用者在开发并发程序的时候限制协程数量。

## 🚀 功能

- 以任务类型区分划分池，池子具有单一功能。
- 自动调度大量的`goroutine`，回收复用`goroutine`，遵循遍历最先执行最先释放原则。
- 向外抛出异常时提供友好的访问。
- 库兼容并发安全，向用户屏蔽并发细节。
- 异步安全关闭，主要线程不受影响。
- 采用链表，无需考虑内存，极大节省内存使用量。
- 在大规模批量短作业并发任务场景下更能体现高效低耗，极大地提升了性能。

## 🧰 安装
``` powershell
go get -u github.com/sugtex/gribbon
```

## 🛠 使用

### 两种任务类别
``` 
func hello(ctx context.Context) {
	// TODO 业务逻辑
}

func helloWithArg(ctx context.Context,arg interface{}){
	// TODO 断言逻辑
	// TODO 业务逻辑
}
```

### 默认池
``` 
if err := gribbon.Submit(context.Background(),hello); err != nil {
   // 处理err
}
if err := gribbon.SubmitWithArg(context.Background(),1,helloWithArg); err != nil {
   // 处理err
}
```

### 异常处理
```
// 是否为库内置异常
func IsGribbonErr(err error)bool{
	return IsInvalidCap(err)||IsWrongSubmit(err)||IsOverMaxCap(err)||IsClosed(err)
}

// 是否无效容量
func IsInvalidCap(err error)bool{
	return strings.EqualFold(err.Error(),errInvalidCap.Error())
}

// 是否错误任务提交
func IsWrongSubmit(err error)bool{
	return strings.EqualFold(err.Error(),errWrongSubmit.Error())
}

// 是否到达容量限制
func IsOverMaxCap(err error) bool {
	return strings.EqualFold(err.Error(), errOverMaxCap.Error())
}

// 是否处于关闭状态
func IsClosed(err error) bool {
	return strings.EqualFold(err.Error(), errClosed.Error())
}
```

### 实践
``` 
// 构建无参数池[提交含参数任务则抛出异常]
pool, err := gribbon.NewGoLink(10, false)
if err != nil {
	// 处理err
	return
}

if err := pool.Submit(context.Background(), hello); err != nil {
	// 处理err
	return
}

// 构建含参数池[提交无参数任务则抛出异常]
pool, err := gribbon.NewGoLink(10, true)
if err != nil {
	// 处理err
	return
}

if err := pool.SubmitWithArg(context.Background(), 1, helloWithArg); err != nil {
	// 处理err
	return
}

// 关闭池[异步并发安全]
if err:=pool.Close();err!=nil{
	// 处理err
	return
}
```

## 📚 附言

- 协程池不适用于死循环任务。
- `gribbon`在分配完用户提交的`task`后，“发条”回归，只需维护`head`节点，再次分配亦是如此。
- `gribbon`面对长延迟的`task`无法体现复用能力，`worker`处于忙碌状态无法回收，不超过用户设置或库默认最大容量的情况下会开辟新的`worker`进行`task`调度。
# gribbon
