# 基于goLang语言开发的日志库

####有终端打印日志方法和文件存储型方法，并分为五种日志级别
```
日志的级别
1.debug //调试时的日志
2.trace  //搞分析的会用
3.info  //正常运行的日志
4.warning //警告性错误
5.error //有错误
6.fatal //严重错误
```

###具体应用
1.可设置日志级别依次排列设置相应级别，就会记录日常级别下的日志

2.可记录调用文件、方法名、以及行号

3.文件日志可自动分割，设置最大分割点后，自动分割并对老日志进行备份

4.可格式化输出日志内容

5.文件日志自动将erroe级别日志另存储为后缀.err文件

###终端打印使用实例
```
package main

import (
	"mylogger"
	"time"
)

func main() {
	log := mylogger.Newlog("debug") //设置日志级别
	//启动一个定时器
	ti := time.Tick(time.Second)
	for t := range ti {
		id := 1001000101
		name := "张三"
		log.Debug("这是一条debug日志 id:%d,name:%s,在频繁登录", id, name)//记录日志Debug级别
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条fatal日志 id:%d,name:%s,在频繁登录", id, name)
		t.Second()
	}
}
```

###日志文件存储使用实例

```
package main

import (
	"mylogger"
	"time"
)

func main() {
    //NewFileLogger（日志级别，日志文件名，最大日志容量（byte型））
	log := mylogger.NewFileLogger("info", "../object/log/runlog", 10*1024)
	//启动一个定时器
	ti := time.Tick(time.Second)
	for t := range ti {
		id := 1001000101
		name := "张三"
		log.Debug("这是一条debug日志 id:%d,name:%s,在频繁登录", id, name)
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条fatal日志 id:%d,name:%s,在频繁登录", id, name)
		t.Second()
	}
}
```