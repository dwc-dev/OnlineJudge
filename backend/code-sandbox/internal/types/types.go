package types

import "time"

type LanguageStrategy interface {
	Prepare() error
	Compile() error
	Execute() ([]ExecResult, error)
}

// ExecResult 保存执行结果数据
type ExecResult struct {
	Output      string        // 程序输出
	Time        time.Duration // 运行时间，单位为纳秒（Nanoseconds）
	MemoryUsage int64         // 占用内存
}
