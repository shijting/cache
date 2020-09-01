package cache

import "time"

/**
一个简易的内存缓存系统
1.支持设定过期时间，精度为秒级。
2．支持设定最大内存，当内存超出时候做出合理的处理。
3．支持并发安全。
4.为简化编程细节，无需实现数据落地。
*/
type ICache interface {
	//size 是⼀一个字符串串。⽀支持以下参数: 1KB，100KB，1MB，2MB 等
	SetMaxMemory(size string)
	// 设置⼀个缓存项，并且在expire时间之后过期
	Set(key string, val interface{}, exp time.Duration)
	// 获取⼀个值
	Get(key string) (interface{}, bool)
	// 删除⼀个值
	Del(key string) bool
	//检测⼀个值 是否存在
	Exists(key string) bool
	//清空所有值
	Flush() bool
	//返回所有的key数量
	Keys() int64
}
